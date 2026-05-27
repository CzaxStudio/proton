# Proton

A GUI library for Go. Built on [Gio](https://gioui.org). No C deps, pure Go.

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
```

## Install

```
go get github.com/CzaxStudio/proton
```

Linux system deps:
```
apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

macOS and Windows need nothing extra.

## How it works

Gio is immediate mode — your draw function runs every frame. Widgets called
directly in your draw function stack vertically by default. Use `Row()` or
`Column()` for other arrangements.

State (button clicks, text, checkboxes) lives in your own structs using
Proton's re-exported types so you only need one import.

## Layouts

| Function | What it does |
|---|---|
| `Column(win, ...fns)` | vertical stack |
| `Row(win, ...fns)` | horizontal row |
| `RowSpread(win, ...fns)` | horizontal, space between items |
| `RowEnd(win, ...fns)` | horizontal, pushed to right |
| `GrowRow(win, ...children)` | horizontal with stretch control |
| `GrowColumn(win, ...children)` | vertical with stretch control |
| `GrowItem(win, fn)` | stretchy child for GrowRow/GrowColumn |
| `FixedItem(win, fn)` | fixed child for GrowRow/GrowColumn |
| `Split(win, fraction, left, right)` | side-by-side split pane |
| `HSplit(win, fraction, top, bottom)` | top-bottom split pane |
| `Center(win, fn)` | centered in available space |
| `Pad(win, dp, fn)` | uniform padding |
| `PadH(win, dp, fn)` | left+right padding |
| `PadV(win, dp, fn)` | top+bottom padding |
| `PadSides(win, t, r, b, l, fn)` | per-edge padding |
| `Gap(win, dp)` | blank space between widgets |
| `Grid(win, cols, gap, ...fns)` | fixed-column grid |

## Widgets

| Function | Returns | Notes |
|---|---|---|
| `Label(win, text)` | — | body text |
| `H1`–`H6(win, text)` | — | headings |
| `Body2(win, text)` | — | smaller body |
| `Caption(win, text)` | — | small text |
| `Text(win, s, size, color, bold)` | — | custom text |
| `Button(win, &state, label)` | bool | true if clicked |
| `OutlineButton(win, &state, label)` | bool | ghost style |
| `IconButton(win, &state, icon, desc)` | bool | icon only |
| `Tappable(win, &state, fn)` | bool | custom clickable area |
| `Input(win, &state, hint)` | — | single-line text field |
| `TextArea(win, &state, hint)` | — | multi-line text field |
| `Checkbox(win, &state, label)` | bool | true if changed |
| `Toggle(win, &state, label)` | bool | switch, true if changed |
| `RadioButton(win, &group, key, label)` | bool | true if changed |
| `Slider(win, &state)` | float32 | current value 0–1 |
| `ProgressBar(win, progress)` | — | 0–1 |
| `List(win, &scroll, n, fn)` | — | virtual scrolling list |
| `HList(win, &scroll, n, fn)` | — | horizontal list |
| `Scroll(win, &scroll, fn)` | — | scrollable content area |
| `Divider(win)` | — | horizontal rule |
| `Rect(win, color, w, h)` | — | filled rectangle |
| `RoundRect(win, color, w, h, r)` | — | rounded rectangle |
| `Card(win, bg, corner, pad, fn)` | — | content in a card |
| `Badge(win, bg, fg, text)` | — | small colored chip |
| `Image(win, img, w, h)` | — | draw an image |
| `MinSize(win, w, h, fn)` | — | minimum size constraint |
| `MaxWidth(win, w, fn)` | — | maximum width constraint |
| `Tooltip(win, &state, tip, fn)` | — | hover tooltip |
| `Toast(win, &state)` | — | timed notification overlay |

## Keyboard

```go
// fire a function on Ctrl+S
proton.OnKey(win, key.ModCtrl, "S", func() { save() })

// fire on Escape
proton.OnKey(win, 0, key.NameEscape, func() { closeDialog() })
```

## Image loading

```go
// load once at startup
img, err := proton.LoadImage("photo.png")
if err != nil {
    log.Fatal(err)
}

// draw every frame
proton.Image(win, img, 200, 150)
```

## Toast notifications

```go
type UI struct {
    toast proton.ToastState
}

// trigger from anywhere (goroutine-safe)
u.toast.Show("File saved!", 2*time.Second)

// in your draw function (call last so it renders on top)
proton.Toast(win, &u.toast)
```

## Theming

```go
a.ApplyPalette(proton.DarkPalette)
a.ApplyPalette(proton.NordPalette)
a.ApplyPalette(proton.RosePinePalette)
a.ApplyPalette(proton.CatppuccinPalette)

// custom
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x1e1e2e),
    Fg:        proton.RGB(0xcdd6f4),
    Primary:   proton.RGB(0x89b4fa),
    PrimaryFg: proton.RGB(0x1e1e2e),
})

a.SetFontScale(1.1)
```

## State types

Declare these in your UI state struct — no imports beyond `proton` needed:

```go
type UI struct {
    btn     proton.Clickable   // button / tappable
    name    proton.Editor      // input / textarea
    checked proton.Bool        // checkbox / toggle
    choice  proton.Enum        // radio group
    vol     proton.Float       // slider
    scroll  proton.Scrollable  // list / scroll
}
```

## Examples

```
go run ./examples/hello
go run ./examples/todo
go run ./examples/calculator
go run ./examples/showcase
```

## License

MIT
