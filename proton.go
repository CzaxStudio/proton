package proton

import (
	"image"
	"image/color"
	"log"
	"os"

	"sync"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// App is the top-level application handle. Create one with New().
type App struct {
	theme      *material.Theme
	windows    []*winDef
	bgColor    *color.NRGBA
	bgGradient *gradient
	logo       *logoState

	// cached gradient op — rebuilt only when window size changes
	gradientCache struct {
		w, h int
		ops  op.Ops
		call op.CallOp
		ok   bool
	}
}

type winDef struct {
	title string
	w, h  int
	extra []WindowOption
	draw  func(Context)
}

// New creates a Proton application.
func New(name string) *App {
	return &App{theme: material.NewTheme()}
}

func (a *App) Window(title string, width, height int, draw func(Context)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, draw: draw})
}

func (a *App) WindowEx(title string, width, height int, opts []WindowOption, draw func(Context)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, extra: opts, draw: draw})
}

func (a *App) Run() {
	if len(a.windows) == 0 {
		log.Fatal("proton: Run() called with no windows registered")
	}
	for _, w := range a.windows[:len(a.windows)-1] {
		go runWin(a, w)
	}
	go func() {
		runWin(a, a.windows[len(a.windows)-1])
		os.Exit(0)
	}()
	app.Main()
}

// background helpers

type gradient struct {
	from, to color.NRGBA
	dir      string
}

func (a *App) SetBackground(c color.NRGBA) { a.bgColor = &c; a.bgGradient = nil }
func (a *App) SetBackgroundCode(code string) {
	c := parseHex(code)
	a.bgColor = &c
	a.bgGradient = nil
}
func (a *App) SetBackgroundRGB(r, g, b uint8) {
	c := color.NRGBA{R: r, G: g, B: b, A: 255}
	a.bgColor = &c
	a.bgGradient = nil
}
func (a *App) SetBackgroundGradient(from, to, dir string) {
	a.bgGradient = &gradient{from: parseHex(from), to: parseHex(to), dir: dir}
	a.bgColor = nil
	a.gradientCache.ok = false
}
func (a *App) SetBackgroundRainbow() {
	a.bgGradient = &gradient{dir: "rainbow"}
	a.bgColor = nil
	a.gradientCache.ok = false
}

// WindowOption

type WindowOption struct{ apply func(*app.Window) }

func Fullscreen() WindowOption {
	return WindowOption{apply: func(w *app.Window) { w.Option(app.Fullscreen.Option()) }}
}
func Maximized() WindowOption {
	return WindowOption{apply: func(w *app.Window) { w.Option(app.Maximized.Option()) }}
}

// frame loop

func runWin(a *App, def *winDef) {
	w := new(app.Window)
	w.Option(app.Title(def.title))
	w.Option(app.Size(unit.Dp(float32(def.w)), unit.Dp(float32(def.h))))
	for _, o := range def.extra {
		o.apply(w)
	}

	var ops op.Ops
	var frame int

	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			ops.Reset()
			gtx := app.NewContext(&ops, e)
			frame++

			drawBackground(gtx, a, w, frame)

			// Run the draw function to collect widgets, then lay them out
			// directly as a vertical Flex — no wrapping material.List, which
			// adds a full virtual-list pass even for non-scrolling content.
			layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				c := newWinImpl(a, w)
				def.draw(c)
				return c.flush(gtx)
			})

			e.Frame(gtx.Ops)

		case app.DestroyEvent:
			return
		}
	}
}

func drawBackground(gtx layout.Context, a *App, w *app.Window, frame int) {
	switch {
	case a.bgGradient != nil && a.bgGradient.dir == "rainbow":
		// Rainbow animates continuously — draw fresh and request next frame.
		drawRainbow(gtx, frame)
		w.Invalidate()

	case a.bgGradient != nil:
		// Static gradient: cache the op and only rebuild when size changes
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		if !a.gradientCache.ok || a.gradientCache.w != w || a.gradientCache.h != h {
			a.gradientCache.ops.Reset()
			r := op.Record(&a.gradientCache.ops)
			drawGradientInto(layout.Context{
				Ops:         &a.gradientCache.ops,
				Constraints: gtx.Constraints,
				Metric:      gtx.Metric,
			}, a.bgGradient)
			a.gradientCache.call = r.Stop()
			a.gradientCache.w = w
			a.gradientCache.h = h
			a.gradientCache.ok = true
		}
		a.gradientCache.call.Add(gtx.Ops)

	case a.bgColor != nil:
		paint.Fill(gtx.Ops, *a.bgColor)

	default:
		paint.Fill(gtx.Ops, a.theme.Palette.Bg)
	}
}

func drawGradientInto(gtx layout.Context, g *gradient) {
	w := gtx.Constraints.Max.X
	h := gtx.Constraints.Max.Y
	// 32 steps is plenty — the difference from 64 is invisible
	const steps = 32
	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps)
		t2 := float32(i+1) / float32(steps)
		c := lerpColor(g.from, g.to, t)
		var rect image.Rectangle
		switch g.dir {
		case "horizontal", "diagonal":
			rect = image.Rect(int(float32(w)*t), 0, int(float32(w)*t2), h)
		case "radial":
			c = lerpColor(g.from, g.to, 1-t)
			rect = image.Rect(0, int(float32(h)*t), w, int(float32(h)*t2))
		default: // vertical
			rect = image.Rect(0, int(float32(h)*t), w, int(float32(h)*t2))
		}
		paint.FillShape(gtx.Ops, c, clip.Rect(rect).Op())
	}
}

