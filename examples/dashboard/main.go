// dashboard — a developer stats dashboard.
// Demonstrates: Grid, ProgressRing, Table, LogView, Tabs, Stepper,
// ResizeSplit, StatusDot, Toast, Spinner, Alert, custom-drawn widgets.
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/CzaxStudio/proton"
)

type ui struct {
	tabs    proton.TabState
	tabBtns [3]proton.Clickable

	split      proton.ResizeSplitState
	refreshBtn proton.Clickable
	spinner    proton.SpinnerState
	loading    bool

	logScroll proton.Scrollable
	logText   string
	addLogBtn proton.Clickable
	clearBtn  proton.Clickable

	stepperStep int
	nextBtn     proton.Clickable
	backBtn     proton.Clickable

	toast     proton.ToastState
	showAlert bool

	cpuHist []float32
}

var servers = []proton.TableRow{
	{"us-east-1", "Online", "23%", "12 ms"},
	{"us-east-2", "Online", "82%", "11 ms"},
	{"eu-west-1", "Online", "41%", "38 ms"},
	{"ap-south-1", "Offline", "—", "—"},
}

var builds = []proton.TableRow{
	{"main", "#1042", "2m ago", "passed"},
	{"feature", "#1041", "18m ago", "passed"},
	{"hotfix", "#1040", "1h ago", "failed"},
	{"release", "#1039", "3h ago", "passed"},
}

func newUI() *ui {
	u := &ui{showAlert: true}
	u.logText = "[OK] Dashboard started\n[OK] Connected to API\n[WARN] us-east-2 CPU at 82%\n"
	u.cpuHist = make([]float32, 28)
	for i := range u.cpuHist {
		u.cpuHist[i] = 0.15 + rand.Float32()*0.7
	}
	return u
}

var gApp *proton.App

func main() {
	u := newUI()
	gApp = proton.New("dashboard")
	gApp.ApplyPalette(proton.NordPalette)
	gApp.Window("Dev Dashboard", 980, 640, func(ctx proton.Context) {
		draw(ctx, u)
	})
	gApp.Run()
}

func draw(ctx proton.Context, u *ui) {
	proton.RowSpread(ctx,
		func(ctx proton.Context) {
			proton.Row(ctx,
				func(ctx proton.Context) { proton.H5(ctx, "Dashboard") },
				func(ctx proton.Context) { proton.Gap(ctx, 10) },
				func(ctx proton.Context) { proton.StatusDot(ctx, proton.RGB(0x4ade80), 8) },
				func(ctx proton.Context) { proton.Gap(ctx, 4) },
				func(ctx proton.Context) { proton.Muted(ctx, "All systems operational") },
			)
		},
		func(ctx proton.Context) {
			proton.Pad(ctx, 4, func(ctx proton.Context) {
				if proton.Button(ctx, &u.refreshBtn, "Refresh") && !u.loading {
					u.loading = true
					go func() {
						time.Sleep(1200 * time.Millisecond)
						u.loading = false
						for i := range u.cpuHist {
							u.cpuHist[i] = 0.15 + rand.Float32()*0.7
						}
						u.logText += fmt.Sprintf("[OK] Refreshed at %s\n", time.Now().Format("15:04:05"))
						u.toast.Show("Data refreshed.", 2*time.Second)
						ctx.Invalidate()
					}()
				}
			})
		},
	)
	proton.Gap(ctx, 4)
	if u.loading {
		proton.Row(ctx,
			func(ctx proton.Context) { proton.Spinner(ctx, &u.spinner, 16) },
			func(ctx proton.Context) { proton.Gap(ctx, 6) },
			func(ctx proton.Context) { proton.Muted(ctx, "Refreshing...") },
		)
		proton.Gap(ctx, 4)
	}
	proton.Divider(ctx)
	proton.Gap(ctx, 8)

	proton.Tabs(ctx,
		[]string{"Overview", "Builds & Logs", "Servers"},
		u.tabBtns[:],
		&u.tabs,
		func(ctx proton.Context, i int) {
			proton.Gap(ctx, 12)
			switch i {
			case 0:
				overview(ctx, u)
			case 1:
				buildsTab(ctx, u)
			case 2:
				serversTab(ctx, u)
			}
		},
	)
	proton.Toast(ctx, &u.toast)
}

