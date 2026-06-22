// Package proton is a pure-Go GUI library built on Gio.
//
// Proton's public API never exposes Gio types. Every widget function takes
// a proton.Context — an interface — so if Gio's internals change in a
// future version, only this package's implementation needs updating.
// User code written against proton.Context keeps compiling unchanged.
package proton

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ----- App -----

// App is the top-level application handle. Create one with New().
type App struct {
	theme      *material.Theme
	windows    []*winDef
	bgColor    *color.NRGBA
	bgGradient *gradient
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

// Window registers a window. draw runs every frame with a Context.
// Nothing opens until Run() is called.
//
//	a.Window("Hello", 480, 300, func(ctx proton.Context) {
//	    proton.Label(ctx, "Hello!")
//	})
func (a *App) Window(title string, width, height int, draw func(Context)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, draw: draw})
}

// WindowEx is like Window but accepts extra window options such as
// proton.Fullscreen() or proton.Maximized().
func (a *App) WindowEx(title string, width, height int, opts []WindowOption, draw func(Context)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, extra: opts, draw: draw})
}

// Run opens all registered windows and blocks until they are all closed.
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

// ----- background color -----

type gradient struct {
	from, to color.NRGBA
	dir      string
}

// SetBackground sets a solid window background color.
//
//	a.SetBackground(proton.RGB(0x1a1b26))
func (a *App) SetBackground(c color.NRGBA) {
	a.bgColor = &c
	a.bgGradient = nil
}

// SetBackgroundCode sets the background using a CSS hex string.
// Accepts "#rrggbb", "rrggbb", "#rgb", "rgb".
//
//	a.SetBackgroundCode("#1a1b26")
func (a *App) SetBackgroundCode(code string) {
	c := parseHex(code)
	a.bgColor = &c
	a.bgGradient = nil
}

// SetBackgroundRGB sets the background from r, g, b values (0–255 each).
//
//	a.SetBackgroundRGB(26, 27, 38)
func (a *App) SetBackgroundRGB(r, g, b uint8) {
	c := color.NRGBA{R: r, G: g, B: b, A: 255}
	a.bgColor = &c
	a.bgGradient = nil
}

// SetBackgroundGradient sets a two-color linear gradient background.
// from and to are hex strings. dir is "horizontal", "vertical",
// "diagonal", or "radial".
//
//	a.SetBackgroundGradient("#1a1b26", "#2d1b69", "vertical")
func (a *App) SetBackgroundGradient(from, to, dir string) {
	a.bgGradient = &gradient{from: parseHex(from), to: parseHex(to), dir: dir}
	a.bgColor = nil
}

// SetBackgroundRainbow sets an animated full-spectrum rainbow gradient.
// Cycles slowly over time. A fun default for demos and novelty apps.
func (a *App) SetBackgroundRainbow() {
	a.bgGradient = &gradient{dir: "rainbow"}
	a.bgColor = nil
}

// ----- WindowOption -----

// WindowOption configures extra window behavior.
// Build these with the provided constructors, not directly.
type WindowOption struct {
	apply func(*app.Window)
}

// Fullscreen starts the window in fullscreen mode.
func Fullscreen() WindowOption {
	return WindowOption{apply: func(w *app.Window) { w.Option(app.Fullscreen.Option()) }}
}

// Maximized starts the window maximized.
func Maximized() WindowOption {
	return WindowOption{apply: func(w *app.Window) { w.Option(app.Maximized.Option()) }}
}

// ----- frame loop -----

func runWin(a *App, def *winDef) {
	w := new(app.Window)
	w.Option(app.Title(def.title))
	w.Option(app.Size(unit.Dp(float32(def.w)), unit.Dp(float32(def.h))))
	for _, o := range def.extra {
		o.apply(w)
	}

	var rootList widget.List
	rootList.Axis = layout.Vertical

	var ops op.Ops
	var frame int
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			ops.Reset()
			gtx := app.NewContext(&ops, e)
			frame++

			drawBackground(gtx, a, frame)

			layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				c := &winImpl{th: a.theme, win: w}
				def.draw(c)
				return material.List(a.theme, &rootList).Layout(gtx,
					len(c.widgets),
					func(gtx layout.Context, i int) layout.Dimensions {
						return c.widgets[i](gtx)
					},
				)
			})

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return
		}
	}
}

