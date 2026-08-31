package proton

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Grid arranges widgets in a fixed-column grid.
//
//	proton.Grid(win, 3, 8,
//	    func(win proton.Context) { proton.Label(win, "one") },
//	    func(win proton.Context) { proton.Label(win, "two") },
//	    func(win proton.Context) { proton.Label(win, "three") },
//	)
func Grid(win Context, cols int, gapDp float32, cells ...func(Context)) {
	if cols <= 0 {
		cols = 1
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		gap := gtx.Dp(unit.Dp(gapDp))
		cellW := (gtx.Constraints.Max.X - gap*(cols-1)) / cols

		var rows []layout.FlexChild
		for i := 0; i < len(cells); i += cols {
			end := i + cols
			if end > len(cells) {
				end = len(cells)
			}
			slice := cells[i:end]
			cw := cellW

			rows = append(rows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				rowChildren := make([]layout.FlexChild, 0, len(slice)*2)
				for j := range slice {
					j := j
					cell := slice[j]
					rowChildren = append(rowChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.X = cw
						gtx.Constraints.Max.X = cw
						return child(win, cell)(gtx)
					}))
					if j < len(slice)-1 && gap > 0 {
						rowChildren = append(rowChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: unit.Dp(gapDp)}.Layout(gtx)
						}))
					}
				}
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, rowChildren...)
			}))

			if end < len(cells) && gap > 0 {
				rows = append(rows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: unit.Dp(gapDp)}.Layout(gtx)
				}))
			}
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
	})
}
