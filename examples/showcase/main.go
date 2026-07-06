
package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/CzaxStudio/proton"
)

type ui struct {
	// nav
	tabs    proton.TabState
	tabBtns [5]proton.Clickable

	// buttons
	btn1 proton.Clickable
	btn2 proton.Clickable
	btn3 proton.Clickable
	tap  proton.Clickable

	// inputs
	nameEd  proton.Editor
	bioEd   proton.Editor
	search  proton.SearchState

	// toggles
	darkMode  proton.Bool
	notify    proton.Bool
	plan      proton.Enum

	// slider
	vol     proton.Float
	rating  proton.NumberState

	// scroll / list
	listScroll proton.Scrollable

	// visual
	selectedColor int
	swatchBtns    [6]proton.Clickable

	// extra
	acc1    proton.AccordionState
	acc1btn proton.Clickable
	acc2    proton.AccordionState
	acc2btn proton.Clickable
	sel     proton.SelectBoxState
	modal   proton.OverlayState
	openBtn proton.Clickable
	closeBtn proton.Clickable
	spin    proton.SpinnerState

	// log
	logScroll proton.Scrollable
	logText   string
	logBtn    proton.Clickable

	// feedback
	toast    proton.ToastState
	widgetScroll proton.Scrollable
	clickCount int
}

var palette = []color.NRGBA{
	proton.RGB(0xf87171),
	proton.RGB(0xfbbf24),
	proton.RGB(0x4ade80),
	proton.RGB(0x60a5fa),
	proton.RGB(0xa78bfa),
	proton.RGB(0xf472b6),
}

var langs = []string{"Go", "Rust", "Zig", "C", "Python"}

func main() {
	u := &ui{selectedColor: 0}
	u.logText = "[OK] Showcase started\n"

	a := proton.New("showcase")
	a.ApplyPalette(proton.NordPalette)
	a.Window("Proton Widget Showcase", 820, 640, func(ctx proton.Context) {
		draw(ctx, u)
	})
	a.Run()
}

func draw(ctx proton.Context, u *ui) {
	proton.OnKey(ctx, proton.ModCtrl, "T", func() {
		u.toast.Show("Ctrl+T works!", 2*time.Second)
	})

	proton.Tabs(ctx,
		[]string{"Buttons & Input", "Toggles", "Visuals", "Extra", "Async"},
		u.tabBtns[:],
		&u.tabs,
		func(ctx proton.Context, i int) {
			proton.Gap(ctx, 12)
			switch i {
			case 0: tabButtons(ctx, u)
			case 1: tabToggles(ctx, u)
			case 2: tabVisuals(ctx, u)
			case 3: tabExtra(ctx, u)
			case 4: tabAsync(ctx, u)
			}
		},
	)
	proton.Toast(ctx, &u.toast)
}

