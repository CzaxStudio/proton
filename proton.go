package proton

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type App struct {
	theme   *material.Theme
	windows []*winDef
}

type winDef struct {
	title string
	w, h  int
	opts  []app.Option
	draw  func(*Win)
}

func New(name string) *App {
	return &App{theme: material.NewTheme()}
}

func (a *App) Theme() *material.Theme { return a.theme }

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
		go runWin(a.theme, w)
	}
	go func() {
		runWin(a.theme, a.windows[len(a.windows)-1])
		os.Exit(0)
	}()
	app.Main()
}

func runWin(th *material.Theme, def *winDef) {
	w := new(app.Window)
	w.Option(app.Title(def.title))
	w.Option(app.Size(unit.Dp(float32(def.w)), unit.Dp(float32(def.h))))
	for _, o := range def.opts {
		w.Option(o)
	}

	var rootList widget.List
	rootList.Axis = layout.Vertical

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			ops.Reset()
			gtx := app.NewContext(&ops, e)

			layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// We run the draw function once to collect widget fns,
				// then pass them to material.List which calls each fn
				// with a correctly scoped gtx — so state.Clicked(gtx),
				// state.Update(gtx) etc. all fire with the real gtx.
				//
				// The draw function itself does NOT return click results
				// at this stage — all results are captured inside the
				// closures via pointer variables and read back on the
				// NEXT call to def.draw (next frame). This is standard
				// immediate-mode GUI behaviour: clicks fire, state
				// updates, next frame sees the new state.
				win := &Win{th: th, raw: w}
				def.draw(win)

				return material.List(th, &rootList).Layout(gtx,
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
// They run inside Gio's layout engine so each gets a correctly scoped gtx.
type Win struct {
	th      *material.Theme
	raw     *app.Window
	widgets []func(gtx layout.Context) layout.Dimensions
}

func (w *Win) add(fn func(gtx layout.Context) layout.Dimensions) {
	w.widgets = append(w.widgets, fn)
}

// run executes all collected widgets as a vertical Flex.
// Called by child() for nested layouts.
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

// child creates a nested Win, runs fn on it, and returns a layout.Widget
// that flushes all collected widgets through a Flex with the live gtx.
// This is what makes buttons inside Pad/Row/Column work correctly —
// the fn runs during Gio's layout pass, not before it.
func child(parent *Win, fn func(*Win)) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		w := &Win{th: parent.th, raw: parent.raw}
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
