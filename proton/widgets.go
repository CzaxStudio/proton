package proton

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// clickResults, boolResults, enumResults store the click/change state from
// the previous frame for each widget state pointer. This lets Button(),
// Checkbox() etc. return correct values even though their layout closures
// run after the draw function returns.
var clickResults = map[*Clickable]bool{}
var boolResults = map[*Bool]bool{}
var enumResults = map[*Enum]bool{}

// ----- text -----

func Label(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Body1(win.theme(), text).Layout(gtx)
	})
}

func H1(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H1(win.theme(), text).Layout(gtx) })
}
func H2(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H2(win.theme(), text).Layout(gtx) })
}
func H3(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H3(win.theme(), text).Layout(gtx) })
}
func H4(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H4(win.theme(), text).Layout(gtx) })
}
func H5(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H5(win.theme(), text).Layout(gtx) })
}
func H6(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H6(win.theme(), text).Layout(gtx) })
}
func Body2(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Body2(win.theme(), text).Layout(gtx) })
}
func Caption(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Caption(win.theme(), text).Layout(gtx) })
}

func Text(win Context, s string, size float32, c color.NRGBA, bold bool) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(win.theme(), unit.Sp(size), s)
		if c != (color.NRGBA{}) {
			lbl.Color = c
		}
		if bold {
			lbl.Font.Weight = font.Bold
		}
		return lbl.Layout(gtx)
	})
}

// ----- buttons -----

// Button draws a filled button. Returns true if it was clicked.
// Click state is read from the previous frame's event processing —
// at 60fps this one-frame latency is imperceptible.
func Button(win Context, state *Clickable, label string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return material.Button(win.theme(), state, label).Layout(gtx)
	})
	return result
}

func OutlineButton(win Context, state *Clickable, label string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		btn := material.Button(win.theme(), state, label)
		btn.Background = color.NRGBA{}
		btn.Color = win.theme().Palette.ContrastBg
		return btn.Layout(gtx)
	})
	return result
}

func IconButton(win Context, state *Clickable, icon *Icon, desc string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return material.IconButton(win.theme(), state, icon, desc).Layout(gtx)
	})
	return result
}

// Tappable makes any content clickable. Returns true if clicked.
func Tappable(win Context, state *Clickable, content func(Context)) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return state.Layout(gtx, child(win, content))
	})
	return result
}

// ----- inputs -----

func Input(win Context, state *Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.SingleLine = true
		return material.Editor(win.theme(), state, hint).Layout(gtx)
	})
}

func TextArea(win Context, state *Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Editor(win.theme(), state, hint).Layout(gtx)
	})
}

// ----- toggle / checkbox / radio -----

func Toggle(win Context, state *Bool, label string) bool {
	result := boolResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		boolResults[state] = state.Update(gtx)
		return material.Switch(win.theme(), state, label).Layout(gtx)
	})
	return result
}

func Checkbox(win Context, state *Bool, label string) bool {
	result := boolResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		boolResults[state] = state.Update(gtx)
		return material.CheckBox(win.theme(), state, label).Layout(gtx)
	})
	return result
}

func RadioButton(win Context, group *Enum, key, label string) bool {
	result := enumResults[group]
	win.add(func(gtx layout.Context) layout.Dimensions {
		enumResults[group] = group.Update(gtx)
		return material.RadioButton(win.theme(), group, key, label).Layout(gtx)
	})
	return result
}

// ----- slider / progress -----

func Slider(win Context, state *Float) float32 {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Slider(win.theme(), state).Layout(gtx)
	})
	return state.Value
}

func ProgressBar(win Context, progress float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.ProgressBar(win.theme(), progress).Layout(gtx)
	})
}

// ----- list / scroll -----

func List(win Context, state *Scrollable, length int, draw func(Context, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.theme(), state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w Context) { draw(w, i) })(gtx)
		})
	})
}

func HList(win Context, state *Scrollable, length int, draw func(Context, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Horizontal
		return material.List(win.theme(), state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w Context) { draw(w, i) })(gtx)
		})
	})
}

func Scroll(win Context, state *Scrollable, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.theme(), state).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
			return child(win, content)(gtx)
		})
	})
}

// ----- visual -----

func Divider(win Context) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		h := gtx.Dp(unit.Dp(1))
		w := gtx.Constraints.Max.X
		paint.FillShape(gtx.Ops, win.theme().Palette.Fg, clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func Rect(win Context, c color.NRGBA, widthDp, heightDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func RoundRect(win Context, c color.NRGBA, widthDp, heightDp, radiusDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		r := gtx.Dp(unit.Dp(radiusDp))
		rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
		paint.FillShape(gtx.Ops, c, rrect.Op(gtx.Ops))
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func Card(win Context, bg color.NRGBA, cornerDp, padDp float32, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w, h := gtx.Constraints.Min.X, gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(cornerDp))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(padDp)).Layout(gtx, child(win, content))
			}),
		)
	})
}

func Badge(win Context, bg, fg color.NRGBA, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w, h := gtx.Constraints.Min.X, gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(12))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Caption(win.theme(), text)
				lbl.Color = fg
				return layout.UniformInset(unit.Dp(4)).Layout(gtx, lbl.Layout)
			}),
		)
	})
}

// ----- size helpers -----

func MinSize(win Context, widthDp, heightDp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if widthDp > 0 {
			if m := gtx.Dp(unit.Dp(widthDp)); gtx.Constraints.Min.X < m {
				gtx.Constraints.Min.X = m
			}
		}
		if heightDp > 0 {
			if m := gtx.Dp(unit.Dp(heightDp)); gtx.Constraints.Min.Y < m {
				gtx.Constraints.Min.Y = m
			}
		}
		return child(win, fn)(gtx)
	})
}

func MaxWidth(win Context, widthDp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if m := gtx.Dp(unit.Dp(widthDp)); gtx.Constraints.Max.X > m {
			gtx.Constraints.Max.X = m
		}
		return child(win, fn)(gtx)
	})
}

// ----- image -----

type ImageOp struct {
	op paint.ImageOp
	sz image.Point
}

func LoadImage(path string) (ImageOp, error) {
	f, err := os.Open(path)
	if err != nil {
		return ImageOp{}, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return ImageOp{}, err
	}
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		bounds := img.Bounds()
		nrgba = image.NewNRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				nrgba.Set(x, y, img.At(x, y))
			}
		}
	}
	return ImageOp{op: paint.NewImageOp(nrgba), sz: nrgba.Bounds().Size()}, nil
}

func Image(win Context, img ImageOp, widthDp, heightDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w, h := img.sz.X, img.sz.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		sz := image.Pt(w, h)
		stack := clip.Rect{Max: sz}.Push(gtx.Ops)
		img.op.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		stack.Pop()
		return layout.Dimensions{Size: sz}
	})
}
