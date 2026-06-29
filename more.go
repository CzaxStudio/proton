package proton

// more.go — additional widgets: Tag, Avatar, ProgressRing, Stepper, Table, SearchInput.

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ----- Tag -----

// TagState holds state for a dismissable tag chip.
//
//	type UI struct {
//	    tags    []string
//	    tagBtns []proton.Clickable
//	}
type TagState struct {
	btn Clickable
}

// Tag draws a small colored chip with optional dismiss button.
// If onClose is not nil, an × button appears and onClose is called when clicked.
//
//	proton.Tag(win, &u.tagState, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "golang", func() {
//	    removTag("golang")
//	})
func Tag(win Context, state *TagState, bg, fg color.NRGBA, label string, onClose func()) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if onClose != nil && clickResults[&state.btn] {
			onClose()
		}

		r := gtx.Dp(unit.Dp(12))
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				h := gtx.Constraints.Min.Y
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: unit.Dp(8), Right: unit.Dp(4), Top: unit.Dp(3), Bottom: unit.Dp(3)}.
					Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl := material.Caption(win.theme(), label)
								lbl.Color = fg
								return lbl.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if onClose == nil {
									return layout.Spacer{Width: unit.Dp(4)}.Layout(gtx)
								}
								return layout.Spacer{Width: unit.Dp(4)}.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if onClose == nil {
									return layout.Dimensions{}
								}
								clickResults[&state.btn] = state.btn.Clicked(gtx)
								return state.btn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									lbl := material.Caption(win.theme(), "×")
									c := fg
									c.A = 180
									lbl.Color = c
									return layout.Inset{Right: unit.Dp(4)}.Layout(gtx, lbl.Layout)
								})
							}),
						)
					})
			}),
		)
	})
}

// ----- Avatar -----

// Avatar draws a circular badge showing initials or a short label.
// Good for user profile pictures when no image is available.
//
//	proton.Avatar(win, "AJ", proton.RGB(0x5e81ac), proton.RGB(0xeceff4), 40)
func Avatar(win Context, initials string, bg, fg color.NRGBA, sizeDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		sz := gtx.Dp(unit.Dp(sizeDp))
		half := sz / 2
		paint.FillShape(gtx.Ops, bg,
			clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Op(gtx.Ops))
		// center the text
		lbl := material.Body2(win.theme(), initials)
		lbl.Color = fg
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				paint.FillShape(gtx.Ops, bg,
					clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(sz, sz)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				_ = half
				return layout.Center.Layout(gtx, lbl.Layout)
			}),
		)
	})
}

// ----- ProgressRing -----

// ProgressRing draws a circular progress indicator.
// progress is 0.0–1.0. sizeDp is the diameter. strokeDp is the ring width.
//
//	proton.ProgressRing(win, 0.72, 48, 5, proton.RGB(0x88c0d0))
func ProgressRing(win Context, progress, sizeDp, strokeDp float32, c color.NRGBA) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		sz := gtx.Dp(unit.Dp(sizeDp))
		stroke := gtx.Dp(unit.Dp(strokeDp))

		cx := float32(sz) / 2
		cy := float32(sz) / 2
		r := cx - float32(stroke)/2

		// background ring
		bg := c
		bg.A = 30
		drawArc(gtx, cx, cy, r, float32(stroke), 0, 1, bg)

		// progress arc
		if progress > 0 {
			drawArc(gtx, cx, cy, r, float32(stroke), 0, progress, c)
		}

		return layout.Dimensions{Size: image.Pt(sz, sz)}
	})
}

func drawArc(gtx layout.Context, cx, cy, r, strokeW, from, to float32, c color.NRGBA) {
	steps := int(48 * (to - from))
	if steps < 2 {
		steps = 2
	}
	for i := 0; i < steps; i++ {
		t := from + (to-from)*float32(i)/float32(steps)
		a := float64(t*2*math.Pi) - math.Pi/2
		x := cx + r*float32(math.Cos(a))
		y := cy + r*float32(math.Sin(a))
		bx := int(x - strokeW/2)
		by := int(y - strokeW/2)
		bw := int(strokeW + 0.5)
		if bw < 1 {
			bw = 1
		}
		paint.FillShape(gtx.Ops, c,
			clip.Rect{Min: image.Pt(bx, by), Max: image.Pt(bx+bw, by+bw)}.Op())
	}
}

// ----- Stepper -----

// Stepper draws a horizontal step-progress indicator.
// current is the active step index (0-based). steps is the list of step names.
//
//	proton.Stepper(win, 1, []string{"Account", "Profile", "Payment", "Done"})
func Stepper(win Context, current int, steps []string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		accent := win.theme().Palette.ContrastBg
		muted := win.theme().Palette.Fg
		muted.A = 60
		textMuted := win.theme().Palette.Fg
		textMuted.A = 120

		stepW := gtx.Constraints.Max.X / len(steps)
		lineH := gtx.Dp(unit.Dp(2))
		dotSz := gtx.Dp(unit.Dp(24))
		totalH := dotSz + gtx.Dp(unit.Dp(20))

		children := make([]layout.FlexChild, len(steps))
		for i, name := range steps {
			i, name := i, name
			children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = stepW
				gtx.Constraints.Max.X = stepW

				return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								if i == 0 {
									return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X/2, lineH)}
								}
								c := muted
								if i <= current {
									c = accent
								}
								paint.FillShape(gtx.Ops, c,
									clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, lineH)}.Op())
								return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, lineH)}
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								dotC := muted
								if i <= current {
									dotC = accent
								}
								paint.FillShape(gtx.Ops, dotC,
									clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(dotSz, dotSz)}.Op(gtx.Ops))
								// number inside dot
								lbl := material.Caption(win.theme(), fmt.Sprintf("%d", i+1))
								if i <= current {
									lbl.Color = win.theme().Palette.ContrastFg
								} else {
									lbl.Color = win.theme().Palette.Fg
									lbl.Color.A = 180
								}
								layout.Stack{Alignment: layout.Center}.Layout(gtx,
									layout.Expanded(func(gtx layout.Context) layout.Dimensions {
										return layout.Dimensions{Size: image.Pt(dotSz, dotSz)}
									}),
									layout.Stacked(lbl.Layout),
								)
								return layout.Dimensions{Size: image.Pt(dotSz, dotSz)}
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								if i == len(steps)-1 {
									return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X/2, lineH)}
								}
								c := muted
								if i < current {
									c = accent
								}
								paint.FillShape(gtx.Ops, c,
									clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, lineH)}.Op())
								return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, lineH)}
							}),
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{Height: unit.Dp(6)}.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Caption(win.theme(), name)
						if i == current {
							lbl.Color = accent
							lbl.Font.Weight = 500
						} else {
							lbl.Color = textMuted
						}
						return layout.Center.Layout(gtx, lbl.Layout)
					}),
				)
			})
		}

		_ = totalH
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
}

