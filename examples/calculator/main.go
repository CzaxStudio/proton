package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/CzaxStudio/proton"
)

type ui struct {
	tabs    proton.TabState
	tabBtns [5]proton.Clickable

	// accordion
	acc1    proton.AccordionState
	acc1btn proton.Clickable
	acc2    proton.AccordionState
	acc2btn proton.Clickable

	// spinner
	spin proton.SpinnerState

	// selectbox
	langSel proton.SelectBoxState

	// overlay
	modal    proton.OverlayState
	openBtn  proton.Clickable
	closeBtn proton.Clickable
	confBtn  proton.Clickable

	// number input
	qty    proton.NumberState
	rating proton.NumberState

	// hover cards
	cards        [4]proton.Clickable
	selectedCard int

	// color swatch
	swatches      [6]proton.Clickable
	selectedColor int

	// alerts
	showInfo    bool
	showWarn    bool
	dismissInfo proton.Clickable
	dismissWarn proton.Clickable

	// links
	githubLink proton.Clickable
	docsLink   proton.Clickable

	// log view
	logScroll proton.Scrollable
	logText   string
	logBtn    proton.Clickable

	// resizable split
	split proton.ResizeSplitState

	// toast
	toast proton.ToastState

	// scroll for widget tab
	widgetScroll proton.Scrollable

	// general
	loading bool
	loadBtn proton.Clickable
}

var langs = []string{"Go", "Rust", "Zig", "C", "Python", "TypeScript"}

var palette = []color.NRGBA{
	proton.RGB(0xf87171),
	proton.RGB(0xfbbf24),
	proton.RGB(0x4ade80),
	proton.RGB(0x60a5fa),
	proton.RGB(0xa78bfa),
	proton.RGB(0xf472b6),
}

func main() {
	u := &ui{
		selectedCard:  -1,
		selectedColor: -1,
		showInfo:      true,
		showWarn:      true,
	}
	u.logText = "[OK] Proton initialized\n"

	a := proton.New("kitchen")
	a.ApplyPalette(proton.NordPalette)
	a.Window("Proton Kitchen Sink", 800, 620, func(win *proton.Win) {
		draw(win, u)
	})
	a.Run()
}

func draw(win *proton.Win, u *ui) {
	proton.Tabs(win,
		[]string{"Widgets", "Text", "Layout", "Async", "Overlay"},
		u.tabBtns[:],
		&u.tabs,
		func(win *proton.Win, i int) {
			switch i {
			case 0:
				tabWidgets(win, u)
			case 1:
				tabText(win, u)
			case 2:
				tabLayout(win, u)
			case 3:
				tabAsync(win, u)
			case 4:
				tabOverlay(win, u)
			}
		},
	)
	proton.Toast(win, &u.toast)
}

