package main

import (
	"fmt"

	"github.com/CzaxStudio/proton"
)

type UI struct {
	name proton.Editor
	btn  proton.Clickable
	mbox proton.MessageBoxState
}

func main() {
	u := &UI{}
	app := proton.New("Proton Demo")

	app.Window("MessageBox Example", 400, 300, func(win *proton.Win) {
		proton.H5(win, "Greetings")
		proton.Gap(win, 10)

		proton.Input(win, &u.name, "Enter your name")
		proton.Gap(win, 10)

		if proton.Button(win, &u.btn, "Greet Me") {

			u.mbox.Show("Welcome", fmt.Sprintf("Hello %s!", u.name.Text()))
		}

		// This must be called last in the window function to render as a modal overlay
		if win.MessageBox(&u.mbox) == proton.MsgOk {

			fmt.Println("User clicked OK")
		}
	})

	app.Run()
}