func overview(ctx proton.Context, u *ui) {
	if u.showAlert {
		proton.Alert(ctx, proton.AlertWarning, "Server us-east-2 is at 82% CPU. Consider scaling.")
		proton.Gap(ctx, 10)
	}

	type card struct {
		label string
		value string
		sub   string
		ring  float32
		col   proton.Palette
	}
	cards := []card{
		{"Uptime", "99.97%", "30-day avg", 0.9997, proton.EverforestDarkPalette},
		{"Requests", "1.24M", "Last 24h", 0.62, proton.NordPalette},
		{"Error Rate", "0.03%", "Well below SLA", 0.0003, proton.CatppuccinPalette},
		{"P95 Latency", "42 ms", "All regions", 0.42, proton.RosePinePalette},
	}
	fns := make([]func(proton.Context), len(cards))
	for i, c := range cards {
		c := c
		fns[i] = func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x3b4252), 10, 14, func(ctx proton.Context) {
				proton.Row(ctx,
					func(ctx proton.Context) {
						proton.Column(ctx,
							func(ctx proton.Context) { proton.Caption(ctx, c.label) },
							func(ctx proton.Context) { proton.Gap(ctx, 6) },
							func(ctx proton.Context) { proton.H4(ctx, c.value) },
							func(ctx proton.Context) { proton.Gap(ctx, 2) },
							func(ctx proton.Context) { proton.Muted(ctx, c.sub) },
						)
					},
					func(ctx proton.Context) { proton.Gap(ctx, 12) },
					func(ctx proton.Context) {
						proton.ProgressRing(ctx, c.ring, 50, 5, c.col.Primary)
					},
				)
			})
		}
	}
	proton.Grid(ctx, 2, 10, fns...)
	proton.Gap(ctx, 20)

	proton.LabeledDivider(ctx, "CPU history — 28 samples")
	proton.Gap(ctx, 10)
	proton.Card(ctx, proton.RGB(0x3b4252), 8, 12, func(ctx proton.Context) {
		proton.Row(ctx, func(ctx proton.Context) {
			for _, v := range u.cpuHist {
				v := v
				c := proton.RGB(0x88c0d0)
				if v > 0.75 {
					c = proton.RGB(0xbf616a)
				} else if v > 0.5 {
					c = proton.RGB(0xebcb8b)
				}
				proton.Column(ctx,
					func(ctx proton.Context) { proton.Rect(ctx, c, 10, v*56) },
				)
				proton.Gap(ctx, 3)
			}
		})
	})
	proton.Gap(ctx, 20)

	proton.LabeledDivider(ctx, "Deploy pipeline")
	proton.Gap(ctx, 12)
	proton.Stepper(ctx, u.stepperStep, []string{"Build", "Test", "Stage", "Deploy"})
	proton.Gap(ctx, 14)
	proton.Row(ctx,
		func(ctx proton.Context) {
			proton.Pad(ctx, 4, func(ctx proton.Context) {
				if proton.OutlineButton(ctx, &u.backBtn, "Back") && u.stepperStep > 0 {
					u.stepperStep--
				}
			})
		},
		func(ctx proton.Context) { proton.Gap(ctx, 8) },
		func(ctx proton.Context) {
			proton.Pad(ctx, 4, func(ctx proton.Context) {
				lbl := "Next"
				if u.stepperStep == 3 {
					lbl = "Deploy Now"
				}
				if proton.Button(ctx, &u.nextBtn, lbl) {
					if u.stepperStep < 3 {
						u.stepperStep++
					} else {
						u.stepperStep = 0
						u.logText += "[OK] Deployed to production\n"
						u.toast.Show("Deployed!", 3*time.Second)
					}
				}
			})
		},
	)
}

func buildsTab(ctx proton.Context, u *ui) {
	proton.ResizeSplit(ctx, &u.split, 0.5,
		func(ctx proton.Context) {
			proton.H6(ctx, "Recent builds")
			proton.Gap(ctx, 8)
			proton.Table(ctx, []string{"Branch", "Build", "When", "Status"}, builds)
		},
		func(ctx proton.Context) {
			proton.PadH(ctx, 12, func(ctx proton.Context) {
				proton.Row(ctx,
					func(ctx proton.Context) { proton.H6(ctx, "Build log") },
					func(ctx proton.Context) { proton.Gap(ctx, 8) },
					func(ctx proton.Context) {
						proton.Pad(ctx, 2, func(ctx proton.Context) {
							if proton.Button(ctx, &u.addLogBtn, "+ Line") {
								u.logText += fmt.Sprintf("[OK] Step %d at %s\n",
									rand.Intn(20)+1, time.Now().Format("15:04:05"))
							}
						})
					},
					func(ctx proton.Context) { proton.Gap(ctx, 6) },
					func(ctx proton.Context) {
						proton.Pad(ctx, 2, func(ctx proton.Context) {
							if proton.OutlineButton(ctx, &u.clearBtn, "Clear") {
								u.logText = ""
							}
						})
					},
				)
				proton.Gap(ctx, 8)
				proton.MinSize(ctx, 0, 220, func(ctx proton.Context) {
					proton.Card(ctx, proton.RGB(0x2e3440), 6, 8, func(ctx proton.Context) {
						proton.LogView(ctx, &u.logScroll, u.logText)
					})
				})
			})
		},
	)
}

func serversTab(ctx proton.Context, u *ui) {
	proton.H6(ctx, "Servers")
	proton.Gap(ctx, 12)
	for _, s := range servers {
		s := s
		proton.PadV(ctx, 4, func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x3b4252), 6, 12, func(ctx proton.Context) {
				proton.Row(ctx,
					func(ctx proton.Context) {
						c := proton.RGB(0x4ade80)
						if s[1] != "Online" {
							c = proton.RGB(0xf87171)
						}
						proton.StatusDot(ctx, c, 8)
					},
					func(ctx proton.Context) { proton.Gap(ctx, 10) },
					func(ctx proton.Context) { proton.Label(ctx, s[0]) },
					func(ctx proton.Context) { proton.Gap(ctx, 20) },
					func(ctx proton.Context) { proton.Muted(ctx, "CPU: "+s[2]) },
					func(ctx proton.Context) { proton.Gap(ctx, 20) },
					func(ctx proton.Context) { proton.Muted(ctx, "Ping: "+s[3]) },
				)
			})
		})
	}
	proton.Gap(ctx, 16)
	proton.LabeledDivider(ctx, "Table view")
	proton.Gap(ctx, 8)
	proton.Table(ctx, []string{"Region", "Status", "CPU", "Latency"}, servers)
}
