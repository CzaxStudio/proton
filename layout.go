package proton

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Column stacks widgets vertically.
func Column(win Context, widgets ...func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	})
}

// Row places widgets side by side.
func Row(win Context, widgets ...func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
}

// RowSpread is like Row but puts leftover space between children.
func RowSpread(win Context, widgets ...func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx, children...)
	})
}

// RowEnd pushes all children to the right edge.
func RowEnd(win Context, widgets ...func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceStart}.Layout(gtx, children...)
	})
}

// GrowRow is a horizontal row with explicit stretch control.
// Use FixedItem for natural-size children, GrowItem for stretchy ones.
//
//	proton.GrowRow(win,
//	    proton.FixedItem(win, func(win proton.Context) { proton.Label(win, "Name:") }),
//	    proton.GrowItem(win, func(win proton.Context) { proton.Input(win, &e, "") }),
//	    proton.FixedItem(win, func(win proton.Context) { proton.Button(win, &b, "Go") }),
//	)
func GrowRow(win Context, children ...FlexItem) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		fc := make([]layout.FlexChild, len(children))
		for i, c := range children {
			fc[i] = c.flexChild
		}
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, fc...)
	})
}

// GrowColumn is a vertical column with explicit stretch control.
func GrowColumn(win Context, children ...FlexItem) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		fc := make([]layout.FlexChild, len(children))
		for i, c := range children {
			fc[i] = c.flexChild
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, fc...)
	})
}

// FlexItem is one child of a GrowRow or GrowColumn.
// Build one with GrowItem or FixedItem — don't construct it directly.
// This wraps Gio's flex child type so it never appears in Proton's public API.
type FlexItem struct {
	flexChild layout.FlexChild
}

// GrowItem makes a child fill remaining space. Use inside GrowRow/GrowColumn.
func GrowItem(win Context, fn func(Context)) FlexItem {
	return FlexItem{flexChild: layout.Flexed(1, child(win, fn))}
}

// FixedItem makes a child take only as much space as it needs. Use inside GrowRow/GrowColumn.
func FixedItem(win Context, fn func(Context)) FlexItem {
	return FlexItem{flexChild: layout.Rigid(child(win, fn))}
}

// Split gives left and right a fraction of the width. leftFraction is 0.0–1.0.
func Split(win Context, leftFraction float32, left, right func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(leftFraction, child(win, left)),
			layout.Flexed(1-leftFraction, child(win, right)),
		)
	})
}

// HSplit splits vertically. topFraction is 0.0–1.0.
func HSplit(win Context, topFraction float32, top, bottom func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(topFraction, child(win, top)),
			layout.Flexed(1-topFraction, child(win, bottom)),
		)
	})
}

// Center places a widget in the center of available space.
func Center(win Context, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, child(win, fn))
	})
}

// Pad adds uniform padding around a widget.
func Pad(win Context, dp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(dp)).Layout(gtx, child(win, fn))
	})
}

// PadH adds left+right padding.
func PadH(win Context, dp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Left: unit.Dp(dp), Right: unit.Dp(dp)}.Layout(gtx, child(win, fn))
	})
}

// PadV adds top+bottom padding.
func PadV(win Context, dp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: unit.Dp(dp), Bottom: unit.Dp(dp)}.Layout(gtx, child(win, fn))
	})
}

// PadSides gives per-edge padding control.
func PadSides(win Context, top, right, bottom, left float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top:    unit.Dp(top),
			Right:  unit.Dp(right),
			Bottom: unit.Dp(bottom),
			Left:   unit.Dp(left),
		}.Layout(gtx, child(win, fn))
	})
}

// Gap inserts a fixed blank space.
func Gap(win Context, dp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Width: unit.Dp(dp), Height: unit.Dp(dp)}.Layout(gtx)
	})
}
