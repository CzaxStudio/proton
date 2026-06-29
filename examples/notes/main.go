// notes — a minimal but fully functional note-taking app.
// Demonstrates: ResizeSplit, SearchInput, HoverCard, TextArea, TextView,
// Overlay modal, Toast, OnKey shortcuts, Badge tags, Accordion, Muted.
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzaxStudio/proton"
)

type note struct {
	title string
	body  string
	tags  []string
}

type ui struct {
	split  proton.ResizeSplitState
	search proton.SearchState
	scroll proton.Scrollable

	rowBtns [256]proton.Clickable
	newBtn  proton.Clickable
	delBtn  proton.Clickable
	saveBtn proton.Clickable
	editBtn proton.Clickable

	titleEd proton.Editor
	bodyEd  proton.Editor

	modal      proton.OverlayState
	confirmDel proton.Clickable
	cancelDel  proton.Clickable

	bodyScroll proton.Scrollable
	toast      proton.ToastState

	notes    []note
	selected int
	editing  bool

	acc1    proton.AccordionState
	acc1btn proton.Clickable
}

func newUI() *ui {
	u := &ui{selected: -1}
	u.notes = []note{
		{
			"Welcome to Notes",
			"This is a note-taking app built entirely with Proton.\n\n" +
				"Press Ctrl+N to create a new note.\n" +
				"Press Ctrl+S to save.\n" +
				"Click Edit to modify a note.\n" +
				"Use the search bar to filter notes in real time.",
			[]string{"intro", "proton"},
		},
		{
			"Proton is fast",
			"Built on Gio — immediate mode GUI at ~60fps.\n" +
				"Pure Go. No CGo. No Electron.\n" +
				"Works on Linux, macOS, and Windows.",
			[]string{"go", "proton"},
		},
		{
			"Shopping list",
			"- Milk\n- Eggs\n- Coffee\n- More coffee",
			[]string{"personal"},
		},
		{
			"Meeting notes",
			"Q3 planning:\n" +
				"- Review roadmap\n" +
				"- Assign owners\n" +
				"- Set deadlines\n" +
				"- Actually follow up this time",
			[]string{"work"},
		},
	}
	return u
}

var gApp *proton.App

func main() {
	u := newUI()
	gApp = proton.New("notes")
	gApp.ApplyPalette(proton.NordPalette)
	gApp.Window("Notes", 880, 580, func(ctx proton.Context) {
		draw(ctx, u)
	})
	gApp.Run()
}

func draw(ctx proton.Context, u *ui) {
	proton.OnKey(ctx, proton.ModCtrl, "N", func() { startNew(u) })
	proton.OnKey(ctx, proton.ModCtrl, "S", func() { saveNote(u, ctx) })

	proton.ResizeSplit(ctx, &u.split, 0.30,
		func(ctx proton.Context) { sidebar(ctx, u) },
		func(ctx proton.Context) {
			proton.PadH(ctx, 14, func(ctx proton.Context) {
				editor(ctx, u)
			})
		},
	)

	proton.Overlay(ctx, &u.modal, func(ctx proton.Context) {
		proton.MinSize(ctx, 310, 0, func(ctx proton.Context) {
			proton.Card(ctx, proton.RGB(0x2e3440), 12, 24, func(ctx proton.Context) {
				proton.H5(ctx, "Delete note?")
				proton.Gap(ctx, 8)
				if u.selected >= 0 && u.selected < len(u.notes) {
					proton.Muted(ctx, fmt.Sprintf(`"%s" will be permanently deleted.`, u.notes[u.selected].title))
				}
				proton.Gap(ctx, 20)
				proton.RowEnd(ctx,
					func(ctx proton.Context) {
						proton.Pad(ctx, 4, func(ctx proton.Context) {
							if proton.OutlineButton(ctx, &u.cancelDel, "Cancel") {
								u.modal.Hide()
							}
						})
					},
					func(ctx proton.Context) { proton.Gap(ctx, 8) },
					func(ctx proton.Context) {
						proton.Pad(ctx, 4, func(ctx proton.Context) {
							if proton.Button(ctx, &u.confirmDel, "Delete") {
								doDelete(u)
								u.modal.Hide()
								u.toast.Show("Note deleted.", 2*time.Second)
							}
						})
					},
				)
			})
		})
	})

	proton.Toast(ctx, &u.toast)
}