// ----- Table -----

// TableRow is one row in a Table. It is a slice of cell strings.
type TableRow = []string

// Table draws a simple data table with a header row and data rows.
// columns is the list of column headers. rows is the data.
//
//	proton.Table(win,
//	    []string{"Name", "Status", "Score"},
//	    []proton.TableRow{
//	        {"Alice", "Active", "98"},
//	        {"Bob",   "Away",   "74"},
//	    },
//	)
func Table(win Context, columns []string, rows []TableRow) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if len(columns) == 0 {
			return layout.Dimensions{}
		}

		colW := gtx.Constraints.Max.X / len(columns)
		rowH := gtx.Dp(unit.Dp(36))
		totalH := rowH * (1 + len(rows))
		bg := win.theme().Palette.Fg
		bg.A = 8
		border := win.theme().Palette.Fg
		border.A = 25
		headerBg := win.theme().Palette.ContrastBg
		headerBg.A = 30

		// draw background
		paint.FillShape(gtx.Ops, bg,
			clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, totalH)}.Op())

		allRows := make([]layout.FlexChild, 0, 1+len(rows))

		// header
		allRows = append(allRows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, headerBg,
				clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, rowH)}.Op())
			cells := make([]layout.FlexChild, len(columns))
			for i, col := range columns {
				i, col := i, col
				cells[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints = layout.Exact(image.Pt(colW, rowH))
					lbl := material.Body2(win.theme(), col)
					lbl.Font.Weight = 600
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, lbl.Layout)
				})
			}
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, cells...)
		}))

		// divider
		allRows = append(allRows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			h := gtx.Dp(unit.Dp(1))
			paint.FillShape(gtx.Ops, border,
				clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, h)}.Op())
			return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, h)}
		}))

		for ri, row := range rows {
			ri, row := ri, row
			allRows = append(allRows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if ri%2 == 1 {
					stripe := win.theme().Palette.Fg
					stripe.A = 6
					paint.FillShape(gtx.Ops, stripe,
						clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, rowH)}.Op())
				}
				cells := make([]layout.FlexChild, len(columns))
				for ci := range columns {
					ci := ci
					text := ""
					if ci < len(row) {
						text = row[ci]
					}
					cells[ci] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints = layout.Exact(image.Pt(colW, rowH))
						lbl := material.Body2(win.theme(), text)
						return layout.UniformInset(unit.Dp(8)).Layout(gtx, lbl.Layout)
					})
				}
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, cells...)
			}))

			if ri < len(rows)-1 {
				allRows = append(allRows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					h := gtx.Dp(unit.Dp(1))
					paint.FillShape(gtx.Ops, border,
						clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, h)}.Op())
					return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, h)}
				}))
			}
		}

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, allRows...)
	})
}

// ----- SearchInput -----

// SearchState holds state for a SearchInput.
//
//	type UI struct {
//	    search proton.SearchState
//	}
type SearchState struct {
	Editor   widget.Editor
	clearBtn widget.Clickable
}

// SearchInput draws a text field with a placeholder and a clear (×) button.
// Returns the current query string.
//
//	q := proton.SearchInput(win, &u.search, "Search notes...")
//	filtered := filter(items, q)
func SearchInput(win Context, state *SearchState, placeholder string) string {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if clickResults[&state.clearBtn] {
			state.Editor.SetText("")
		}
		clickResults[&state.clearBtn] = state.clearBtn.Clicked(gtx)

		border := win.theme().Palette.Fg
		border.A = 50
		bg := win.theme().Palette.Fg
		bg.A = 12

		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				if w == 0 {
					w = gtx.Constraints.Max.X
				}
				h := gtx.Constraints.Min.Y
				if h == 0 {
					h = gtx.Dp(unit.Dp(38))
				}
				r := gtx.Dp(unit.Dp(6))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				paint.FillShape(gtx.Ops, border,
					clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1}.Op())
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Caption(win.theme(), "⌕")
							c := win.theme().Palette.Fg
							c.A = 120
							lbl.Color = c
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: unit.Dp(6)}.Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							state.Editor.SingleLine = true
							return material.Editor(win.theme(), &state.Editor, placeholder).Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if state.Editor.Text() == "" {
								return layout.Dimensions{}
							}
							return state.clearBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Caption(win.theme(), "×")
								c := win.theme().Palette.Fg
								c.A = 150
								lbl.Color = c
								return lbl.Layout(gtx)
							})
						}),
					)
				})
			}),
		)
	})
	return state.Editor.Text()
}