func drawBackground(gtx layout.Context, a *App, frame int) {
	switch {
	case a.bgGradient != nil && a.bgGradient.dir == "rainbow":
		drawRainbow(gtx, frame)
	case a.bgGradient != nil:
		drawGradient(gtx, a.bgGradient)
	case a.bgColor != nil:
		paint.Fill(gtx.Ops, *a.bgColor)
	default:
		paint.Fill(gtx.Ops, a.theme.Palette.Bg)
	}
}

// drawGradient fills the window with a linear blend between two colors.
func drawGradient(gtx layout.Context, g *gradient) {
	w := gtx.Constraints.Max.X
	h := gtx.Constraints.Max.Y
	const steps = 64

	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps)
		t2 := float32(i+1) / float32(steps)
		c := lerpColor(g.from, g.to, t)

		var rect image.Rectangle
		switch g.dir {
		case "horizontal":
			rect = image.Rect(int(float32(w)*t), 0, int(float32(w)*t2), h)
		case "diagonal":
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

// drawRainbow fills the window with a slowly shifting full-spectrum gradient.
func drawRainbow(gtx layout.Context, frame int) {
	w := gtx.Constraints.Max.X
	h := gtx.Constraints.Max.Y
	const steps = 48
	shift := float32(frame%600) / 600.0

	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps)
		hue := t + shift
		hue -= float32(int(hue)) // wrap to 0..1
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

// hsvColor converts hue/saturation/value (all 0..1) to RGB.
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

// ----- Context -----

// Context is passed to every draw function and layout callback.
// It is Proton's only public handle into the rendering system —
// no Gio types are ever exposed through it. This is what makes
// Proton immune to breaking changes in Gio's internal API: as long
// as this interface's contract stays the same, your code keeps working
// across any future Gio version bump.
//
//	a.Window("App", 480, 300, func(ctx proton.Context) {
//	    proton.Label(ctx, "Hello")
//	    if proton.Button(ctx, &btn, "Click") {
//	        // handle click
//	    }
//	})
type Context interface {
	// Invalidate requests a redraw on the next frame.
	// Call after changing state from a goroutine.
	Invalidate()

	// internal — unexported so no implementation details leak publicly
	add(fn gioWidget)
	theme() *material.Theme
	rawWindow() *app.Window
}

// gioWidget is the internal draw closure type. Never exported.
type gioWidget = func(gtx layout.Context) layout.Dimensions

// winImpl is the concrete, unexported implementation of Context.
// Renaming, restructuring, or swapping this out never affects user code
// because users only ever interact with the Context interface.
type winImpl struct {
	th      *material.Theme
	win     *app.Window
	widgets []gioWidget
}

func (c *winImpl) Invalidate()                { c.win.Invalidate() }
func (c *winImpl) add(fn gioWidget)            { c.widgets = append(c.widgets, fn) }
func (c *winImpl) theme() *material.Theme      { return c.th }
func (c *winImpl) rawWindow() *app.Window      { return c.win }

// run lays out all collected widgets as a vertical Flex.
func (c *winImpl) run(gtx layout.Context) layout.Dimensions {
	if len(c.widgets) == 0 {
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}
	children := make([]layout.FlexChild, len(c.widgets))
	for i, fn := range c.widgets {
		fn := fn
		children[i] = layout.Rigid(fn)
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

// child creates a nested Context, runs fn on it to collect widgets, and
// returns a gioWidget that lays them out during Gio's live layout pass.
// Every layout helper (Row, Column, Pad, Split, ...) uses this.
func child(parent Context, fn func(Context)) gioWidget {
	return func(gtx layout.Context) layout.Dimensions {
		c := &winImpl{th: parent.theme(), win: parent.rawWindow()}
		fn(c)
		return c.run(gtx)
	}
}

// ----- color helpers -----

// RGB builds an opaque color from a 24-bit hex value.
//
//	proton.RGB(0xff6b6b)
func RGB(hex uint32) color.NRGBA {
	return color.NRGBA{R: uint8(hex >> 16), G: uint8(hex >> 8), B: uint8(hex), A: 0xff}
}

// RGBA builds a color with explicit alpha, all values 0–255.
func RGBA(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{R: r, G: g, B: b, A: a}
}

// HexColor parses a CSS hex string into a color.
// Accepts "#rrggbb", "rrggbb", "#rgb", "rgb", "#rrggbbaa".
//
//	c := proton.HexColor("#ff6b6b")
func HexColor(code string) color.NRGBA {
	return parseHex(code)
}
