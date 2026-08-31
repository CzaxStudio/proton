package proton

// extra.go — additional widgets that build on the core.

import (
	"image"
	"image/color"
	"math"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ----- Tabs -----

// TabState tracks which tab is selected.
//
//	type UI struct {
//	    tabs    proton.TabState
//	    tabBtns [3]proton.Clickable
//	}
type TabState struct {
	Selected int
}

// Tabs draws a horizontal tab bar. labels are the tab names, btns is one
// Clickable per tab. content is called with the selected tab index.
//
//	proton.Tabs(win, []string{"Files", "Settings"}, u.tabBtns[:], &u.tabs,
//	    func(win proton.Context, i int) {
//	        switch i {
//	        case 0: drawFiles(win)
//	        case 1: drawSettings(win)
//	        }
//	    },
//	)
func Tabs(win Context, labels []string, btns []Clickable, state *TabState, content func(Context, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				tabChildren := make([]layout.FlexChild, len(labels))
				for i, label := range labels {
					i, label := i, label
					tabChildren[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if btns[i].Clicked(gtx) {
							state.Selected = i
						}
						active := state.Selected == i
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return btns[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										lbl := material.Body2(win.theme(), label)
										if active {
											lbl.Color = win.theme().Palette.ContrastBg
											lbl.Font.Weight = 500
										} else {
											c := win.theme().Palette.Fg
											c.A = 150
											lbl.Color = c
										}
										return lbl.Layout(gtx)
									})
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if !active {
									return layout.Dimensions{}
								}
								h := gtx.Dp(unit.Dp(2))
								w := gtx.Constraints.Max.X
								paint.FillShape(gtx.Ops, win.theme().Palette.ContrastBg,
									clip.Rect{Max: image.Pt(w, h)}.Op())
								return layout.Dimensions{Size: image.Pt(w, h)}
							}),
						)
					})
				}
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, tabChildren...)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				h := gtx.Dp(unit.Dp(1))
				w := gtx.Constraints.Max.X
				c := win.theme().Palette.Fg
				c.A = 35
				paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(12)).Layout(gtx,
					child(win, func(w Context) { content(w, state.Selected) }))
			}),
		)
	})
}

// ----- Accordion -----

// AccordionState tracks whether a collapsible section is open or closed.
//
//	type UI struct {
//	    sec1    proton.AccordionState
//	    sec1btn proton.Clickable
//	}
type AccordionState struct {
	Open bool
}

// Accordion draws a collapsible section with a clickable header.
//
//	proton.Accordion(win, &u.sec1, &u.sec1btn, "Advanced", func(win proton.Context) {
//	    proton.Label(win, "Hidden until expanded.")
//	})
func Accordion(win Context, state *AccordionState, btn *Clickable, title string, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if btn.Clicked(gtx) {
			state.Open = !state.Open
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return btn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return drawChevron(gtx, win, state.Open)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Width: unit.Dp(8)}.Layout(gtx)
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Body1(win.theme(), title)
								lbl.Font.Weight = 500
								return lbl.Layout(gtx)
							}),
						)
					})
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !state.Open {
					return layout.Dimensions{}
				}
				return layout.Inset{Left: unit.Dp(24), Top: unit.Dp(4), Bottom: unit.Dp(8), Right: unit.Dp(0)}.
					Layout(gtx, child(win, content))
			}),
		)
	})
}

// drawChevron draws a small triangle pointing right (collapsed) or down (open).
func drawChevron(gtx layout.Context, win Context, open bool) layout.Dimensions {
	sz := gtx.Dp(unit.Dp(8))
	c := win.theme().Palette.Fg
	c.A = 180
	var path clip.Path
	path.Begin(gtx.Ops)
	if open {
		path.MoveTo(f32.Pt(0, 0))
		path.LineTo(f32.Pt(float32(sz), 0))
		path.LineTo(f32.Pt(float32(sz)/2, float32(sz)))
	} else {
		path.MoveTo(f32.Pt(0, 0))
		path.LineTo(f32.Pt(float32(sz), float32(sz)/2))
		path.LineTo(f32.Pt(0, float32(sz)))
	}
	path.Close()
	paint.FillShape(gtx.Ops, c, clip.Outline{Path: path.End()}.Op())
	return layout.Dimensions{Size: image.Pt(sz+4, sz)}
}

