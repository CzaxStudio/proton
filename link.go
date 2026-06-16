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

// Link draws a clickable underlined text label styled like a hyperlink.
// Returns true if clicked. Handle the click yourself — open a browser,
// navigate in-app, whatever makes sense.
//
//	var lnk proton.Clickable
//	if proton.Link(win, &lnk, "View on GitHub") {
//	    openBrowser("https://github.com/CzaxStudio/proton")
//	}
func Link(win *Win, state *Clickable, text string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)

		c := win.th.Palette.ContrastBg
		if state.Hovered() {
			c.A = 255
		} else {
			c.A = 200
		}

		return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Body1(win.th, text)
			lbl.Color = c

			dims := lbl.Layout(gtx)

			// underline
			uh := gtx.Dp(unit.Dp(1))
			uy := dims.Size.Y - gtx.Dp(unit.Dp(2))
			paint.FillShape(gtx.Ops, c,
				clip.Rect{
					Min: image.Pt(0, uy),
					Max: image.Pt(dims.Size.X, uy+uh),
				}.Op())

			return dims
		})
	})
	return result
}

// LinkSmall is like Link but uses caption-sized text.
func LinkSmall(win *Win, state *Clickable, text string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)

		c := win.th.Palette.ContrastBg
		if state.Hovered() {
			c.A = 255
		} else {
			c.A = 190
		}

		return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Caption(win.th, text)
			lbl.Color = c
			dims := lbl.Layout(gtx)
			uh := gtx.Dp(unit.Dp(1))
			uy := dims.Size.Y - gtx.Dp(unit.Dp(1))
			paint.FillShape(gtx.Ops, c,
				clip.Rect{
					Min: image.Pt(0, uy),
					Max: image.Pt(dims.Size.X, uy+uh),
				}.Op())
			return dims
		})
	})
	return result
}

// CodeBlock draws a monospace text box styled like a code snippet.
// Good for showing commands, file paths, API keys, config values, etc.
//
//	proton.CodeBlock(win, "go get github.com/CzaxStudio/proton")
func CodeBlock(win *Win, code string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		bg := win.th.Palette.Fg
		bg.A = 15
		border := win.th.Palette.Fg
		border.A = 40

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				if w == 0 {
					w = gtx.Constraints.Max.X
				}
				h := gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(5))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				paint.FillShape(gtx.Ops, border,
					clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Body2(win.th, code)
					lbl.Font.Typeface = "monospace"
					accent := win.th.Palette.ContrastBg
					accent.A = 230
					lbl.Color = accent
					return lbl.Layout(gtx)
				})
			}),
		)
	})
}

// ColoredText draws a label with an explicit color without needing to build
// a full Text() call. Saves a line when you just want one color change.
//
//	proton.ColoredText(win, "Warning: this cannot be undone", proton.RGB(0xfbbf24))
func ColoredText(win *Win, text string, c color.NRGBA) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Body1(win.th, text)
		lbl.Color = c
		return lbl.Layout(gtx)
	})
}

// Muted draws body text in a dimmer color — good for secondary info,
// placeholders, descriptions that shouldn't compete with the main content.
//
//	proton.Label(win, "Alice Johnson")
//	proton.Muted(win, "alice@example.com")
func Muted(win *Win, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Body2(win.th, text)
		c := win.th.Palette.Fg
		c.A = 140
		lbl.Color = c
		return lbl.Layout(gtx)
	})
}

// ErrorText draws text in the theme's error/red color.
// Pass an empty string to draw nothing (useful for conditional errors).
//
//	proton.ErrorText(win, validationErr)
func ErrorText(win *Win, text string) {
	if text == "" {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Body2(win.th, text)
		lbl.Color = color.NRGBA{R: 248, G: 113, B: 113, A: 255}
		return lbl.Layout(gtx)
	})
}

// SuccessText draws text in green. Same empty-string shortcut as ErrorText.
//
//	proton.SuccessText(win, "Saved successfully!")
func SuccessText(win *Win, text string) {
	if text == "" {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Body2(win.th, text)
		lbl.Color = color.NRGBA{R: 74, G: 222, B: 128, A: 255}
		return lbl.Layout(gtx)
	})
}

// WarningText draws text in yellow/amber.
func WarningText(win *Win, text string) {
	if text == "" {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Body2(win.th, text)
		lbl.Color = color.NRGBA{R: 251, G: 191, B: 36, A: 255}
		return lbl.Layout(gtx)
	})
}
