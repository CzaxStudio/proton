package proton

import (
	"image"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// ResizeSplitState tracks the drag position for a resizable split pane.
//
//	type UI struct {
//	    split proton.ResizeSplitState
//	}
type ResizeSplitState struct {
	Fraction float32
	dragging bool
	tag      struct{}
}

// ResizeSplit is like Split but the user can drag the divider to resize the panes.
//
//	proton.ResizeSplit(win, &u.split, 0.35,
//	    func(win proton.Context) { drawSidebar(win) },
//	    func(win proton.Context) { drawContent(win) },
//	)
func ResizeSplit(win Context, state *ResizeSplitState, defaultFraction float32, left, right func(Context)) {
	if state.Fraction == 0 {
		state.Fraction = defaultFraction
	}

	win.add(func(gtx layout.Context) layout.Dimensions {
		handleW := gtx.Dp(unit.Dp(5))
		totalW := gtx.Constraints.Max.X
		totalH := gtx.Constraints.Max.Y

		// register for pointer events on the handle area
		event.Op(gtx.Ops, &state.tag)
		for {
			ev, ok := gtx.Source.Event(pointer.Filter{
				Target: &state.tag,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release,
			})
			if !ok {
				break
			}
			if pe, ok := ev.(pointer.Event); ok {
				switch pe.Kind {
				case pointer.Press:
					state.dragging = true
				case pointer.Drag:
					if state.dragging && totalW > 0 {
						f := float32(pe.Position.X) / float32(totalW)
						if f < 0.1 {
							f = 0.1
						}
						if f > 0.9 {
							f = 0.9
						}
						state.Fraction = f
					}
				case pointer.Release:
					state.dragging = false
				}
			}
		}

		leftW := int(float32(totalW-handleW) * state.Fraction)
		rightW := totalW - handleW - leftW

		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(leftW, totalH))
				return child(win, left)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(handleW, totalH))

				// register hit area
				area := clip.Rect{Max: image.Pt(handleW, totalH)}.Push(gtx.Ops)
				event.Op(gtx.Ops, &state.tag)
				area.Pop()

				c := win.theme().Palette.Fg
				c.A = 50
				if state.dragging {
					c = win.theme().Palette.ContrastBg
					c.A = 180
				}
				lineW := gtx.Dp(unit.Dp(1))
				xOffset := (handleW - lineW) / 2
				paint.FillShape(gtx.Ops, c,
					clip.Rect{
						Min: image.Pt(xOffset, 0),
						Max: image.Pt(xOffset+lineW, totalH),
					}.Op())
				return layout.Dimensions{Size: image.Pt(handleW, totalH)}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(rightW, totalH))
				return child(win, right)(gtx)
			}),
		)
	})
}

// ResizeHSplit is the vertical version of ResizeSplit — top and bottom panes
// with a draggable horizontal divider.
func ResizeHSplit(win Context, state *ResizeSplitState, defaultFraction float32, top, bottom func(Context)) {
	if state.Fraction == 0 {
		state.Fraction = defaultFraction
	}

	win.add(func(gtx layout.Context) layout.Dimensions {
		handleH := gtx.Dp(unit.Dp(5))
		totalW := gtx.Constraints.Max.X
		totalH := gtx.Constraints.Max.Y

		event.Op(gtx.Ops, &state.tag)
		for {
			ev, ok := gtx.Source.Event(pointer.Filter{
				Target: &state.tag,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release,
			})
			if !ok {
				break
			}
			if pe, ok := ev.(pointer.Event); ok {
				switch pe.Kind {
				case pointer.Press:
					state.dragging = true
				case pointer.Drag:
					if state.dragging && totalH > 0 {
						f := float32(pe.Position.Y) / float32(totalH)
						if f < 0.1 {
							f = 0.1
						}
						if f > 0.9 {
							f = 0.9
						}
						state.Fraction = f
					}
				case pointer.Release:
					state.dragging = false
				}
			}
		}

		topH := int(float32(totalH-handleH) * state.Fraction)
		botH := totalH - handleH - topH

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(totalW, topH))
				return child(win, top)(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(totalW, handleH))

				area := clip.Rect{Max: image.Pt(totalW, handleH)}.Push(gtx.Ops)
				event.Op(gtx.Ops, &state.tag)
				area.Pop()

				c := win.theme().Palette.Fg
				c.A = 50
				if state.dragging {
					c = win.theme().Palette.ContrastBg
					c.A = 180
				}
				lineH := gtx.Dp(unit.Dp(1))
				yOffset := (handleH - lineH) / 2
				paint.FillShape(gtx.Ops, c,
					clip.Rect{
						Min: image.Pt(0, yOffset),
						Max: image.Pt(totalW, yOffset+lineH),
					}.Op())
				return layout.Dimensions{Size: image.Pt(totalW, handleH)}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(image.Pt(totalW, botH))
				return child(win, bottom)(gtx)
			}),
		)
	})
}