// ----- Spinner -----

// SpinnerState tracks the animation start time.
//
//	type UI struct {
//	    spin proton.SpinnerState
//	}
//
//	proton.Spinner(win, &u.spin, 32)
type SpinnerState struct {
	start time.Time
}

// Spinner draws an animated circular loading indicator. sizeDp is the diameter.
func Spinner(win Context, state *SpinnerState, sizeDp float32) {
	if state.start.IsZero() {
		state.start = time.Now()
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		sz := gtx.Dp(unit.Dp(sizeDp))
		elapsed := time.Since(state.start).Seconds()
		angle := float32(elapsed * 2 * math.Pi * 1.1)

		cx := float32(sz) / 2
		cy := float32(sz) / 2
		r := cx * 0.7
		strokeW := cx * 0.13

		// background ring
		bgC := win.theme().Palette.Fg
		bgC.A = 30
		steps := 48
		for i := 0; i < steps; i++ {
			a := float32(i) / float32(steps) * 2 * math.Pi
			x := cx + r*float32(math.Cos(float64(a)))
			y := cy + r*float32(math.Sin(float64(a)))
			bx := int(x - strokeW/2)
			by := int(y - strokeW/2)
			bw := int(strokeW + 0.5)
			if bw < 1 {
				bw = 1
			}
			paint.FillShape(gtx.Ops, bgC,
				clip.Rect{Min: image.Pt(bx, by), Max: image.Pt(bx+bw, by+bw)}.Op())
		}

		// arc (~75% of circle)
		arcC := win.theme().Palette.ContrastBg
		arcSteps := int(float32(steps) * 0.75)
		for i := 0; i < arcSteps; i++ {
			frac := float32(i) / float32(arcSteps)
			a := angle + frac*float32(math.Pi)*1.5
			x := cx + r*float32(math.Cos(float64(a)))
			y := cy + r*float32(math.Sin(float64(a)))
			c := arcC
			c.A = uint8(60 + frac*195)
			bx := int(x - strokeW/2)
			by := int(y - strokeW/2)
			bw := int(strokeW + 0.5)
			if bw < 1 {
				bw = 1
			}
			paint.FillShape(gtx.Ops, c,
				clip.Rect{Min: image.Pt(bx, by), Max: image.Pt(bx+bw, by+bw)}.Op())
		}

		win.Invalidate()
		return layout.Dimensions{Size: image.Pt(sz, sz)}
	})
}

// ----- SelectBox -----

// SelectBoxState holds the open/closed state and per-item clickables.
//
//	type UI struct {
//	    lang proton.SelectBoxState
//	}
//
//	i := proton.SelectBox(win, &u.lang, []string{"Go", "Rust", "Zig"})
type SelectBoxState struct {
	Open     bool
	Selected int
	mainBtn  widget.Clickable
	rowBtns  []widget.Clickable
}

