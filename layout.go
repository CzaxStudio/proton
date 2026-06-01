package proton

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Column stacks widgets vertically.
func Column(win *Win, widgets ...func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	})
}

// Row places widgets side by side.
func Row(win *Win, widgets ...func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
}

// RowSpread is like Row but puts leftover space between children.
func RowSpread(win *Win, widgets ...func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		children := make([]layout.FlexChild, len(widgets))
		for i, fn := range widgets {
			children[i] = layout.Rigid(child(win, fn))
		}
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx, children...)
	})
}

// RowEnd pushes all children to the right edge.
func RowEnd(win *Win, widgets ...func(*Win)) {
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
//	    proton.FixedItem(win, func(win *proton.Win) { proton.Label(win, "Name:") }),
//	    proton.GrowItem(win, func(win *proton.Win) { proton.Input(win, &e, "") }),
//	    proton.FixedItem(win, func(win *proton.Win) { proton.Button(win, &b, "Go") }),
//	)
func GrowRow(win *Win, children ...layout.FlexChild) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
}

// GrowColumn is a vertical column with explicit stretch control.
func GrowColumn(win *Win, children ...layout.FlexChild) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	})
}

// GrowItem makes a child fill remaining space. Use inside GrowRow/GrowColumn.
func GrowItem(win *Win, fn func(*Win)) layout.FlexChild {
	return layout.Flexed(1, child(win, fn))
}

// FixedItem makes a child take only as much space as it needs. Use inside GrowRow/GrowColumn.
func FixedItem(win *Win, fn func(*Win)) layout.FlexChild {
	return layout.Rigid(child(win, fn))
}

// Split gives left and right a fraction of the width. leftFraction is 0.0–1.0.
func Split(win *Win, leftFraction float32, left, right func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(leftFraction, child(win, left)),
			layout.Flexed(1-leftFraction, child(win, right)),
		)
	})
}

// HSplit splits vertically. topFraction is 0.0–1.0.
func HSplit(win *Win, topFraction float32, top, bottom func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(topFraction, child(win, top)),
			layout.Flexed(1-topFraction, child(win, bottom)),
		)
	})
}

// Center places a widget in the center of available space.
func Center(win *Win, fn func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, child(win, fn))
	})
}

// Pad adds uniform padding around a widget.
func Pad(win *Win, dp float32, fn func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(dp)).Layout(gtx, child(win, fn))
	})
}

// PadH adds left+right padding.
func PadH(win *Win, dp float32, fn func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Left: unit.Dp(dp), Right: unit.Dp(dp)}.Layout(gtx, child(win, fn))
	})
}

// PadV adds top+bottom padding.
func PadV(win *Win, dp float32, fn func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: unit.Dp(dp), Bottom: unit.Dp(dp)}.Layout(gtx, child(win, fn))
	})
}

// PadSides gives per-edge padding control.
func PadSides(win *Win, top, right, bottom, left float32, fn func(*Win)) {
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
func Gap(win *Win, dp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Width: unit.Dp(dp), Height: unit.Dp(dp)}.Layout(gtx)
	})
}

// Sub returns a layout.Widget that runs fn with a fresh Win.
// Use when mixing Proton with raw Gio layout code.
func Sub(win *Win, fn func(*Win)) func(gtx layout.Context) layout.Dimensions {
	return child(win, fn)
}
