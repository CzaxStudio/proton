package proton

import (
	"image"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ContextMenuState tracks whether the menu is open and the click positions.
//
//	type UI struct {
//	    menu     proton.ContextMenuState
//	    menuTag  proton.FrameTag
//	}
type ContextMenuState struct {
	Open bool
	x, y int
	btns []widget.Clickable
}

// ContextMenuItem is one entry in a context menu.
type ContextMenuItem struct {
	Label    string
	Disabled bool
}

// ContextMenu draws a right-click context menu over the target area.
// items defines the menu entries. The returned int is the index of the
// clicked item, or -1 if nothing was selected this frame.
//
//	items := []proton.ContextMenuItem{
//	    {Label: "Copy"},
//	    {Label: "Paste"},
//	    {Label: "Delete", Disabled: false},
//	}
//
//	i := proton.ContextMenu(win, &u.menu, &u.menuTag, items, func(win *proton.Win) {
//	    proton.Label(win, "right-click me")
//	})
//	if i >= 0 {
//	    fmt.Println("selected:", items[i].Label)
//	}
func ContextMenu(win Context, state *ContextMenuState, tag *FrameTag, items []ContextMenuItem, content func(Context)) int {
	for len(state.btns) < len(items) {
		state.btns = append(state.btns, widget.Clickable{})
	}

	selected := -1

	win.add(func(gtx layout.Context) layout.Dimensions {
		// register for pointer events on the content area
		event.Op(gtx.Ops, tag)

		// check for right-click (secondary button press)
		for {
			ev, ok := gtx.Source.Event(pointer.Filter{
				Target: tag,
				Kinds:  pointer.Press,
			})
			if !ok {
				break
			}
			if pe, ok := ev.(pointer.Event); ok && pe.Buttons.Contain(pointer.ButtonSecondary) {
				state.Open = true
				state.x = int(pe.Position.X)
				state.y = int(pe.Position.Y)
			}
			// primary click outside menu closes it
			if pe, ok := ev.(pointer.Event); ok && pe.Buttons.Contain(pointer.ButtonPrimary) && state.Open {
				state.Open = false
			}
		}

		// check menu item clicks
		for i := range items {
			if !items[i].Disabled && state.btns[i].Clicked(gtx) {
				selected = i
				state.Open = false
			}
		}

		// draw the content area
		dims := child(win, content)(gtx)

		if !state.Open {
			return dims
		}

		// draw menu as overlay at click position
		itemH := gtx.Dp(unit.Dp(32))
		menuW := gtx.Dp(unit.Dp(160))
		menuH := itemH * len(items)

		// clamp position so menu stays on screen
		mx := state.x
		my := state.y
		if mx+menuW > gtx.Constraints.Max.X {
			mx = gtx.Constraints.Max.X - menuW
		}
		if my+menuH > gtx.Constraints.Max.Y {
			my = gtx.Constraints.Max.Y - menuH
		}

		macro := op.Record(gtx.Ops)

		bg := win.theme().Palette.Bg
		bg.R += 20
		bg.G += 20
		bg.B += 20
		shadow := win.theme().Palette.Fg
		shadow.A = 40

		r := gtx.Dp(unit.Dp(5))
		menuRect := clip.RRect{
			Rect: image.Rect(0, 0, menuW, menuH+gtx.Dp(unit.Dp(4))),
			NW: r, NE: r, SE: r, SW: r,
		}
		// shadow
		shadowRect := clip.RRect{
			Rect: image.Rect(2, 3, menuW+2, menuH+gtx.Dp(unit.Dp(4))+3),
			NW: r, NE: r, SE: r, SW: r,
		}
		paint.FillShape(gtx.Ops, shadow, shadowRect.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, bg, menuRect.Op(gtx.Ops))

		menuGtx := gtx
		menuGtx.Constraints = layout.Exact(image.Pt(menuW, menuH))

		rowChildren := make([]layout.FlexChild, len(items))
		for i := range items {
			i := i
			rowChildren[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return state.btns[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					rowBg := bg
					if !items[i].Disabled && state.btns[i].Hovered() {
						rowBg = win.theme().Palette.ContrastBg
						rowBg.A = 60
					}
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							w := menuW
							paint.FillShape(gtx.Ops, rowBg,
								clip.Rect{Max: image.Pt(w, itemH)}.Op())
							return layout.Dimensions{Size: image.Pt(w, itemH)}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Body2(win.theme(), items[i].Label)
								if items[i].Disabled {
									lbl.Color.A = 90
								}
								return lbl.Layout(gtx)
							})
						}),
					)
				})
			})
		}
		layout.Flex{Axis: layout.Vertical}.Layout(menuGtx, rowChildren...)

		call := macro.Stop()
		t := op.Offset(image.Pt(mx, my)).Push(gtx.Ops)
		call.Add(gtx.Ops)
		t.Pop()

		return dims
	})

	return selected
}