func tabButtons(ctx proton.Context, u *ui) {
	proton.Scroll(ctx, &u.widgetScroll, func(ctx proton.Context) {
		proton.LabeledDivider(ctx, "Buttons")
		proton.Gap(ctx, 10)
		proton.Row(ctx,
			func(ctx proton.Context) {
				proton.Pad(ctx, 4, func(ctx proton.Context) {
					if proton.Button(ctx, &u.btn1, "Primary") {
						u.clickCount++
						u.toast.Show(fmt.Sprintf("Clicked %d times", u.clickCount), time.Second)
					}
				})
			},
			func(ctx proton.Context) { proton.Gap(ctx, 8) },
			func(ctx proton.Context) {
				proton.Pad(ctx, 4, func(ctx proton.Context) {
					if proton.OutlineButton(ctx, &u.btn2, "Outline") {
						u.toast.Show("Outline clicked", time.Second)
					}
				})
			},
			func(ctx proton.Context) { proton.Gap(ctx, 8) },
			func(ctx proton.Context) {
				if proton.Tappable(ctx, &u.tap, func(ctx proton.Context) {
					proton.Badge(ctx, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "Tappable badge")
				}) {
					u.toast.Show("Tappable clicked", time.Second)
				}
			},
		)
		proton.Gap(ctx, 6)
		proton.Caption(ctx, fmt.Sprintf("Button clicks: %d  •  Ctrl+T for keyboard shortcut", u.clickCount))

		proton.Gap(ctx, 20)
		proton.LabeledDivider(ctx, "Text inputs")
		proton.Gap(ctx, 10)
		proton.Label(ctx, "Name")
		proton.Gap(ctx, 4)
		proton.Input(ctx, &u.nameEd, "Your name…")
		proton.Gap(ctx, 12)
		proton.Label(ctx, "Bio")
		proton.Gap(ctx, 4)
		proton.TextArea(ctx, &u.bioEd, "Tell us something…")
		proton.Gap(ctx, 12)
		proton.Label(ctx, "Search")
		proton.Gap(ctx, 4)
		q := proton.SearchInput(ctx, &u.search, "Search anything…")
		proton.Gap(ctx, 4)
		if q != "" {
			proton.Caption(ctx, fmt.Sprintf("query: %q", q))
		}
		if u.nameEd.Text() != "" {
			proton.Gap(ctx, 4)
			proton.Caption(ctx, "Hello, "+u.nameEd.Text()+"!")
		}

		proton.Gap(ctx, 20)
		proton.LabeledDivider(ctx, "NumberInput")
		proton.Gap(ctx, 10)
		v := proton.NumberInput(ctx, &u.rating, 0, 10, 0.5)
		proton.Gap(ctx, 4)
		proton.Caption(ctx, fmt.Sprintf("Value: %.1f", v))
	})
}

func tabToggles(ctx proton.Context, u *ui) {
	proton.LabeledDivider(ctx, "Toggle (switch)")
	proton.Gap(ctx, 10)
	proton.Pad(ctx, 4, func(ctx proton.Context) {
		if proton.Toggle(ctx, &u.darkMode, "Dark mode") {
			u.toast.Show(fmt.Sprintf("Dark mode: %v", u.darkMode.Value), time.Second)
		}
	})
	proton.Gap(ctx, 8)
	proton.Pad(ctx, 4, func(ctx proton.Context) {
		proton.Toggle(ctx, &u.notify, "Notifications")
	})

	proton.Gap(ctx, 20)
	proton.LabeledDivider(ctx, "Checkbox")
	proton.Gap(ctx, 10)
	proton.Pad(ctx, 4, func(ctx proton.Context) {
		proton.Checkbox(ctx, &u.darkMode, "I agree to the terms")
	})
	proton.Gap(ctx, 8)
	proton.Pad(ctx, 4, func(ctx proton.Context) {
		proton.Checkbox(ctx, &u.notify, "Send me emails (nobody reads these)")
	})

	proton.Gap(ctx, 20)
	proton.LabeledDivider(ctx, "RadioButton")
	proton.Gap(ctx, 10)
	proton.Pad(ctx, 4, func(ctx proton.Context) { proton.RadioButton(ctx, &u.plan, "free", "Free") })
	proton.Gap(ctx, 6)
	proton.Pad(ctx, 4, func(ctx proton.Context) { proton.RadioButton(ctx, &u.plan, "pro", "Pro ($9/mo)") })
	proton.Gap(ctx, 6)
	proton.Pad(ctx, 4, func(ctx proton.Context) { proton.RadioButton(ctx, &u.plan, "team", "Team ($29/mo)") })
	proton.Gap(ctx, 8)
	proton.Caption(ctx, "Selected: "+u.plan.Value)

	proton.Gap(ctx, 20)
	proton.LabeledDivider(ctx, "Slider")
	proton.Gap(ctx, 10)
	v := proton.Slider(ctx, &u.vol)
	proton.Gap(ctx, 6)
	proton.Caption(ctx, fmt.Sprintf("Volume: %.0f%%", v*100))
	proton.Gap(ctx, 8)
	proton.ProgressBar(ctx, v)
}

