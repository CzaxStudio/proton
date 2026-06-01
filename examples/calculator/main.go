package main

import (
	"fmt"
	"strconv"

	"github.com/CzaxStudio/proton"
)

type calc struct {
	display string
	prev    float64
	op      string
	fresh   bool
	btns    [16]proton.Clickable
}

var keys = [16]string{
	"7", "8", "9", "/",
	"4", "5", "6", "*",
	"1", "2", "3", "-",
	"C", "0", "=", "+",
}

func main() {
	c := &calc{display: "0"}

	a := proton.New("calc")
	a.ApplyPalette(proton.DarkPalette)
	a.SetLogo(Proton_data)

	a.Window("Calculator", 300, 450, func(win *proton.Win) {
		proton.Center(win, func(win *proton.Win) {
			proton.Logo(win, 64, 64)
		})
		proton.Gap(win, 10)

		drawCalc(win, c)
	})
	a.Run()
}

func drawCalc(win *proton.Win, c *calc) {
	proton.Column(win,
		func(win *proton.Win) {
			proton.PadSides(win, 8, 8, 16, 8, func(win *proton.Win) {
				proton.H4(win, c.display)
			})
		},
		func(win *proton.Win) { drawGrid(win, c) },
	)
}

func drawGrid(win *proton.Win, c *calc) {
	proton.Column(win,
		func(win *proton.Win) { drawRow(win, c, 0) },
		func(win *proton.Win) { drawRow(win, c, 1) },
		func(win *proton.Win) { drawRow(win, c, 2) },
		func(win *proton.Win) { drawRow(win, c, 3) },
	)
}

func drawRow(win *proton.Win, c *calc, row int) {
	proton.GrowRow(win,
		proton.GrowItem(win, func(win *proton.Win) { drawKey(win, c, row*4+0) }),
		proton.GrowItem(win, func(win *proton.Win) { drawKey(win, c, row*4+1) }),
		proton.GrowItem(win, func(win *proton.Win) { drawKey(win, c, row*4+2) }),
		proton.GrowItem(win, func(win *proton.Win) { drawKey(win, c, row*4+3) }),
	)
}

func drawKey(win *proton.Win, c *calc, idx int) {
	proton.Pad(win, 3, func(win *proton.Win) {
		if proton.Button(win, &c.btns[idx], keys[idx]) {
			press(c, keys[idx])
		}
	})
}

func press(c *calc, key string) {
	switch key {
	case "C":
		c.display = "0"
		c.prev = 0
		c.op = ""
		c.fresh = false
	case "+", "-", "*", "/":
		v, _ := strconv.ParseFloat(c.display, 64)
		c.prev = v
		c.op = key
		c.fresh = true
	case "=":
		if c.op == "" {
			return
		}
		v, _ := strconv.ParseFloat(c.display, 64)
		c.display = cleanNum(operate(c.prev, v, c.op))
		c.op = ""
		c.fresh = false
	default:
		if c.display == "0" || c.fresh {
			c.display = key
			c.fresh = false
		} else {
			c.display += key
		}
	}
}

func operate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	}
	return b
}

func cleanNum(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return strconv.FormatFloat(v, 'f', -1, 64)
}