// SelectBox draws a dropdown selector. Returns the index of the selected option.
func SelectBox(win Context, state *SelectBoxState, options []string) int {
	for len(state.rowBtns) < len(options) {
		state.rowBtns = append(state.rowBtns, widget.Clickable{})
	}

	win.add(func(gtx layout.Context) layout.Dimensions {
		if state.mainBtn.Clicked(gtx) {
			state.Open = !state.Open
		}
		for i := range options {
			if state.rowBtns[i].Clicked(gtx) {
				state.Selected = i
				state.Open = false
			}
		}

		label := "Select..."
		if state.Selected >= 0 && state.Selected < len(options) {
			label = options[state.Selected]
		}

		// header button
		headerDims := state.mainBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			bg := win.theme().Palette.Fg
			bg.A = 18
			border := win.theme().Palette.Fg
			border.A = 55
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					w := gtx.Constraints.Max.X
					h := gtx.Dp(unit.Dp(38))
					r := gtx.Dp(unit.Dp(5))
					var rrect clip.RRect
					if state.Open {
						rrect = clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: 0, SW: 0}
					} else {
						rrect = clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
					}
					paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
					paint.FillShape(gtx.Ops, border,
						clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1}.Op())
					return layout.Dimensions{Size: image.Pt(w, h)}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(9)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, material.Body1(win.theme(), label).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return drawChevron(gtx, win, state.Open)
							}),
						)
					})
				}),
			)
		})

		if !state.Open {
			return headerDims
		}

		// draw dropdown below header
		rowH := gtx.Dp(unit.Dp(36))
		totalH := rowH * len(options)
		bg := win.theme().Palette.Fg
		bg.A = 18
		border := win.theme().Palette.Fg
		border.A = 55

		macro := op.Record(gtx.Ops)
		dropGtx := gtx
		dropGtx.Constraints = layout.Exact(image.Pt(headerDims.Size.X, totalH))

		r := gtx.Dp(unit.Dp(5))
		dropRect := clip.RRect{
			Rect: image.Rect(0, 0, headerDims.Size.X, totalH),
			NW:   0, NE: 0, SE: r, SW: r,
		}
		paint.FillShape(gtx.Ops, win.theme().Palette.Bg,
			clip.Rect{Max: image.Pt(headerDims.Size.X, totalH)}.Op())
		paint.FillShape(gtx.Ops, bg, dropRect.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, border,
			clip.Stroke{Path: dropRect.Path(gtx.Ops), Width: 1}.Op())

		rowChildren := make([]layout.FlexChild, len(options))
		for i, opt := range options {
			i, opt := i, opt
			rowChildren[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return state.rowBtns[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					rowBg := color.NRGBA{}
					if i == state.Selected {
						rowBg = win.theme().Palette.ContrastBg
						rowBg.A = 50
					} else if state.rowBtns[i].Hovered() {
						rowBg = win.theme().Palette.Fg
						rowBg.A = 20
					}
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							w := gtx.Constraints.Min.X
							h := rowH
							if rowBg.A > 0 {
								paint.FillShape(gtx.Ops, rowBg,
									clip.Rect{Max: image.Pt(w, h)}.Op())
							}
							return layout.Dimensions{Size: image.Pt(w, h)}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(9)).Layout(gtx,
								material.Body1(win.theme(), opt).Layout)
						}),
					)
				})
			})
		}
		layout.Flex{Axis: layout.Vertical}.Layout(dropGtx, rowChildren...)
		call := macro.Stop()

		stack := op.Offset(image.Pt(0, headerDims.Size.Y-1)).Push(gtx.Ops)
		call.Add(gtx.Ops)
		stack.Pop()

		return headerDims
	})

	if state.Selected < 0 || state.Selected >= len(options) {
		return 0
	}
	return state.Selected
}

// ----- LabeledDivider -----

// LabeledDivider draws a horizontal rule with an optional centered label.
// Pass an empty string for a plain divider.
//
//	proton.LabeledDivider(win, "Advanced Settings")
//	proton.LabeledDivider(win, "")
func LabeledDivider(win Context, label string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		c := win.theme().Palette.Fg
		c.A = 55
		if label == "" {
			h := gtx.Dp(unit.Dp(1))
			w := gtx.Constraints.Max.X
			paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
			return layout.Dimensions{Size: image.Pt(w, h)}
		}
		lbl := material.Caption(win.theme(), label)
		lbl.Color = c
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				h := gtx.Dp(unit.Dp(1))
				w := gtx.Constraints.Max.X
				paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, lbl.Layout)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				h := gtx.Dp(unit.Dp(1))
				w := gtx.Constraints.Max.X
				paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
		)
	})
}

// ----- ZStack (z-axis layering) -----

