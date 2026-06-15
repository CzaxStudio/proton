package proton

import (
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// TextView draws a read-only scrollable text area — useful for log output,
// file previews, help text, or any large block of text.
//
//	var tv proton.Scrollable
//	proton.TextView(win, &tv, longText)
func TextView(win *Win, state *widget.List, text string) {
	lines := strings.Split(text, "\n")
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.th, state).Layout(gtx, len(lines), func(gtx layout.Context, i int) layout.Dimensions {
			return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Body2(win.th, lines[i])
				lbl.Font.Typeface = "monospace"
				return lbl.Layout(gtx)
			})
		})
	})
}

// LogView is like TextView but new lines get appended at the bottom and it
// auto-scrolls to keep the latest output visible.
// Pass the full log string each frame.
//
//	type UI struct {
//	    log    string
//	    logScroll proton.Scrollable
//	}
//
//	u.log += "Build started...\n"
//	proton.LogView(win, &u.logScroll, u.log)
func LogView(win *Win, state *widget.List, text string) {
	lines := strings.Split(text, "\n")
	// remove trailing empty line from the split
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		// scroll to bottom
		if len(lines) > 0 {
			state.ScrollTo(len(lines) - 1)
		}
		return material.List(win.th, state).Layout(gtx, len(lines), func(gtx layout.Context, i int) layout.Dimensions {
			return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Body2(win.th, lines[i])
				lbl.Font.Typeface = "monospace"
				// color-code common log prefixes
				line := lines[i]
				switch {
				case strings.HasPrefix(line, "ERROR") || strings.HasPrefix(line, "[ERROR]"):
					lbl.Color = RGB(0xf87171)
				case strings.HasPrefix(line, "WARN") || strings.HasPrefix(line, "[WARN]"):
					lbl.Color = RGB(0xfbbf24)
				case strings.HasPrefix(line, "OK") || strings.HasPrefix(line, "[OK]") ||
					strings.HasPrefix(line, "DONE") || strings.HasPrefix(line, "SUCCESS"):
					lbl.Color = RGB(0x4ade80)
				default:
					c := win.th.Palette.Fg
					c.A = 190
					lbl.Color = c
				}
				return lbl.Layout(gtx)
			})
		})
	})
}
