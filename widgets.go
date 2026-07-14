package proton

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// clickResults, boolResults, enumResults store the click/change state from
// the previous frame so Button(), Checkbox() etc. can return correct values
// even though their layout closures run after the draw function returns.
var clickResults = map[*Clickable]bool{}
var boolResults  = map[*Bool]bool{}
var enumResults  = map[*Enum]bool{}

func Label(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return material.Body1(win.theme(), text).Layout(gtx)
	})
}

func H1(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H1(win.theme(), text).Layout(gtx) })
}
func H2(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H2(win.theme(), text).Layout(gtx) })
}
func H3(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H3(win.theme(), text).Layout(gtx) })
}
func H4(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H4(win.theme(), text).Layout(gtx) })
}
func H5(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H5(win.theme(), text).Layout(gtx) })
}
func H6(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.H6(win.theme(), text).Layout(gtx) })
}
func Body2(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Body2(win.theme(), text).Layout(gtx) })
}
func Caption(win Context, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions { return material.Caption(win.theme(), text).Layout(gtx) })
}

func Text(win Context, s string, size float32, c color.NRGBA, bold bool) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(win.theme(), unit.Sp(size), s)
		if c != (color.NRGBA{}) {
			lbl.Color = c
		}
		if bold {
			lbl.Font.Weight = font.Bold
		}
		return lbl.Layout(gtx)
	})
}

// Button draws a filled button with hover and press states.
func Button(win Context, state *Clickable, label string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return drawButton(gtx, win, state, label, false)
	})
	return result
}

// OutlineButton draws a ghost/outline button with hover and press states.
func OutlineButton(win Context, state *Clickable, label string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return drawButton(gtx, win, state, label, true)
	})
	return result
}

func drawButton(gtx layout.Context, win Context, state *Clickable, label string, outline bool) layout.Dimensions {
	th := win.theme()
	primary := th.Palette.ContrastBg
	primaryFg := th.Palette.ContrastFg
	r := gtx.Dp(unit.Dp(6))

	// compute colors based on interaction state
	bg := primary
	fg := primaryFg
	borderC := primary

	if outline {
		bg = color.NRGBA{}
		fg = primary
		borderC = primary
		borderC.A = 180
	}

	if state.Hovered() {
		if outline {
			bg = primary
			bg.A = 25
		} else {
			bg = lightenNRGBA(primary, 20)
		}
	}
	if state.Pressed() {
		if outline {
			bg = primary
			bg.A = 45
		} else {
			bg = darkenNRGBA(primary, 20)
		}
	}

	return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				h := gtx.Constraints.Min.Y
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				if bg.A > 0 {
					paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				}
				if outline {
					paint.FillShape(gtx.Ops, borderC,
						clip.Stroke{Path: rrect.Path(gtx.Ops), Width: float32(gtx.Dp(unit.Dp(1)))}.Op())
				}
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Body2(th, label)
					lbl.Color = fg
					lbl.Font.Weight = font.Medium
					return lbl.Layout(gtx)
				})
			}),
		)
	})
}

func IconButton(win Context, state *Clickable, icon *Icon, desc string) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return material.IconButton(win.theme(), state, icon, desc).Layout(gtx)
	})
	return result
}

// Tappable makes any content clickable. Returns true if clicked.
func Tappable(win Context, state *Clickable, content func(Context)) bool {
	result := clickResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		clickResults[state] = state.Clicked(gtx)
		return state.Layout(gtx, child(win, content))
	})
	return result
}

// Input draws a styled single-line text field with focus ring.
func Input(win Context, state *Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.SingleLine = true
		return drawInput(gtx, win, state, hint, false)
	})
}

// TextArea draws a styled multi-line text field with focus ring.
func TextArea(win Context, state *Editor, hint string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return drawInput(gtx, win, state, hint, true)
	})
}

