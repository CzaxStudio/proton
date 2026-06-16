package main

import "github.com/CzaxStudio/proton"

func main() {
	a := proton.New("hello")
	a.Window("Hello", 400, 200, func(win *proton.Win) {
		proton.H3(win, "Hello from Proton!")
	})
	a.Run()
}
