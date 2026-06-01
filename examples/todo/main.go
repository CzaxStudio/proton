package main

import "github.com/CzaxStudio/proton"

type task struct {
	text string
	done proton.Bool
}

type ui struct {
	input  proton.Editor
	add    proton.Clickable
	tasks  []task
	scroll proton.Scrollable
}

func main() {
	u := &ui{}

	a := proton.New("todo")
	a.ApplyPalette(proton.NordPalette)
	a.Window("Todo", 420, 580, func(win *proton.Win) {
		draw(win, u)
	})
	a.Run()
}

func draw(win *proton.Win, u *ui) {
	proton.Column(win,
		func(win *proton.Win) {
			proton.Row(win,
				func(win *proton.Win) {
					proton.Input(win, &u.input, "what needs doing?")
				},
				func(win *proton.Win) {
					if proton.Button(win, &u.add, "Add") {
						t := u.input.Text()
						if t != "" {
							u.tasks = append(u.tasks, task{text: t})
							u.input.SetText("")
						}
					}
				},
			)
		},
		func(win *proton.Win) {
			proton.Divider(win)
		},
		func(win *proton.Win) {
			proton.List(win, &u.scroll, len(u.tasks), func(win *proton.Win, i int) {
				t := &u.tasks[i]
				proton.Checkbox(win, &t.done, t.text)
			})
		},
	)
}