// ZStack draws multiple widgets layered on top of each other.
// The first is drawn at the bottom, the last is on top.
// All layers share the same available space.
//
//	proton.ZStack(win,
//	    func(win proton.Context) { proton.Rect(win, proton.RGB(0x1e1e2e), 0, 120) },
//	    func(win proton.Context) {
//	        proton.Center(win, func(win proton.Context) { proton.Label(win, "on top") })
//	    },
//	)
func ZStack(win Context, layers ...func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		stacked := make([]layout.StackChild, len(layers))
		for i, fn := range layers {
			fn := fn
			if i == 0 {
				stacked[i] = layout.Expanded(child(win, fn))
			} else {
				stacked[i] = layout.Stacked(child(win, fn))
			}
		}
		return layout.Stack{}.Layout(gtx, stacked...)
	})
}

// ----- Overlay / Modal -----

// OverlayState controls whether the overlay is visible.
//
//	type UI struct {
//	    modal    proton.OverlayState
//	    closeBtn proton.Clickable
//	}
//
//	if proton.Button(win, &u.openBtn, "Open") {
//	    u.modal.Show()
//	}
//	proton.Overlay(win, &u.modal, func(win proton.Context) {
//	    proton.Card(win, proton.RGB(0x1e1e2e), 12, 24, func(win proton.Context) {
//	        proton.H5(win, "Dialog Title")
//	        proton.Gap(win, 8)
//	        proton.Pad(win, 4, func(win proton.Context) {
//	            if proton.Button(win, &u.closeBtn, "Close") { u.modal.Hide() }
//	        })
//	    })
//	})
type OverlayState struct {
	Visible bool
}

// Show makes the overlay visible.
func (o *OverlayState) Show() { o.Visible = true }

// Hide dismisses the overlay.
func (o *OverlayState) Hide() { o.Visible = false }

// Toggle flips the overlay visibility.
func (o *OverlayState) Toggle() { o.Visible = !o.Visible }

// Overlay draws a dimmed backdrop with centered content on top of everything.
// Does nothing when state.Visible is false.
func Overlay(win Context, state *OverlayState, content func(Context)) {
	if !state.Visible {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		paint.FillShape(gtx.Ops, color.NRGBA{A: 165},
			clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Center.Layout(gtx, child(win, content))
	})
}

// ----- Spacer -----

// FlexSpacer returns a flexible empty space that fills remaining room in a
// GrowRow or GrowColumn, pushing siblings to opposite edges.
//
//	proton.GrowRow(win,
//	    proton.FixedItem(win, func(win proton.Context) { proton.Label(win, "left") }),
//	    proton.FlexSpacer(),
//	    proton.FixedItem(win, func(win proton.Context) { proton.Label(win, "right") }),
//	)
func FlexSpacer() layout.FlexChild {
	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: gtx.Constraints.Min}
	})
}

// ----- Conditional rendering helper -----

// If only draws content when cond is true.
// Saves you from wrapping everything in a Go if block when you just
// want to show or hide a single widget.
//
//	proton.If(win, user.IsAdmin, func(win proton.Context) {
//	    proton.Button(win, &u.deleteBtn, "Delete All Users")
//	})
func If(win Context, cond bool, content func(Context)) {
	if !cond {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		return child(win, content)(gtx)
	})
}

// ----- Clickable area with hover background -----

// HoverCard wraps content in an area that highlights on hover.
// bg is the normal background, hover is the color when the mouse is over it.
// Returns true if clicked.
//
//	var row proton.Clickable
//	if proton.HoverCard(win, &row, proton.RGB(0x1e1e2e), proton.RGB(0x2a2a3e), 8, func(win proton.Context) {
//	    proton.Label(win, "hover me")
//	}) {
//	    println("clicked")
//	}
func HoverCard(win Context, state *Clickable, bg, hover color.NRGBA, cornerDp float32, content func(Context)) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			currentBg := bg
			if state.Hovered() {
				currentBg = hover
			}
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					w := gtx.Constraints.Min.X
					h := gtx.Constraints.Min.Y
					r := gtx.Dp(unit.Dp(cornerDp))
					rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
					paint.FillShape(gtx.Ops, currentBg, rrect.Op(gtx.Ops))
					return layout.Dimensions{Size: image.Pt(w, h)}
				}),
				layout.Stacked(child(win, content)),
			)
		})
	})
	return result
}
