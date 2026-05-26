# Proton

An ultra easy GUI library for Go. Built on [Gio](https://gioui.org). No C deps, pure Go.

## To be published on June 12th 2026.

```go
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
		proton.Input(win, &u.name, "Your name")
		proton.RowEnd(win, func(win *proton.Win) {
			if proton.Button(win, &u.btn, "Go") {
				println("Hello,", u.name.Text())
			}
		})
	})
	a.Run()
}
```

## Install

```
go get github.com/CzaxStudio/proton
```

Linux needs Gio's system packages:
```
apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

macOS and Windows need nothing extra.

## How it works

Gio is immediate mode — your draw function runs every frame and you just call
widget functions in order. State lives in your own structs; Proton re-exports
the state types so you only need one import.

The `*Win` passed to your draw function carries the current layout context.
When you use layout helpers like `Column` or `Pad`, they scope it for each
child automatically.

```go
type State struct {
    count int
    btn   proton.Clickable
}

s := &State{}
a.Window("Counter", 300, 200, func(win *proton.Win) {
    proton.H4(win, fmt.Sprintf("Count: %d", s.count))
    proton.Gap(win, 8)
    if proton.Button(win, &s.btn, "Increment") {
        s.count++
    }
})
```

## Layouts

```go
// vertical stack
proton.Column(win,
    func(win *proton.Win) { proton.Label(win, "first") },
    func(win *proton.Win) { proton.Label(win, "second") },
)

// horizontal row
proton.Row(win,
    func(win *proton.Win) { proton.Label(win, "left") },
    func(win *proton.Win) { proton.Label(win, "right") },
)

// one child stretches, others are fixed
proton.GrowRow(win,
    proton.FixedItem(win, func(win *proton.Win) { proton.Label(win, "label") }),
    proton.GrowItem(win, func(win *proton.Win) { proton.Input(win, &e, "") }),
    proton.FixedItem(win, func(win *proton.Win) { proton.Button(win, &b, "Go") }),
)

// padding
proton.Pad(win, 16, func(win *proton.Win) { ... })
proton.PadV(win, 8, func(win *proton.Win) { ... })
proton.PadH(win, 8, func(win *proton.Win) { ... })
proton.PadSides(win, top, right, bottom, left, func(win *proton.Win) { ... })

// centering
proton.Center(win, func(win *proton.Win) { ... })

// blank gap inside a row or column
proton.Gap(win, 12)

// split panes
proton.Split(win, 0.3, leftFn, rightFn)
proton.HSplit(win, 0.4, topFn, bottomFn)

// grid
proton.Grid(win, 3, 8,
    func(win *proton.Win) { proton.Label(win, "one") },
    func(win *proton.Win) { proton.Label(win, "two") },
    func(win *proton.Win) { proton.Label(win, "three") },
    func(win *proton.Win) { proton.Label(win, "four") },
)
```

## Widgets

| Function | What it does |
|---|---|
| `Label(win, text)` | body text |
| `H1` – `H6(win, text)` | headings |
| `Body2(win, text)` | smaller body text |
| `Caption(win, text)` | small caption text |
| `Text(win, s, size, color, bold)` | custom text |
| `Button(win, &state, label) bool` | filled button |
| `OutlineButton(win, &state, label) bool` | ghost button |
| `IconButton(win, &state, icon, desc) bool` | icon-only button |
| `Tappable(win, &state, fn) bool` | any content as a button |
| `Input(win, &state, hint)` | single-line text field |
| `TextArea(win, &state, hint)` | multi-line text field |
| `Checkbox(win, &state, label) bool` | checkbox |
| `RadioButton(win, &group, key, label) bool` | radio button |
| `Slider(win, &state) float32` | slider, returns 0–1 |
| `ProgressBar(win, progress)` | progress bar, 0–1 |
| `List(win, &scroll, n, fn)` | vertical scrolling list |
| `HList(win, &scroll, n, fn)` | horizontal scrolling list |
| `Divider(win)` | horizontal rule |
| `Rect(win, color, w, h)` | filled rectangle |
| `RoundRect(win, color, w, h, r)` | rounded rectangle |
| `Card(win, bg, corner, pad, fn)` | content in a card background |
| `Badge(win, bg, fg, text)` | small colored chip |
| `Toast(win, &state)` | overlay notification at bottom |

## Theming

```go
a := proton.New("app")
a.ApplyPalette(proton.DarkPalette)
a.ApplyPalette(proton.NordPalette)
a.ApplyPalette(proton.RosePinePalette)
a.ApplyPalette(proton.CatppuccinPalette)

// or your own
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x1e1e2e),
    Fg:        proton.RGB(0xcdd6f4),
    Primary:   proton.RGB(0x89b4fa),
    PrimaryFg: proton.RGB(0x1e1e2e),
})

a.SetFontScale(1.1)
```

## Examples

All examples live in the same module, so just:

```
cd examples/hello && go run .
cd examples/todo && go run .
cd examples/calculator && go run .
```

## License

MIT
