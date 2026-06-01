package proton

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
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

type App struct {
	theme   *material.Theme
	windows []*winDef
	logo    ImageOp
}

type winDef struct {
	title string
	w, h  int
	opts  []app.Option
	draw  func(*Win)
}

func New(name string) *App {
	a := &App{theme: material.NewTheme()}
	return a
}

func (a *App) Theme() *material.Theme { return a.theme }

// SetLogo loads an image once and caches it for all windows.
func (a *App) SetLogo(data []byte) {
	img, err := LoadImageBytes(data)
	if err != nil {
		fmt.Printf("proton: failed to load logo: %v\n", err)
		return
	}
	a.logo = img
}

func (a *App) Window(title string, width, height int, draw func(*Win)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, draw: draw})
}

func (a *App) WindowEx(title string, width, height int, opts []app.Option, draw func(*Win)) {
	a.windows = append(a.windows, &winDef{title: title, w: width, h: height, opts: opts, draw: draw})
}

func (a *App) Run() {
	if len(a.windows) == 0 {
		log.Fatal("proton: no windows registered")
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

func runWin(a *App, def *winDef) {
	w := new(app.Window)
	w.Option(app.Title(def.title))
	w.Option(app.Size(unit.Dp(float32(def.w)), unit.Dp(float32(def.h))))
	for _, o := range def.opts {
		w.Option(o)
	}

	// rootList is the implicit vertical scroller for the top-level draw function.
	var rootList widget.List
	rootList.Axis = layout.Vertical

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			ops.Reset()
			gtx := app.NewContext(&ops, e)

			// Fill background
			paint.FillShape(gtx.Ops, a.theme.Palette.Bg, clip.Rect{Max: gtx.Constraints.Max}.Op())

			layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				win := &Win{th: a.theme, raw: w, gtx: gtx, logo: a.logo}
				def.draw(win)

				return material.List(a.theme, &rootList).Layout(gtx,
					len(win.widgets),
					func(gtx layout.Context, i int) layout.Dimensions {
						return win.widgets[i](gtx)
					},
				)
			})

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return
		}
	}
}

// Win collects widget draw functions.
type Win struct {
	th      *material.Theme
	raw     *app.Window
	gtx     layout.Context
	logo    ImageOp
	widgets []func(gtx layout.Context) layout.Dimensions
}

// add queues a widget. Called by every widget function (Label, Button, etc).
func (w *Win) add(fn func(gtx layout.Context) layout.Dimensions) {
	w.widgets = append(w.widgets, fn)
}

// run executes all queued widgets through a vertical Flex against the given gtx.
// Used by layout helpers (Column, Row, Split etc) to lay out nested content.
func (w *Win) run(gtx layout.Context) layout.Dimensions {
	if len(w.widgets) == 0 {
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}
	children := make([]layout.FlexChild, len(w.widgets))
	for i, fn := range w.widgets {
		fn := fn
		children[i] = layout.Rigid(fn)
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

// child creates a fresh Win for collecting a nested group of widgets,
// then runs them through a Flex when Gio calls the returned layout.Widget.
// This is what Column, Row, Pad etc. use for their content callbacks.
func child(parent *Win, fn func(*Win)) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		w := &Win{th: parent.th, raw: parent.raw, gtx: gtx}
		fn(w)
		return w.run(gtx)
	}
}

func (w *Win) Invalidate()            { w.raw.Invalidate() }
func (w *Win) Theme() *material.Theme { return w.th }

func RGB(hex uint32) color.NRGBA {
	return color.NRGBA{R: uint8(hex >> 16), G: uint8(hex >> 8), B: uint8(hex), A: 0xff}
}

func RGBA(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{R: r, G: g, B: b, A: a}
}