func tabVisuals(ctx proton.Context, u *ui) {
	proton.LabeledDivider(ctx, "Colors & shapes")
	proton.Gap(ctx, 10)
	proton.Row(ctx,
		func(ctx proton.Context) { proton.Rect(ctx, proton.RGB(0x5e81ac), 60, 40) },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.RoundRect(ctx, proton.RGB(0xa3be8c), 60, 40, 10) },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.RoundRect(ctx, proton.RGB(0xbf616a), 60, 40, 20) },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Rect(ctx, proton.RGB(0xebcb8b), 60, 40) },
	)

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Cards & Badges")
	proton.Gap(ctx, 10)
	proton.Grid(ctx, 3, 8,
		func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x3b4252), 8, 12, func(ctx proton.Context) {
				proton.Label(ctx, "Card one")
				proton.Gap(ctx, 4)
				proton.Muted(ctx, "With shadow")
			})
		},
		func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x4c566a), 8, 12, func(ctx proton.Context) {
				proton.Label(ctx, "Card two")
				proton.Gap(ctx, 4)
				proton.Muted(ctx, "Lighter bg")
			})
		},
		func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x2e3440), 8, 12, func(ctx proton.Context) {
				proton.Label(ctx, "Card three")
				proton.Gap(ctx, 4)
				proton.Muted(ctx, "Darker bg")
			})
		},
	)
	proton.Gap(ctx, 12)
	proton.Row(ctx,
		func(ctx proton.Context) { proton.Badge(ctx, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "stable") },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Badge(ctx, proton.RGB(0xa3be8c), proton.RGB(0x2e3440), "passing") },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Badge(ctx, proton.RGB(0xbf616a), proton.RGB(0xeceff4), "failing") },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Badge(ctx, proton.RGB(0xebcb8b), proton.RGB(0x2e3440), "beta") },
	)

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Color swatch")
	proton.Gap(ctx, 10)
	i := proton.ColorSwatch(ctx, u.swatchBtns[:], palette, u.selectedColor, 28)
	if i >= 0 {
		u.selectedColor = i
	}
	proton.Gap(ctx, 6)
	proton.Row(ctx,
		func(ctx proton.Context) { proton.StatusDot(ctx, palette[u.selectedColor], 10) },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Caption(ctx, fmt.Sprintf("Color %d selected", u.selectedColor+1)) },
	)

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Progress rings")
	proton.Gap(ctx, 10)
	proton.Row(ctx,
		func(ctx proton.Context) { proton.ProgressRing(ctx, 0.25, 48, 5, proton.RGB(0x88c0d0)) },
		func(ctx proton.Context) { proton.Gap(ctx, 12) },
		func(ctx proton.Context) { proton.ProgressRing(ctx, 0.60, 48, 5, proton.RGB(0xa3be8c)) },
		func(ctx proton.Context) { proton.Gap(ctx, 12) },
		func(ctx proton.Context) { proton.ProgressRing(ctx, 0.90, 48, 5, proton.RGB(0xbf616a)) },
	)

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Alerts")
	proton.Gap(ctx, 8)
	proton.Alert(ctx, proton.AlertInfo, "Informational alert.")
	proton.Gap(ctx, 6)
	proton.Alert(ctx, proton.AlertSuccess, "Success alert.")
	proton.Gap(ctx, 6)
	proton.Alert(ctx, proton.AlertWarning, "Warning alert.")
	proton.Gap(ctx, 6)
	proton.Alert(ctx, proton.AlertError, "Error alert.")
}

