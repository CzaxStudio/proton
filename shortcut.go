package proton

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// ShortcutHint draws a small keyboard shortcut badge, e.g. "Ctrl+S".
// Place it next to a button or menu label to show the hotkey.
//
//	proton.Row(win,
//	    func(win *proton.Win) { proton.Label(win, "Save") },
//	    func(win *proton.Win) { proton.Gap(win, 8) },
//	    func(win *proton.Win) { proton.ShortcutHint(win, "Ctrl+S") },
//	)
func ShortcutHint(win *Win, keys string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		bg := win.th.Palette.Fg
		bg.A = 22
		border := win.th.Palette.Fg
		border.A = 50

		lbl := material.Caption(win.th, keys)
		lbl.Color.A = 160

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				h := gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(3))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				paint.FillShape(gtx.Ops, border,
					clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(2), Bottom: unit.Dp(2), Left: unit.Dp(5), Right: unit.Dp(5)}.
					Layout(gtx, lbl.Layout)
			}),
		)
	})
}

// ColorSwatch draws a grid of color circles the user can click to select a color.
// colors is the palette to show. state is one Clickable per color.
// Returns the index of the selected color, or -1 if none selected yet.
//
//	palette := []color.NRGBA{
//	    proton.RGB(0xff6b6b),
//	    proton.RGB(0x4ecdc4),
//	    proton.RGB(0x45b7d1),
//	}
//	var swatches [3]proton.Clickable
//	var selectedColor int
//
//	i := proton.ColorSwatch(win, swatches[:], palette, selectedColor, 28)
//	if i >= 0 { selectedColor = i }
func ColorSwatch(win *Win, btns []Clickable, colors []color.NRGBA, selected int, sizeDp float32) int {
	result := selected
	win.add(func(gtx layout.Context) layout.Dimensions {
		sz := gtx.Dp(unit.Dp(sizeDp))
		gap := gtx.Dp(unit.Dp(6))
		total := sz*len(colors) + gap*(len(colors)-1)
		if total > gtx.Constraints.Max.X {
			total = gtx.Constraints.Max.X
		}

		children := make([]layout.FlexChild, len(colors)*2-1)
		for i := range colors {
			i := i
			children[i*2] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if btns[i].Clicked(gtx) {
					result = i
				}
				return btns[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// outer circle (selection ring)
					outerSz := sz
					if i == result {
						outerSz = sz + 4
					}
					center := outerSz / 2
					inner := sz / 2

					if i == result {
						ring := win.th.Palette.Fg
						ring.A = 200
						paint.FillShape(gtx.Ops, ring,
							clip.Ellipse{
								Min: image.Pt(center-inner-2, center-inner-2),
								Max: image.Pt(center+inner+2, center+inner+2),
							}.Op(gtx.Ops))
					}

					paint.FillShape(gtx.Ops, colors[i],
						clip.Ellipse{
							Min: image.Pt(center-inner, center-inner),
							Max: image.Pt(center+inner, center+inner),
						}.Op(gtx.Ops))

					return layout.Dimensions{Size: image.Pt(outerSz, outerSz)}
				})
			})
			if i < len(colors)-1 {
				children[i*2+1] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Width: unit.Dp(6)}.Layout(gtx)
				})
			}
		}

		dims := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
		_ = total
		return dims
	})
	return result
}

// StatusDot draws a small colored circle — useful for online/offline indicators,
// build status, connection state, etc.
//
//	proton.Row(win,
//	    func(win *proton.Win) { proton.StatusDot(win, proton.RGB(0x4ade80), 10) },
//	    func(win *proton.Win) { proton.Gap(win, 6) },
//	    func(win *proton.Win) { proton.Caption(win, "Connected") },
//	)
func StatusDot(win *Win, c color.NRGBA, sizeDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		sz := gtx.Dp(unit.Dp(sizeDp))
		r := sz / 2
		paint.FillShape(gtx.Ops, c,
			clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Op(gtx.Ops))
		_ = r
		return layout.Dimensions{Size: image.Pt(sz, sz)}
	})
}
