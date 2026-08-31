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

// AlertKind controls the visual style of an alert banner.
type AlertKind int

const (
	AlertInfo    AlertKind = iota // blue — informational
	AlertSuccess                  // green — success
	AlertWarning                  // yellow — caution
	AlertError                    // red — something went wrong
)

// Alert draws a colored banner with a message — like an inline notification
// that stays visible until dismissed. Good for form errors, status messages,
// or anything that needs persistent visibility unlike a toast.
//
//	proton.Alert(win, proton.AlertError, "Invalid email address.")
//	proton.Alert(win, proton.AlertSuccess, "Your changes have been saved.")
//	proton.Alert(win, proton.AlertWarning, "This action cannot be undone.")
//	proton.Alert(win, proton.AlertInfo, "The app will restart to apply updates.")
func Alert(win Context, kind AlertKind, message string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		var bg, fg, accent color.NRGBA
		switch kind {
		case AlertInfo:
			bg = color.NRGBA{R: 30, G: 64, B: 175, A: 40}
			accent = color.NRGBA{R: 96, G: 165, B: 250, A: 255}
			fg = color.NRGBA{R: 191, G: 219, B: 254, A: 230}
		case AlertSuccess:
			bg = color.NRGBA{R: 20, G: 83, B: 45, A: 40}
			accent = color.NRGBA{R: 74, G: 222, B: 128, A: 255}
			fg = color.NRGBA{R: 187, G: 247, B: 208, A: 230}
		case AlertWarning:
			bg = color.NRGBA{R: 120, G: 53, B: 15, A: 40}
			accent = color.NRGBA{R: 251, G: 191, B: 36, A: 255}
			fg = color.NRGBA{R: 254, G: 240, B: 138, A: 230}
		case AlertError:
			bg = color.NRGBA{R: 127, G: 29, B: 29, A: 40}
			accent = color.NRGBA{R: 248, G: 113, B: 113, A: 255}
			fg = color.NRGBA{R: 254, G: 202, B: 202, A: 230}
		}

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				if w == 0 {
					w = gtx.Constraints.Max.X
				}
				h := gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(6))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				// left accent bar
				barW := gtx.Dp(unit.Dp(3))
				barRect := clip.RRect{
					Rect: image.Rect(0, 0, barW, h),
					NW:   r, NE: 0, SE: 0, SW: r,
				}
				paint.FillShape(gtx.Ops, accent, barRect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top: unit.Dp(10), Bottom: unit.Dp(10),
					Left: unit.Dp(14), Right: unit.Dp(12),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Body2(win.theme(), message)
					lbl.Color = fg
					return lbl.Layout(gtx)
				})
			}),
		)
	})
}

// AlertDismissable is like Alert but shows a close button.
// visible controls whether it's shown. Set it to false to dismiss.
// Returns true on the frame the close button is clicked.
//
//	type UI struct {
//	    showAlert bool
//	    alertDismiss proton.Clickable
//	}
//
//	if u.showAlert {
//	    if proton.AlertDismissable(win, &u.alertDismiss, proton.AlertInfo, "Click X to dismiss") {
//	        u.showAlert = false
//	    }
//	}
func AlertDismissable(win Context, closeBtn *Clickable, kind AlertKind, message string) bool {
	result := clickResults[closeBtn]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[closeBtn] = closeBtn.Clicked(gtx)

		var bg, fg, accent color.NRGBA
		switch kind {
		case AlertInfo:
			bg = color.NRGBA{R: 30, G: 64, B: 175, A: 40}
			accent = color.NRGBA{R: 96, G: 165, B: 250, A: 255}
			fg = color.NRGBA{R: 191, G: 219, B: 254, A: 230}
		case AlertSuccess:
			bg = color.NRGBA{R: 20, G: 83, B: 45, A: 40}
			accent = color.NRGBA{R: 74, G: 222, B: 128, A: 255}
			fg = color.NRGBA{R: 187, G: 247, B: 208, A: 230}
		case AlertWarning:
			bg = color.NRGBA{R: 120, G: 53, B: 15, A: 40}
			accent = color.NRGBA{R: 251, G: 191, B: 36, A: 255}
			fg = color.NRGBA{R: 254, G: 240, B: 138, A: 230}
		case AlertError:
			bg = color.NRGBA{R: 127, G: 29, B: 29, A: 40}
			accent = color.NRGBA{R: 248, G: 113, B: 113, A: 255}
			fg = color.NRGBA{R: 254, G: 202, B: 202, A: 230}
		}

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				if w == 0 {
					w = gtx.Constraints.Max.X
				}
				h := gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(6))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				barW := gtx.Dp(unit.Dp(3))
				barRect := clip.RRect{
					Rect: image.Rect(0, 0, barW, h),
					NW:   r, NE: 0, SE: 0, SW: r,
				}
				paint.FillShape(gtx.Ops, accent, barRect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top: unit.Dp(10), Bottom: unit.Dp(10),
					Left: unit.Dp(14), Right: unit.Dp(12),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Body2(win.theme(), message)
							lbl.Color = fg
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return closeBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Body2(win.theme(), "×")
								lbl.Color = fg
								return lbl.Layout(gtx)
							})
						}),
					)
				})
			}),
		)
	})
	return result
}