func drawRainbow(gtx layout.Context, frame int) {
	w := gtx.Constraints.Max.X
	h := gtx.Constraints.Max.Y
	const steps = 32
	shift := float32(frame%360) / 360.0
	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps)
		hue := t + shift
		hue -= float32(int(hue))
		c := hsvColor(hue, 0.55, 0.85)
		t2 := float32(i+1) / float32(steps)
		rect := image.Rect(0, int(float32(h)*t), w, int(float32(h)*t2))
		paint.FillShape(gtx.Ops, c, clip.Rect(rect).Op())
	}
}

func lerpColor(a, b color.NRGBA, t float32) color.NRGBA {
	lerp := func(x, y uint8) uint8 { return uint8(float32(x) + (float32(y)-float32(x))*t) }
	return color.NRGBA{R: lerp(a.R, b.R), G: lerp(a.G, b.G), B: lerp(a.B, b.B), A: 255}
}

func hsvColor(h, s, v float32) color.NRGBA {
	i := int(h * 6)
	f := h*6 - float32(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)
	var r, g, b float32
	switch i % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}
	return color.NRGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 255}
}

// Context is the only type in Proton's public API that touches the renderer.
// No Gio types leak through it.
//
//	a.Window("App", 480, 300, func(ctx proton.Context) {
//	    proton.Label(ctx, "Hello")
//	    if proton.Button(ctx, &btn, "Click") { ... }
//	})
type Context interface {
	// Invalidate requests a redraw on the next frame.
	// Call after changing state from a goroutine.
	Invalidate()

	add(fn gioWidget)
	theme() *material.Theme
	rawWindow() *app.Window
	appLogo() *logoState
}

type gioWidget = func(gtx layout.Context) layout.Dimensions

// winImpl is the unexported concrete implementation of Context.
// We reuse the widgets slice by capping it — this avoids the most
// common allocation hot-spot (growing the slice every frame).
type winImpl struct {
	th      *material.Theme
	win     *app.Window
	logo    *logoState
	widgets []gioWidget
}

// Pool of winImpl objects to reduce per-frame heap allocations.
// child() pulls from here instead of allocating every time.
var implPool = sync.Pool{New: func() any { return &winImpl{widgets: make([]gioWidget, 0, 32)} }}

func newWinImpl(a *App, w *app.Window) *winImpl {
	c := implPool.Get().(*winImpl)
	c.th = a.theme
	c.win = w
	c.logo = a.logo
	c.widgets = c.widgets[:0]
	return c
}

func (c *winImpl) Invalidate()           { c.win.Invalidate() }
func (c *winImpl) add(fn gioWidget)      { c.widgets = append(c.widgets, fn) }
func (c *winImpl) theme() *material.Theme { return c.th }
func (c *winImpl) rawWindow() *app.Window { return c.win }
func (c *winImpl) appLogo() *logoState   { return c.logo }

// flush lays out all collected widgets as a vertical Flex and returns
// the impl to the pool.
func (c *winImpl) flush(gtx layout.Context) layout.Dimensions {
	dims := c.run(gtx)
	implPool.Put(c)
	return dims
}

func (c *winImpl) run(gtx layout.Context) layout.Dimensions {
	if len(c.widgets) == 0 {
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}
	// Build FlexChild slice on the stack for small widget counts.
	// For larger counts this still heap-allocates, but that's unavoidable.
	children := make([]layout.FlexChild, len(c.widgets))
	for i, fn := range c.widgets {
		fn := fn
		children[i] = layout.Rigid(fn)
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

// child creates a nested Context, calls fn to collect its widgets, then
// returns a gioWidget that lays them out during Gio's live layout pass.
// Uses the pool to avoid per-call heap allocations.
func child(parent Context, fn func(Context)) gioWidget {
	return func(gtx layout.Context) layout.Dimensions {
		c := implPool.Get().(*winImpl)
		c.th = parent.theme()
		c.win = parent.rawWindow()
		c.logo = parent.appLogo()
		c.widgets = c.widgets[:0]
		fn(c)
		dims := c.run(gtx)
		implPool.Put(c)
		return dims
	}
}

// color helpers

// RGB builds an opaque color from a 24-bit hex value.
func RGB(hex uint32) color.NRGBA {
	return color.NRGBA{R: uint8(hex >> 16), G: uint8(hex >> 8), B: uint8(hex), A: 0xff}
}

// RGBA builds a color with explicit alpha (0–255 each).
func RGBA(r, g, b, a uint8) color.NRGBA { return color.NRGBA{R: r, G: g, B: b, A: a} }

// HexColor parses a CSS hex string. Accepts "#rrggbb", "rrggbb", "#rgb", "#rrggbbaa".
func HexColor(code string) color.NRGBA { return parseHex(code) }
// Code explanation credit: Robert Carpenter
