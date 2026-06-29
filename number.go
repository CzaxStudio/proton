package proton

import (
	"fmt"
	"image"
	"math"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// NumberState holds state for a numeric stepper.
//
//	type UI struct {
//	    qty proton.NumberState
//	}
type NumberState struct {
	Value float64
	dec   widget.Clickable
	inc   widget.Clickable
}

// NumberInput draws a − value + stepper row.
// min/max clamp the value, step is the increment per click.
// Returns the current value.
//
//	qty := proton.NumberInput(win, &u.qty, 1, 99, 1)
func NumberInput(win Context, state *NumberState, min, max, step float64) float64 {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if state.dec.Clicked(gtx) {
			state.Value = math.Max(min, state.Value-step)
		}
		if state.inc.Clicked(gtx) {
			state.Value = math.Min(max, state.Value+step)
		}
		if state.Value < min {
			state.Value = min
		}

		btnSz := image.Pt(gtx.Dp(unit.Dp(36)), gtx.Dp(unit.Dp(36)))
		numSz := image.Pt(gtx.Dp(unit.Dp(60)), gtx.Dp(unit.Dp(36)))

		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(btnSz)
				return material.Button(win.theme(), &state.dec, "−").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(numSz)
				return layout.Center.Layout(gtx, material.Body1(win.theme(), numFmt(state.Value, step)).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints = layout.Exact(btnSz)
				return material.Button(win.theme(), &state.inc, "+").Layout(gtx)
			}),
		)
	})
	return state.Value
}

func numFmt(v, step float64) string {
	if step >= 1 {
		return fmt.Sprintf("%d", int(v))
	}
	return fmt.Sprintf("%.2f", v)
}