func drawInput(gtx layout.Context, win Context, state *widget.Editor, hint string, multiLine bool) layout.Dimensions {
	th := win.theme()
	r := gtx.Dp(unit.Dp(6))

	bg := th.Palette.Fg
	bg.A = 14
	border := th.Palette.Fg
	border.A = 45

	focused := state.Len() > 0 // approximation — use focused ring when has content
	_ = focused

	// focus ring: use primary color border when active
	// (Gio doesn't expose focus state directly in v0.8 without event.Op — we
	// approximate with a slightly brighter border, always present)
	focusBorder := th.Palette.ContrastBg
	focusBorder.A = 0 // hidden by default

	h := gtx.Dp(unit.Dp(38))
	if multiLine {
		h = gtx.Dp(unit.Dp(100))
	}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			w := gtx.Constraints.Max.X
			if gtx.Constraints.Min.X > 0 {
				w = gtx.Constraints.Min.X
			}
			rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
			paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
			paint.FillShape(gtx.Ops, border,
				clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1}.Op())
			return layout.Dimensions{Size: image.Pt(w, h)}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min = image.Pt(gtx.Constraints.Max.X, h)
			gtx.Constraints.Max.Y = h
			return layout.UniformInset(unit.Dp(9)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				ed := material.Editor(th, state, hint)
				ed.HintColor = th.Palette.Fg
				ed.HintColor.A = 90
				return ed.Layout(gtx)
			})
		}),
	)
}

// Toggle draws a styled on/off switch.
func Toggle(win Context, state *Bool, label string) bool {
	result := boolResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		boolResults[state] = state.Update(gtx)
		return drawToggle(gtx, win, state, label)
	})
	return result
}

func drawToggle(gtx layout.Context, win Context, state *widget.Bool, label string) layout.Dimensions {
	th := win.theme()
	trackW := gtx.Dp(unit.Dp(40))
	trackH := gtx.Dp(unit.Dp(22))
	knobSz := gtx.Dp(unit.Dp(16))
	pad := (trackH - knobSz) / 2

	trackOn := th.Palette.ContrastBg
	trackOff := th.Palette.Fg
	trackOff.A = 50
	if state.Hovered() {
		trackOff.A = 80
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				r := trackH / 2
				trackC := trackOff
				if state.Value {
					trackC = trackOn
				}
				rrect := clip.RRect{
					Rect: image.Rect(0, 0, trackW, trackH),
					NW: r, NE: r, SE: r, SW: r,
				}
				paint.FillShape(gtx.Ops, trackC, rrect.Op(gtx.Ops))

				// knob
				knobX := pad
				if state.Value {
					knobX = trackW - knobSz - pad
				}
				knobR := knobSz / 2
				knobC := color.NRGBA{R: 255, G: 255, B: 255, A: 230}
				paint.FillShape(gtx.Ops, knobC,
					clip.Ellipse{
						Min: image.Pt(knobX, pad),
						Max: image.Pt(knobX+knobSz, pad+knobSz),
					}.Op(gtx.Ops))
				_ = knobR
				return layout.Dimensions{Size: image.Pt(trackW, trackH)}
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if label == "" {
				return layout.Dimensions{}
			}
			return layout.Spacer{Width: unit.Dp(10)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if label == "" {
				return layout.Dimensions{}
			}
			return material.Body1(th, label).Layout(gtx)
		}),
	)
}

// Checkbox draws a styled checkbox.
func Checkbox(win Context, state *Bool, label string) bool {
	result := boolResults[state]
	win.add(func(gtx layout.Context) layout.Dimensions {
		boolResults[state] = state.Update(gtx)
		return drawCheckbox(gtx, win, state, label)
	})
	return result
}

