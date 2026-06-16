package proton

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Tooltip shows a label when the user hovers over content.
//
//	var hov proton.Clickable
//	proton.Tooltip(win, &hov, "Saves the file", func(win *proton.Win) {
//	    proton.Button(win, &saveBtn, "Save")
//	})
func Tooltip(win *Win, state *Clickable, tip string, content func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		dims := state.Layout(gtx, child(win, content))
		if !state.Hovered() {
			return dims
		}
		lbl := material.Caption(win.th, tip)
		lbl.Color = win.th.Palette.ContrastFg
		stack := op.Offset(image.Pt(0, dims.Size.Y+4)).Push(gtx.Ops)
		tipDims := layout.UniformInset(unit.Dp(6)).Layout(gtx, lbl.Layout)
		w, h := tipDims.Size.X, tipDims.Size.Y
		r := gtx.Dp(unit.Dp(4))
		rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
		paint.FillShape(gtx.Ops, win.th.Palette.ContrastBg, rrect.Op(gtx.Ops))
		layout.UniformInset(unit.Dp(6)).Layout(gtx, lbl.Layout)
		stack.Pop()
		return dims
	})
}