func sidebar(ctx proton.Context, u *ui) {
	proton.PadSides(ctx, 0, 10, 0, 0, func(ctx proton.Context) {
		proton.RowSpread(ctx,
			func(ctx proton.Context) { proton.H5(ctx, "Notes") },
			func(ctx proton.Context) {
				proton.Pad(ctx, 2, func(ctx proton.Context) {
					if proton.Button(ctx, &u.newBtn, "+ New") {
						startNew(u)
					}
				})
			},
		)
		proton.Gap(ctx, 10)

		q := proton.SearchInput(ctx, &u.search, "Search...")
		proton.Gap(ctx, 8)

		// filter
		type item struct {
			n   note
			idx int
		}
		var items []item
		for i, n := range u.notes {
			ql := strings.ToLower(strings.TrimSpace(q))
			if ql == "" ||
				strings.Contains(strings.ToLower(n.title), ql) ||
				strings.Contains(strings.ToLower(n.body), ql) {
				items = append(items, item{n, i})
			}
		}

		proton.List(ctx, &u.scroll, len(items), func(ctx proton.Context, i int) {
			it := items[i]
			active := u.selected == it.idx
			bg := proton.RGB(0x3b4252)
			if active {
				bg = proton.RGB(0x4c566a)
			}
			hov := proton.RGB(0x434c5e)
			if active {
				hov = bg
			}
			proton.PadV(ctx, 2, func(ctx proton.Context) {
				if proton.HoverCard(ctx, &u.rowBtns[i], bg, hov, 7, func(ctx proton.Context) {
					proton.PadV(ctx, 10, func(ctx proton.Context) {
						proton.PadH(ctx, 10, func(ctx proton.Context) {
							t := it.n.title
							if t == "" {
								t = "Untitled"
							}
							proton.Label(ctx, t)
							proton.Gap(ctx, 3)
							prev := strings.ReplaceAll(it.n.body, "\n", " ")
							if len(prev) > 55 {
								prev = prev[:55] + "…"
							}
							proton.Muted(ctx, prev)
						})
					})
				}) {
					if !u.editing {
						u.selected = it.idx
					}
				}
			})
		})
	})
}

func editor(ctx proton.Context, u *ui) {
	if u.selected < 0 || u.selected >= len(u.notes) {
		proton.Center(ctx, func(ctx proton.Context) {
			proton.Muted(ctx, "Select a note or create one with Ctrl+N")
		})
		return
	}
	n := &u.notes[u.selected]

	// toolbar
	proton.RowSpread(ctx,
		func(ctx proton.Context) {
			if u.editing {
				proton.GrowRow(ctx,
					proton.GrowItem(ctx, func(ctx proton.Context) {
						proton.Input(ctx, &u.titleEd, "Title")
					}),
				)
			} else {
				t := n.title
				if t == "" {
					t = "Untitled"
				}
				proton.H5(ctx, t)
			}
		},
		func(ctx proton.Context) {
			proton.Row(ctx,
				func(ctx proton.Context) {
					proton.Pad(ctx, 4, func(ctx proton.Context) {
						if u.editing {
							if proton.Button(ctx, &u.saveBtn, "Save") {
								saveNote(u, ctx)
							}
						} else {
							if proton.OutlineButton(ctx, &u.editBtn, "Edit") {
								u.editing = true
								u.titleEd.SetText(n.title)
								u.bodyEd.SetText(n.body)
							}
						}
					})
				},
				func(ctx proton.Context) { proton.Gap(ctx, 6) },
				func(ctx proton.Context) {
					proton.Pad(ctx, 4, func(ctx proton.Context) {
						if proton.OutlineButton(ctx, &u.delBtn, "Delete") {
							u.modal.Show()
						}
					})
				},
			)
		},
	)

	if len(n.tags) > 0 {
		proton.Gap(ctx, 6)
		proton.Row(ctx, tagFns(n.tags)...)
	}

	proton.Gap(ctx, 8)
	proton.Divider(ctx)
	proton.Gap(ctx, 10)

	if u.editing {
		proton.GrowColumn(ctx,
			proton.GrowItem(ctx, func(ctx proton.Context) {
				proton.TextArea(ctx, &u.bodyEd, "Write your note here…")
			}),
		)
	} else {
		proton.TextView(ctx, &u.bodyScroll, n.body)
	}
}

func tagFns(tags []string) []func(proton.Context) {
	colors := []struct{ bg, fg proton.Palette }{
		{proton.NordPalette, proton.NordPalette},
		{proton.EverforestDarkPalette, proton.EverforestDarkPalette},
		{proton.CatppuccinPalette, proton.CatppuccinPalette},
		{proton.DraculaPalette, proton.DraculaPalette},
	}
	var fns []func(proton.Context)
	for i, tag := range tags {
		tag := tag
		c := colors[i%len(colors)]
		fns = append(fns, func(ctx proton.Context) {
			proton.Badge(ctx, c.bg.Primary, c.fg.PrimaryFg, tag)
		})
		fns = append(fns, func(ctx proton.Context) { proton.Gap(ctx, 5) })
	}
	return fns
}

func startNew(u *ui) {
	u.notes = append(u.notes, note{})
	u.selected = len(u.notes) - 1
	u.editing = true
	u.titleEd.SetText("")
	u.bodyEd.SetText("")
}

func saveNote(u *ui, ctx proton.Context) {
	if !u.editing || u.selected < 0 || u.selected >= len(u.notes) {
		return
	}
	u.notes[u.selected].title = u.titleEd.Text()
	u.notes[u.selected].body = u.bodyEd.Text()
	u.editing = false
	u.toast.Show("Saved.", 2*time.Second)
	ctx.Invalidate()
}

func doDelete(u *ui) {
	if u.selected < 0 || u.selected >= len(u.notes) {
		return
	}
	u.notes = append(u.notes[:u.selected], u.notes[u.selected+1:]...)
	if u.selected >= len(u.notes) {
		u.selected = len(u.notes) - 1
	}
	u.editing = false
}
