package main

import "github.com/CzaxStudio/proton"

type ui struct {
	picker proton.ThemePickerState
	btn    proton.Clickable
	check  proton.Bool
	input  proton.Editor
	vol    proton.Float
}

var app *proton.App

func main() {
	u := &ui{}
	app = proton.New("theme picker")
	app.ApplyPalette(proton.NordPalette)
	app.Window("Theme Picker", 560, 600, func(win *proton.Win) {
		draw(win, u)
	})
	app.Run()
}

func draw(win *proton.Win, u *ui) {
	proton.Split(win, 0.45,
		func(win *proton.Win) {
			proton.H5(win, "Themes")
			proton.Gap(win, 8)
			proton.Caption(win, "Click any theme to apply it live.")
			proton.Gap(win, 8)
			proton.ThemePicker(win, &u.picker, app)
		},
		func(win *proton.Win) {
			proton.PadH(win, 16, func(win *proton.Win) {
				proton.H5(win, "Preview")
				proton.Gap(win, 12)
				proton.Input(win, &u.input, "Type something...")
				proton.Gap(win, 8)
				proton.Pad(win, 4, func(win *proton.Win) {
					proton.Button(win, &u.btn, "Primary Button")
				})
				proton.Gap(win, 8)
				proton.Checkbox(win, &u.check, "A checkbox")
				proton.Gap(win, 8)
				proton.Slider(win, &u.vol)
				proton.Gap(win, 12)
				proton.Divider(win)
				proton.Gap(win, 12)
				proton.Alert(win, proton.AlertInfo, "Info alert")
				proton.Gap(win, 6)
				proton.Alert(win, proton.AlertSuccess, "Success alert")
				proton.Gap(win, 6)
				proton.Alert(win, proton.AlertWarning, "Warning alert")
				proton.Gap(win, 6)
				proton.Alert(win, proton.AlertError, "Error alert")
			})
		},
	)
}