func tabWidgets(win *proton.Win, u *ui) {
	proton.Scroll(win, &u.widgetScroll, func(win *proton.Win) {
		proton.LabeledDivider(win, "SelectBox")
		proton.Gap(win, 8)
		proton.MaxWidth(win, 240, func(win *proton.Win) {
			proton.SelectBox(win, &u.langSel, langs)
		})
		proton.Gap(win, 4)
		if u.langSel.Selected >= 0 && u.langSel.Selected < len(langs) {
			proton.Caption(win, "Selected: "+langs[u.langSel.Selected])
		}

		proton.Gap(win, 16)
		proton.LabeledDivider(win, "NumberInput")
		proton.Gap(win, 8)
		proton.Row(win,
			func(win *proton.Win) {
				proton.Column(win,
					func(win *proton.Win) { proton.Caption(win, "Quantity") },
					func(win *proton.Win) { proton.Gap(win, 4) },
					func(win *proton.Win) { proton.NumberInput(win, &u.qty, 1, 99, 1) },
				)
			},
			func(win *proton.Win) { proton.Gap(win, 24) },
			func(win *proton.Win) {
				proton.Column(win,
					func(win *proton.Win) { proton.Caption(win, "Rating (0.5 steps)") },
					func(win *proton.Win) { proton.Gap(win, 4) },
					func(win *proton.Win) { proton.NumberInput(win, &u.rating, 0, 5, 0.5) },
				)
			},
		)

		proton.Gap(win, 16)
		proton.LabeledDivider(win, "ColorSwatch")
		proton.Gap(win, 8)
		i := proton.ColorSwatch(win, u.swatches[:], palette, u.selectedColor, 26)
		if i >= 0 {
			u.selectedColor = i
		}
		proton.Gap(win, 4)
		if u.selectedColor >= 0 {
			proton.Row(win,
				func(win *proton.Win) {
					proton.StatusDot(win, palette[u.selectedColor], 10)
				},
				func(win *proton.Win) { proton.Gap(win, 6) },
				func(win *proton.Win) {
					proton.Caption(win, fmt.Sprintf("Color %d selected", u.selectedColor+1))
				},
			)
		}

		proton.Gap(win, 16)
		proton.LabeledDivider(win, "Accordion")
		proton.Gap(win, 6)
		proton.Accordion(win, &u.acc1, &u.acc1btn, "What is Proton?", func(win *proton.Win) {
			proton.Gap(win, 4)
			proton.Label(win, "A GUI library for Go. No C deps. Just go build.")
		})
		proton.Gap(win, 2)
		proton.Accordion(win, &u.acc2, &u.acc2btn, "Does it support Windows?", func(win *proton.Win) {
			proton.Gap(win, 4)
			proton.Label(win, "Yes. macOS and Linux too. One codebase.")
		})

		proton.Gap(win, 16)
		proton.LabeledDivider(win, "Status Dots")
		proton.Gap(win, 8)
		statuses := []struct {
			label string
			c     color.NRGBA
		}{
			{"Online", proton.RGB(0x4ade80)},
			{"Away", proton.RGB(0xfbbf24)},
			{"Offline", proton.RGB(0xf87171)},
		}
		proton.Row(win, func(win *proton.Win) {
			for _, s := range statuses {
				s := s
				proton.Row(win,
					func(win *proton.Win) { proton.StatusDot(win, s.c, 9) },
					func(win *proton.Win) { proton.Gap(win, 5) },
					func(win *proton.Win) { proton.Caption(win, s.label) },
					func(win *proton.Win) { proton.Gap(win, 16) },
				)
			}
		})
	})
}

func tabText(win *proton.Win, u *ui) {
	proton.LabeledDivider(win, "Alerts")
	proton.Gap(win, 8)

	proton.Alert(win, proton.AlertInfo, "This is an info alert. Something useful happened.")
	proton.Gap(win, 6)
	proton.Alert(win, proton.AlertSuccess, "Operation completed successfully.")
	proton.Gap(win, 6)
	proton.Alert(win, proton.AlertWarning, "Proceed with caution. This may have side effects.")
	proton.Gap(win, 6)
	proton.Alert(win, proton.AlertError, "Something went wrong. Check the logs.")
	proton.Gap(win, 6)

	if u.showInfo {
		if proton.AlertDismissable(win, &u.dismissInfo, proton.AlertInfo, "Click × to dismiss this alert.") {
			u.showInfo = false
		}
		proton.Gap(win, 6)
	}

	proton.Gap(win, 16)
	proton.LabeledDivider(win, "Text helpers")
	proton.Gap(win, 8)

	proton.ErrorText(win, "This is an error message.")
	proton.Gap(win, 4)
	proton.SuccessText(win, "This is a success message.")
	proton.Gap(win, 4)
	proton.WarningText(win, "This is a warning message.")
	proton.Gap(win, 4)
	proton.Muted(win, "This is muted/secondary text.")

	proton.Gap(win, 16)
	proton.LabeledDivider(win, "Links")
	proton.Gap(win, 8)

	proton.Row(win,
		func(win *proton.Win) {
			if proton.Link(win, &u.githubLink, "View on GitHub") {
				u.toast.Show("Would open github.com/CzaxStudio/proton", 2*time.Second)
			}
		},
		func(win *proton.Win) { proton.Gap(win, 16) },
		func(win *proton.Win) {
			if proton.LinkSmall(win, &u.docsLink, "Read the docs") {
				u.toast.Show("Would open docs", 2*time.Second)
			}
		},
	)

	proton.Gap(win, 16)
	proton.LabeledDivider(win, "CodeBlock")
	proton.Gap(win, 8)
	proton.CodeBlock(win, "go get github.com/CzaxStudio/proton")
	proton.Gap(win, 6)
	proton.CodeBlock(win, `a := proton.New("my app")
a.ApplyPalette(proton.NordPalette)
a.Window("Hello", 480, 300, draw)
a.Run()`)

	proton.Gap(win, 16)
	proton.LabeledDivider(win, "ShortcutHint")
	proton.Gap(win, 8)
	proton.Row(win,
		func(win *proton.Win) { proton.Label(win, "Save file") },
		func(win *proton.Win) { proton.Gap(win, 8) },
		func(win *proton.Win) { proton.ShortcutHint(win, "Ctrl+S") },
		func(win *proton.Win) { proton.Gap(win, 24) },
		func(win *proton.Win) { proton.Label(win, "Undo") },
		func(win *proton.Win) { proton.Gap(win, 8) },
		func(win *proton.Win) { proton.ShortcutHint(win, "Ctrl+Z") },
	)
}

