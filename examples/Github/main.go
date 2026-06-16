package main

import (
	"time"

	"github.com/CzaxStudio/proton"
)

type UI struct {
	// Sidebar
	tabs    proton.TabState
	tabBtns []proton.Clickable

	// Search
	searchEditor proton.Editor
	searchBtn    proton.Clickable

	// Repos
	repoBtns []proton.Clickable

	// Activity
	activityScroll proton.Scrollable
	activityLog    string

	// Feedback + state
	toast proton.ToastState
	spin  proton.SpinnerState

	app *proton.App
}

func main() {
	ui := &UI{
		tabBtns:  make([]proton.Clickable, 4),
		repoBtns: make([]proton.Clickable, 6),
	}

	ui.activityLog = "[INIT] Welcome to Github Studio\n"
	ui.searchEditor.SetText("")

	ui.app = proton.New("Github Studio — Proton Showcase")
	ui.app.ApplyPalette(proton.CatppuccinPalette)
	ui.app.SetFontScale(1.02)

	ui.app.Window("Github Studio", 1100, 700, func(win *proton.Win) {
		drawApp(win, ui)
	})

	ui.app.Run()
}

func drawApp(win *proton.Win, ui *UI) {
	// hero + content split
	proton.Column(win,
		func(win *proton.Win) {
			// Hero
			proton.Card(win, proton.RGBA(10, 10, 20, 160), 12, 20, func(win *proton.Win) {
				proton.Row(win,
					func(win *proton.Win) {
						proton.Column(win,
							func(win *proton.Win) { proton.H1(win, "Github Studio") },
							func(win *proton.Win) { proton.H3(win, "A beautiful showcase built with Proton") },
						)
					},
					func(win *proton.Win) {
						proton.GrowItem(win, func(win *proton.Win) {
							proton.PadH(win, 12, func(win *proton.Win) {
								proton.RowEnd(win,
									func(win *proton.Win) {
										proton.Input(win, &ui.searchEditor, "Search repositories, issues, users...")
									},
									func(win *proton.Win) {
										proton.Gap(win, 8)
										if proton.Button(win, &ui.searchBtn, "Search") {
											ui.activityLog += "[SEARCH] Query executed\n"
											ui.toast.Show("Search complete", 2*time.Second)
										}
									},
								)
							})
						})
					},
				)
			})
		},
		func(win *proton.Win) {
			// Main split: left repos, right activity
			proton.Split(win, 0.68,
				func(win *proton.Win) { drawRepos(win, ui) },
				func(win *proton.Win) { drawActivity(win, ui) },
			)
		},
	)

	// toast & spinner overlay
	proton.Center(win, func(win *proton.Win) { proton.Toast(win, &ui.toast) })
}

func drawRepos(win *proton.Win, ui *UI) {
	proton.Pad(win, 18, func(win *proton.Win) {
		proton.RowSpread(win,
			func(win *proton.Win) { proton.H4(win, "Featured Repositories") },
			func(win *proton.Win) {
				proton.OutlineButton(win, &ui.searchBtn, "New Repository")
			},
		)

		proton.Gap(win, 18)

		// showcase repo cards
		proton.Grid(win, 2, 16,
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[0], "proton", "A gorgeous UI toolkit", 12400) },
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[1], "proton-extras", "Widgets & patterns", 4920) },
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[2], "proton-themes", "Stunning palettes", 2100) },
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[3], "proton-examples", "Cookbook of examples", 890) },
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[4], "proton-cmd", "CLI tooling", 340) },
			func(win *proton.Win) { repoCard(win, &ui.repoBtns[5], "proton-showcase", "Design-first apps", 76) },
		)
	})
}

func repoCard(win *proton.Win, btn *proton.Clickable, name, desc string, stars int) {
	proton.Card(win, proton.RGBA(255, 255, 255, 6), 12, 16, func(win *proton.Win) {
		proton.Row(win,
			func(win *proton.Win) {
				proton.Column(win,
					func(win *proton.Win) { proton.Text(win, name, 18, win.Theme().Palette.ContrastBg, true) },
					func(win *proton.Win) { proton.Muted(win, desc) },
				)
			},
			func(win *proton.Win) {
				proton.GrowItem(win, func(win *proton.Win) {
					proton.RowEnd(win,
						func(win *proton.Win) {
							proton.Badge(win, proton.RGB(0xFFD166), proton.RGB(0x000000), "Star ")
							proton.Gap(win, 8)
							if proton.Button(win, (*proton.Clickable)(btn), "Open") {
								win.Invalidate()
							}
						},
					)
				})
			},
		)
		proton.Gap(win, 12)
		proton.ProgressBar(win, float32((stars%1000))/1000.0)
	})
}

func drawActivity(win *proton.Win, ui *UI) {
	proton.Pad(win, 16, func(win *proton.Win) {
		proton.Column(win,
			func(win *proton.Win) { proton.H5(win, "Activity Feed") },
			func(win *proton.Win) {
				// simulate activity
				ui.activityLog += "[EVENT] Repo indexed\n"
				ui.activityLog += "[EVENT] CI passed\n"
				proton.Card(win, proton.RGBA(255, 255, 255, 6), 8, 12, func(win *proton.Win) {
					proton.LogView(win, &ui.activityScroll, ui.activityLog)
				})
			},
			func(win *proton.Win) {
				proton.Gap(win, 12)
				proton.H6(win, "Now Scanning")
				proton.Row(win,
					func(win *proton.Win) {
						proton.Spinner(win, &ui.spin, 40)
					},
					func(win *proton.Win) {
						proton.Gap(win, 12)
						proton.Body2(win, "Background indexing in progress — lightweight and fast.")
					},
				)
			},
		)
	})
}
