package main

import "github.com/CzaxStudio/proton"

type UI struct {
	name proton.Editor
	btn  proton.Clickable
}

func main() {
	u := &UI{}
	a := proton.New("my app")
	a.Window("Hello", 480, 300, func(win *proton.Win) {
		proton.H3(win, "Hello from Proton!")
		proton.Gap(win, 8)
		proton.Input(win, &u.name, "Your name")
		proton.Gap(win, 8)
		if proton.Button(win, &u.btn, "Go") {
			println("Hello,", u.name.Text())
		}
	})
	a.Run()
}