func tabExtra(ctx proton.Context, u *ui) {
	proton.LabeledDivider(ctx, "SelectBox")
	proton.Gap(ctx, 8)
	proton.MaxWidth(ctx, 220, func(ctx proton.Context) {
		i := proton.SelectBox(ctx, &u.sel, langs)
		_ = i
	})
	if u.sel.Selected >= 0 && u.sel.Selected < len(langs) {
		proton.Gap(ctx, 4)
		proton.Caption(ctx, "Picked: "+langs[u.sel.Selected])
	}

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Accordion")
	proton.Gap(ctx, 6)
	proton.Accordion(ctx, &u.acc1, &u.acc1btn, "What is Proton?", func(ctx proton.Context) {
		proton.Gap(ctx, 4)
		proton.Label(ctx, "A pure-Go GUI library built on Gio.")
		proton.Gap(ctx, 4)
		proton.CodeBlock(ctx, "go get github.com/CzaxStudio/proton")
	})
	proton.Gap(ctx, 4)
	proton.Accordion(ctx, &u.acc2, &u.acc2btn, "Does it work on Windows?", func(ctx proton.Context) {
		proton.Gap(ctx, 4)
		proton.Label(ctx, "Yes — and macOS and Linux too.")
	})

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Overlay / modal")
	proton.Gap(ctx, 10)
	proton.Pad(ctx, 4, func(ctx proton.Context) {
		if proton.Button(ctx, &u.openBtn, "Open modal") {
			u.modal.Show()
		}
	})
	proton.Overlay(ctx, &u.modal, func(ctx proton.Context) {
		proton.MinSize(ctx, 280, 0, func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x2e3440), 12, 24, func(ctx proton.Context) {
				proton.H5(ctx, "Modal overlay")
				proton.Gap(ctx, 8)
				proton.Label(ctx, "Everything behind this is dimmed.")
				proton.Gap(ctx, 16)
				proton.Pad(ctx, 4, func(ctx proton.Context) {
					if proton.Button(ctx, &u.closeBtn, "Close") {
						u.modal.Hide()
					}
				})
			})
		})
	})

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Link, CodeBlock, ShortcutHint")
	proton.Gap(ctx, 8)
	proton.Row(ctx,
		func(ctx proton.Context) {
			if proton.Link(ctx, &u.btn3, "GitHub") {
				u.toast.Show("Link clicked", time.Second)
			}
		},
		func(ctx proton.Context) { proton.Gap(ctx, 16) },
		func(ctx proton.Context) { proton.ShortcutHint(ctx, "Ctrl+T") },
		func(ctx proton.Context) { proton.Gap(ctx, 6) },
		func(ctx proton.Context) { proton.Caption(ctx, "— try it!") },
	)
	proton.Gap(ctx, 8)
	proton.CodeBlock(ctx, `a.Window("App", 480, 300, func(ctx proton.Context) {
    proton.Label(ctx, "Hello!")
})`)

	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Table")
	proton.Gap(ctx, 8)
	proton.Table(ctx,
		[]string{"Name", "Role", "Status"},
		[]proton.TableRow{
			{"Alice", "Engineer", "Active"},
			{"Bob", "Designer", "Away"},
			{"Carol", "PM", "Active"},
		},
	)
}

func tabAsync(ctx proton.Context, u *ui) {
	proton.LabeledDivider(ctx, "Spinner (always running)")
	proton.Gap(ctx, 12)
	proton.Row(ctx,
		func(ctx proton.Context) { proton.Spinner(ctx, &u.spin, 32) },
		func(ctx proton.Context) { proton.Gap(ctx, 12) },
		func(ctx proton.Context) {
			proton.Column(ctx,
				func(ctx proton.Context) { proton.Label(ctx, "Loading widget") },
				func(ctx proton.Context) { proton.Gap(ctx, 4) },
				func(ctx proton.Context) { proton.Muted(ctx, "Animates independently — no Invalidate loop needed") },
			)
		},
	)

	proton.Gap(ctx, 20)
	proton.LabeledDivider(ctx, "LogView + goroutine")
	proton.Gap(ctx, 10)
	proton.Row(ctx,
		func(ctx proton.Context) {
			proton.Pad(ctx, 4, func(ctx proton.Context) {
				if proton.Button(ctx, &u.logBtn, "Add log line") {
					u.logText += fmt.Sprintf("[OK] Entry at %s\n", time.Now().Format("15:04:05.000"))
				}
			})
		},
	)
	proton.Gap(ctx, 8)
	proton.MinSize(ctx, 0, 180, func(ctx proton.Context) {
		proton.Card(ctx, proton.RGB(0x2e3440), 6, 8, func(ctx proton.Context) {
			proton.LogView(ctx, &u.logScroll, u.logText)
		})
	})

	proton.Gap(ctx, 20)
	proton.LabeledDivider(ctx, "Toast")
	proton.Gap(ctx, 10)
	proton.Row(ctx,
		func(ctx proton.Context) {
			proton.Pad(ctx, 4, func(ctx proton.Context) {
				if proton.Button(ctx, &u.btn2, "Show toast") {
					u.toast.Show("Toast is working!", 2*time.Second)
				}
			})
		},
	)
}
