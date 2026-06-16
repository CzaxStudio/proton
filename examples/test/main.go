package main

import (
	"fmt"
	"time"

	"github.com/CzaxStudio/proton"
)

type ui struct {
	// buttons
	primary  proton.Clickable
	outline  proton.Clickable
	tappable proton.Clickable
	incBtn   proton.Clickable

	// inputs
	name    proton.Editor
	message proton.Editor

	// toggles / checks
	darkMode  proton.Bool
	notify    proton.Bool
	colorPref proton.Enum

	// slider / progress
	vol      proton.Float
	progress float32

	// list
	scroll proton.Scrollable
	items  []string

	// scroll region
	contentScroll proton.Scrollable

	// toast
	toast proton.ToastState

	// status line
	status string
}

func main() {
	u := &ui{
		items:    []string{"Apples", "Bananas", "Cherries", "Dates", "Elderberries", "Figs", "Grapes"},
		progress: 0.4,
		status:   "Ready.",
	}

	a := proton.New("showcase")
	a.ApplyPalette(proton.NordPalette)
	a.Window("Proton Showcase", 720, 560, func(win *proton.Win) {
		draw(win, u)
	})
	a.Run()
}

func draw(win *proton.Win, u *ui) {
	// top bar
	proton.PadSides(win, 0, 0, 8, 0, func(win *proton.Win) {
		proton.RowSpread(win,
			func(win *proton.Win) { proton.H4(win, "Proton Showcase") },
			func(win *proton.Win) { proton.Caption(win, "v1.0") },
		)
	})
	proton.Divider(win)
	proton.Gap(win, 12)

	// main two-column layout
	proton.Split(win, 0.45,
		func(win *proton.Win) { leftPanel(win, u) },
		func(win *proton.Win) {
			proton.PadH(win, 12, func(win *proton.Win) {
				rightPanel(win, u)
			})
		},
	)

	proton.Gap(win, 8)
	proton.Divider(win)
	proton.Gap(win, 4)
	proton.Caption(win, u.status)

	// toast on top of everything
	proton.Toast(win, &u.toast)
}

func leftPanel(win *proton.Win, u *ui) {
	proton.H6(win, "Buttons")
	proton.Gap(win, 6)

	proton.Row(win,
		func(win *proton.Win) {
			if proton.Button(win, &u.primary, "Primary") {
				u.status = "Primary clicked."
				u.toast.Show("Primary clicked!", time.Second*2)
			}
		},
		func(win *proton.Win) { proton.Gap(win, 8) },
		func(win *proton.Win) {
			if proton.OutlineButton(win, &u.outline, "Outline") {
				u.status = "Outline clicked."
			}
		},
	)
	proton.Gap(win, 8)

	proton.H6(win, "Input")
	proton.Gap(win, 4)
	proton.Input(win, &u.name, "Your name")
	proton.Gap(win, 4)
	proton.TextArea(win, &u.message, "Write something...")
	proton.Gap(win, 4)
	proton.Caption(win, fmt.Sprintf("name: %q", u.name.Text()))

	proton.Gap(win, 12)
	proton.H6(win, "Toggles & Radio")
	proton.Gap(win, 4)

	proton.Toggle(win, &u.darkMode, "Dark mode")
	proton.Gap(win, 4)
	proton.Checkbox(win, &u.notify, "Enable notifications")
	proton.Gap(win, 8)
	proton.RadioButton(win, &u.colorPref, "red", "Red")
	proton.RadioButton(win, &u.colorPref, "green", "Green")
	proton.RadioButton(win, &u.colorPref, "blue", "Blue")
	proton.Gap(win, 4)
	proton.Caption(win, "color: "+u.colorPref.Value)
}

func rightPanel(win *proton.Win, u *ui) {
	proton.H6(win, "Slider & Progress")
	proton.Gap(win, 6)

	v := proton.Slider(win, &u.vol)
	proton.Gap(win, 4)
	proton.Caption(win, fmt.Sprintf("volume: %.0f%%", v*100))
	proton.Gap(win, 6)
	proton.ProgressBar(win, u.progress)
	proton.Gap(win, 4)
	proton.Row(win,
		func(win *proton.Win) {
			if proton.Button(win, &u.incBtn, "+10%") {
				if u.progress < 1 {
					u.progress += 0.1
				}
			}
		},
	)

	proton.Gap(win, 12)
	proton.H6(win, "List")
	proton.Gap(win, 4)

	proton.Card(win, proton.RGB(0x3b4252), 8, 0, func(win *proton.Win) {
		proton.List(win, &u.scroll, len(u.items), func(win *proton.Win, i int) {
			proton.PadV(win, 5, func(win *proton.Win) {
				proton.PadH(win, 8, func(win *proton.Win) {
					proton.Label(win, u.items[i])
				})
			})
		})
	})

	proton.Gap(win, 12)
	proton.H6(win, "Grid")
	proton.Gap(win, 6)

	proton.Grid(win, 3, 6,
		func(win *proton.Win) { proton.Badge(win, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "Go") },
		func(win *proton.Win) { proton.Badge(win, proton.RGB(0xa3be8c), proton.RGB(0x2e3440), "Pure") },
		func(win *proton.Win) { proton.Badge(win, proton.RGB(0xbf616a), proton.RGB(0xeceff4), "Fast") },
	)
}