func tabLayout(win *proton.Win, u *ui) {
	proton.LabeledDivider(win, "ResizeSplit — drag the handle")
	proton.Gap(win, 8)

	proton.MinSize(win, 0, 200, func(win *proton.Win) {
		proton.ResizeSplit(win, &u.split, 0.35,
			func(win *proton.Win) {
				proton.Pad(win, 8, func(win *proton.Win) {
					proton.Caption(win, "Left pane")
					proton.Gap(win, 8)
					proton.Label(win, "Drag the handle to resize.")
				})
			},
			func(win *proton.Win) {
				proton.Pad(win, 8, func(win *proton.Win) {
					proton.Caption(win, "Right pane")
					proton.Gap(win, 8)
					proton.Label(win, "Both panes are fully independent.")
				})
			},
		)
	})

	proton.Gap(win, 20)
	proton.LabeledDivider(win, "ZStack")
	proton.Gap(win, 8)

	proton.ZStack(win,
		func(win *proton.Win) {
			proton.RoundRect(win, proton.RGB(0x3b4252), 0, 80, 10)
		},
		func(win *proton.Win) {
			proton.Center(win, func(win *proton.Win) {
				proton.Label(win, "Text rendered on top of a shape")
			})
		},
	)

	proton.Gap(win, 20)
	proton.LabeledDivider(win, "FlexSpacer")
	proton.Gap(win, 8)

	proton.GrowRow(win,
		proton.FixedItem(win, func(win *proton.Win) { proton.Caption(win, "pushed left") }),
		proton.FlexSpacer(),
		proton.FixedItem(win, func(win *proton.Win) { proton.Caption(win, "pushed right") }),
	)
}

func tabAsync(win *proton.Win, u *ui) {
	proton.LabeledDivider(win, "Spinner")
	proton.Gap(win, 12)
	proton.Center(win, func(win *proton.Win) {
		proton.Spinner(win, &u.spin, 44)
	})

	proton.Gap(win, 20)
	proton.LabeledDivider(win, "Log viewer")
	proton.Gap(win, 8)

	proton.Pad(win, 4, func(win *proton.Win) {
		if proton.Button(win, &u.logBtn, "Add log line") {
			u.logText += fmt.Sprintf("[OK] Event at %s\n", time.Now().Format("15:04:05"))
		}
	})
	proton.Gap(win, 8)

	proton.MinSize(win, 0, 160, func(win *proton.Win) {
		proton.Card(win, proton.RGB(0x2e3440), 6, 0, func(win *proton.Win) {
			proton.Pad(win, 8, func(win *proton.Win) {
				proton.LogView(win, &u.logScroll, u.logText)
			})
		})
	})
}

func tabOverlay(win *proton.Win, u *ui) {
	proton.Label(win, "Click the button below to open a modal overlay.")
	proton.Gap(win, 16)

	proton.Pad(win, 4, func(win *proton.Win) {
		if proton.Button(win, &u.openBtn, "Open Modal") {
			u.modal.Show()
		}
	})

	proton.Overlay(win, &u.modal, func(win *proton.Win) {
		proton.MinSize(win, 300, 0, func(win *proton.Win) {
			proton.Card(win, proton.RGB(0x2e3440), 12, 24, func(win *proton.Win) {
				proton.H5(win, "Confirm")
				proton.Gap(win, 8)
				proton.Label(win, "Are you sure you want to do this?")
				proton.Gap(win, 4)
				proton.Muted(win, "This action cannot be undone.")
				proton.Gap(win, 20)
				proton.RowEnd(win,
					func(win *proton.Win) {
						proton.Pad(win, 4, func(win *proton.Win) {
							if proton.OutlineButton(win, &u.closeBtn, "Cancel") {
								u.modal.Hide()
							}
						})
					},
					func(win *proton.Win) { proton.Gap(win, 8) },
					func(win *proton.Win) {
						proton.Pad(win, 4, func(win *proton.Win) {
							if proton.Button(win, &u.confBtn, "Confirm") {
								u.modal.Hide()
								u.toast.Show("Done.", 2*time.Second)
							}
						})
					},
				)
			})
		})
	})
}
