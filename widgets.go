package proton

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ----- text -----

func Label(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Body1(win.th, text).Layout(gtx)
	})
}

func H1(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H1(win.th, text).Layout(gtx) })
}
func H2(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H2(win.th, text).Layout(gtx) })
}
func H3(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H3(win.th, text).Layout(gtx) })
}
func H4(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H4(win.th, text).Layout(gtx) })
}
func H5(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H5(win.th, text).Layout(gtx) })
}
func H6(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H6(win.th, text).Layout(gtx) })
}
func Body2(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Body2(win.th, text).Layout(gtx) })
}
func Caption(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Caption(win.th, text).Layout(gtx) })
}

func Text(win *Win, s string, size float32, c color.NRGBA, bold bool) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(win.th, unit.Sp(size), s)
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

// Button draws a filled button. Returns true if clicked.
// The result is from the current frame's layout pass — always correct.
func Button(win *Win, state *widget.Clickable, label string) bool {
	clicked := state.Clicked(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Button(win.th, state, label).Layout(gtx)
	})
	return clicked
}

func OutlineButton(win *Win, state *widget.Clickable, label string) bool {
	clicked := state.Clicked(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(win.th, state, label)
		btn.Background = color.NRGBA{}
		btn.Color = win.th.Palette.ContrastBg
		return btn.Layout(gtx)
	})
	return clicked
}

func IconButton(win *Win, state *widget.Clickable, icon *widget.Icon, desc string) bool {
	clicked := state.Clicked(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.IconButton(win.th, state, icon, desc).Layout(gtx)
	})
	return clicked
}

// Tappable makes any content clickable. Returns true if clicked.
func Tappable(win *Win, state *widget.Clickable, content func(*Win)) bool {
	clicked := state.Clicked(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return state.Layout(gtx, child(win, content))
	})
	return clicked
}

// ----- inputs -----

func Input(win *Win, state *widget.Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.SingleLine = true
		return material.Editor(win.th, state, hint).Layout(gtx)
	})
}

func TextArea(win *Win, state *widget.Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Editor(win.th, state, hint).Layout(gtx)
	})
}

// ----- toggle / checkbox / radio -----

func Toggle(win *Win, state *widget.Bool, label string) bool {
	changed := state.Update(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Switch(win.th, state, label).Layout(gtx)
	})
	return changed
}

func Checkbox(win *Win, state *widget.Bool, label string) bool {
	changed := state.Update(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.CheckBox(win.th, state, label).Layout(gtx)
	})
	return changed
}

func RadioButton(win *Win, group *widget.Enum, key, label string) bool {
	changed := group.Update(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.RadioButton(win.th, group, key, label).Layout(gtx)
	})
	return changed
}

// ----- slider / progress -----

func Slider(win *Win, state *widget.Float) float32 {
	state.Update(win.gtx)
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Slider(win.th, state).Layout(gtx)
	})
	return state.Value
}

func ProgressBar(win *Win, progress float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.ProgressBar(win.th, progress).Layout(gtx)
	})
}

// ----- list / scroll -----

func List(win *Win, state *widget.List, length int, draw func(*Win, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.th, state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w *Win) { draw(w, i) })(gtx)
		})
	})
}

func HList(win *Win, state *widget.List, length int, draw func(*Win, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Horizontal
		return material.List(win.th, state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w *Win) { draw(w, i) })(gtx)
		})
	})
}

func Scroll(win *Win, state *widget.List, content func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.th, state).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
			return child(win, content)(gtx)
		})
	})
}

// ----- visual -----

func Divider(win *Win) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		h := gtx.Dp(unit.Dp(1))
		w := gtx.Constraints.Max.X
		paint.FillShape(gtx.Ops, win.th.Palette.Fg, clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func Rect(win *Win, c color.NRGBA, widthDp, heightDp float32) {
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

func RoundRect(win *Win, c color.NRGBA, widthDp, heightDp, radiusDp float32) {
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

func Card(win *Win, bg color.NRGBA, cornerDp, padDp float32, content func(*Win)) {
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

func Badge(win *Win, bg, fg color.NRGBA, text string) {
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
				lbl := material.Caption(win.th, text)
				lbl.Color = fg
				return layout.UniformInset(unit.Dp(4)).Layout(gtx, lbl.Layout)
			}),
		)
	})
}

// ----- size helpers -----

func MinSize(win *Win, widthDp, heightDp float32, fn func(*Win)) {
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

func MaxWidth(win *Win, widthDp float32, fn func(*Win)) {
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

func (i ImageOp) Size() image.Point {
	return i.sz
}

func LoadImage(path string) (ImageOp, error) {
	f, err := os.Open(path)
	if err != nil {
		return ImageOp{}, err
	}
	defer f.Close()
	return decodeImage(f)
}

func LoadImageBytes(data []byte) (ImageOp, error) {
	return decodeImage(bytes.NewReader(data))
}

func decodeImage(r io.Reader) (ImageOp, error) {
	img, _, err := image.Decode(r)
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

func Image(win *Win, img ImageOp, widthDp, heightDp float32) {
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

// Logo draws the application logo if one was set via App.SetLogo().
func Logo(win *Win, widthDp, heightDp float32) {
	if win.logo.Size().X > 0 {
		Image(win, win.logo, widthDp, heightDp)
	}
}