func drawCheckbox(gtx layout.Context, win Context, state *widget.Bool, label string) layout.Dimensions {
	th := win.theme()
	sz := gtx.Dp(unit.Dp(18))
	r := gtx.Dp(unit.Dp(4))
	primary := th.Palette.ContrastBg

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return state.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				rrect := clip.RRect{Rect: image.Rect(0, 0, sz, sz), NW: r, NE: r, SE: r, SW: r}
				if state.Value {
					paint.FillShape(gtx.Ops, primary, rrect.Op(gtx.Ops))
					// checkmark: two lines
					checkC := th.Palette.ContrastFg
					lx, ly := sz/5, sz/2
					mx, my := sz*2/5, sz*3/4
					rx, ry := sz*4/5, sz/4
					drawLine(gtx, lx, ly, mx, my, checkC, 2)
					drawLine(gtx, mx, my, rx, ry, checkC, 2)
				} else {
					border := th.Palette.Fg
					border.A = 120
					if state.Hovered() {
						border.A = 200
					}
					bg := th.Palette.Fg
					bg.A = 10
					paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
					paint.FillShape(gtx.Ops, border,
						clip.Stroke{Path: rrect.Path(gtx.Ops), Width: 1.5}.Op())
				}
				return layout.Dimensions{Size: image.Pt(sz, sz)}
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(8)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(th, label).Layout(gtx)
		}),
	)
}

// drawLine draws a 1-2px line between two points using tiny rects (no path API needed).
func drawLine(gtx layout.Context, x0, y0, x1, y1 int, c color.NRGBA, thickness int) {
	steps := 12
	for i := 0; i <= steps; i++ {
		t := float32(i) / float32(steps)
		x := int(float32(x0) + t*float32(x1-x0))
		y := int(float32(y0) + t*float32(y1-y0))
		paint.FillShape(gtx.Ops, c,
			clip.Rect{Min: image.Pt(x, y), Max: image.Pt(x+thickness, y+thickness)}.Op())
	}
}

// RadioButton draws a styled radio button.
func RadioButton(win Context, group *Enum, key, label string) bool {
	result := enumResults[group]
	win.add(func(gtx layout.Context) layout.Dimensions {
		enumResults[group] = group.Update(gtx)
		return drawRadio(gtx, win, group, key, label)
	})
	return result
}

func drawRadio(gtx layout.Context, win Context, group *widget.Enum, key, label string) layout.Dimensions {
	th := win.theme()
	sz := gtx.Dp(unit.Dp(18))
	primary := th.Palette.ContrastBg
	selected := group.Value == key

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return group.Layout(gtx, key, func(gtx layout.Context) layout.Dimensions {
				border := th.Palette.Fg
				border.A = 120

				bg := th.Palette.Fg
				bg.A = 10
				paint.FillShape(gtx.Ops, bg,
					clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Op(gtx.Ops))

				if selected {
					paint.FillShape(gtx.Ops, primary,
						clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Op(gtx.Ops))
					inner := sz / 3
					off := (sz - inner) / 2
					white := th.Palette.ContrastFg
					paint.FillShape(gtx.Ops, white,
						clip.Ellipse{Min: image.Pt(off, off), Max: image.Pt(off+inner, off+inner)}.Op(gtx.Ops))
				} else {
					paint.FillShape(gtx.Ops, border,
						clip.Stroke{
							Path: clip.Ellipse{Min: image.Pt(0, 0), Max: image.Pt(sz, sz)}.Path(gtx.Ops),
							Width: 1.5,
						}.Op())
				}
				return layout.Dimensions{Size: image.Pt(sz, sz)}
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(8)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Body1(th, label)
			if selected {
				lbl.Color = primary
			}
			return lbl.Layout(gtx)
		}),
	)
}

func Slider(win Context, state *Float) float32 {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return drawSlider(gtx, win, state)
	})
	return state.Value
}

func drawSlider(gtx layout.Context, win Context, state *widget.Float) layout.Dimensions {
	th := win.theme()
	h := gtx.Dp(unit.Dp(4))
	trackH := gtx.Dp(unit.Dp(4))
	knobSz := gtx.Dp(unit.Dp(16))
	w := gtx.Constraints.Max.X
	totalH := knobSz + 4
	yCenter := totalH / 2

	// position knob
	if state.Value < 0 {
		state.Value = 0
	}
	if state.Value > 1 {
		state.Value = 1
	}
	knobX := int(state.Value*float32(w-knobSz))

	// track background
	trackY := yCenter - trackH/2
	trackBg := th.Palette.Fg
	trackBg.A = 40
	paint.FillShape(gtx.Ops, trackBg,
		clip.RRect{
			Rect: image.Rect(0, trackY, w, trackY+trackH),
			NW: trackH / 2, NE: trackH / 2, SE: trackH / 2, SW: trackH / 2,
		}.Op(gtx.Ops))

	// filled portion
	fillW := knobX + knobSz/2
	if fillW > 0 {
		primary := th.Palette.ContrastBg
		paint.FillShape(gtx.Ops, primary,
			clip.RRect{
				Rect: image.Rect(0, trackY, fillW, trackY+trackH),
				NW: trackH / 2, NE: trackH / 2, SE: trackH / 2, SW: trackH / 2,
			}.Op(gtx.Ops))
	}

	// knob
	knobC := th.Palette.Bg
	knobR := knobSz / 2
	shadow := th.Palette.Fg
	shadow.A = 60
	// shadow dot (slightly offset)
	paint.FillShape(gtx.Ops, shadow,
		clip.Ellipse{Min: image.Pt(knobX+1, yCenter-knobR+1), Max: image.Pt(knobX+knobSz+1, yCenter+knobR+1)}.Op(gtx.Ops))
	paint.FillShape(gtx.Ops, th.Palette.ContrastBg,
		clip.Ellipse{Min: image.Pt(knobX, yCenter-knobR), Max: image.Pt(knobX+knobSz, yCenter+knobR)}.Op(gtx.Ops))
	paint.FillShape(gtx.Ops, knobC,
		clip.Ellipse{Min: image.Pt(knobX+2, yCenter-knobR+2), Max: image.Pt(knobX+knobSz-2, yCenter+knobR-2)}.Op(gtx.Ops))

	// interaction overlay (invisible but handles events via material.Slider)
	_ = h
	macro := material.Slider(th, state)
	macro.Layout(gtx)

	return layout.Dimensions{Size: image.Pt(w, totalH)}
}

// ProgressBar draws a styled filled progress bar.
func ProgressBar(win Context, progress float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return drawProgressBar(gtx, win, progress)
	})
}

func drawProgressBar(gtx layout.Context, win Context, progress float32) layout.Dimensions {
	th := win.theme()
	w := gtx.Constraints.Max.X
	h := gtx.Dp(unit.Dp(6))
	r := h / 2

	// track
	trackC := th.Palette.Fg
	trackC.A = 30
	paint.FillShape(gtx.Ops, trackC,
		clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}.Op(gtx.Ops))

	// fill
	if progress > 0 {
		fillW := int(float32(w) * progress)
		if fillW > w {
			fillW = w
		}
		paint.FillShape(gtx.Ops, th.Palette.ContrastBg,
			clip.RRect{Rect: image.Rect(0, 0, fillW, h), NW: r, NE: r, SE: r, SW: r}.Op(gtx.Ops))
	}
	return layout.Dimensions{Size: image.Pt(w, h)}
}

func List(win Context, state *Scrollable, length int, draw func(Context, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.theme(), state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w Context) { draw(w, i) })(gtx)
		})
	})
}

func HList(win Context, state *Scrollable, length int, draw func(Context, int)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Horizontal
		return material.List(win.theme(), state).Layout(gtx, length, func(gtx layout.Context, i int) layout.Dimensions {
			return child(win, func(w Context) { draw(w, i) })(gtx)
		})
	})
}

func Scroll(win Context, state *Scrollable, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		state.Axis = layout.Vertical
		return material.List(win.theme(), state).Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
			return child(win, content)(gtx)
		})
	})
}

func Divider(win Context) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		h := gtx.Dp(unit.Dp(1))
		w := gtx.Constraints.Max.X
		c := win.theme().Palette.Fg
		c.A = 35
		paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func Rect(win Context, c color.NRGBA, widthDp, heightDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, h)}.Op())
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func RoundRect(win Context, c color.NRGBA, widthDp, heightDp, radiusDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w := gtx.Constraints.Max.X
		h := gtx.Constraints.Max.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		r := gtx.Dp(unit.Dp(radiusDp))
		rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
		paint.FillShape(gtx.Ops, c, rrect.Op(gtx.Ops))
		return layout.Dimensions{Size: image.Pt(w, h)}
	})
}

func Card(win Context, bg color.NRGBA, cornerDp, padDp float32, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		r := gtx.Dp(unit.Dp(cornerDp))
		// subtle drop shadow
		shadow := win.theme().Palette.Fg
		shadow.A = 18
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w := gtx.Constraints.Min.X
				h := gtx.Constraints.Min.Y
				// shadow (offset by 2px)
				shadowRect := clip.RRect{Rect: image.Rect(2, 3, w+2, h+3), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, shadow, shadowRect.Op(gtx.Ops))
				// card background
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(padDp)).Layout(gtx, child(win, content))
			}),
		)
	})
}

func Badge(win Context, bg, fg color.NRGBA, text string) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w, h := gtx.Constraints.Min.X, gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(12))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, bg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Caption(win.theme(), text)
				lbl.Color = fg
				return layout.Inset{Left: unit.Dp(8), Right: unit.Dp(8), Top: unit.Dp(3), Bottom: unit.Dp(3)}.
					Layout(gtx, lbl.Layout)
			}),
		)
	})
}

func MinSize(win Context, widthDp, heightDp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if widthDp > 0 {
			if m := gtx.Dp(unit.Dp(widthDp)); gtx.Constraints.Min.X < m {
				gtx.Constraints.Min.X = m
			}
		}
		if heightDp > 0 {
			if m := gtx.Dp(unit.Dp(heightDp)); gtx.Constraints.Min.Y < m {
				gtx.Constraints.Min.Y = m
			}
		}
		return child(win, fn)(gtx)
	})
}

func MaxWidth(win Context, widthDp float32, fn func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		if m := gtx.Dp(unit.Dp(widthDp)); gtx.Constraints.Max.X > m {
			gtx.Constraints.Max.X = m
		}
		return child(win, fn)(gtx)
	})
}

type ImageOp struct {
	op paint.ImageOp
	sz image.Point
}

func LoadImage(path string) (ImageOp, error) {
	f, err := os.Open(path)
	if err != nil {
		return ImageOp{}, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return ImageOp{}, err
	}
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		bounds := img.Bounds()
		nrgba = image.NewNRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				nrgba.Set(x, y, img.At(x, y))
			}
		}
	}
	return ImageOp{op: paint.NewImageOp(nrgba), sz: nrgba.Bounds().Size()}, nil
}

func Image(win Context, img ImageOp, widthDp, heightDp float32) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		w, h := img.sz.X, img.sz.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		sz := image.Pt(w, h)
		stack := clip.Rect{Max: sz}.Push(gtx.Ops)
		img.op.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		stack.Pop()
		return layout.Dimensions{Size: sz}
	})
}

func lightenNRGBA(c color.NRGBA, amt uint8) color.NRGBA {
	add := func(v uint8) uint8 {
		if int(v)+int(amt) > 255 {
			return 255
		}
		return v + amt
	}
	return color.NRGBA{R: add(c.R), G: add(c.G), B: add(c.B), A: c.A}
}

func darkenNRGBA(c color.NRGBA, amt uint8) color.NRGBA {
	sub := func(v uint8) uint8 {
		if int(v)-int(amt) < 0 {
			return 0
		}
		return v - amt
	}
	return color.NRGBA{R: sub(c.R), G: sub(c.G), B: sub(c.B), A: c.A}
}
