# Proton

A GUI library for Go. Built on [Gio](https://gioui.org), no C deps, pure Go.

```go
package main

import "github.com/CzaxStudio/proton"

func main() {
    a := proton.New("my app")
    a.Window("Hello", 400, 300, func(win *proton.Win) {
        proton.H3(win, "Hello from Proton!")
    })
    a.Run()
}
```

## Install

```
go get github.com/CzaxStudio/proton
```

Linux needs Gio's system dependencies:
```
apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

macOS and Windows need nothing extra.

## How it works

Gio is immediate mode — your draw function runs every frame. State (button
clicks, text input, checkboxes) lives in your own structs using Proton's
re-exported types. Proton handles the event loop, window setup, and passes
you a `*Win` to draw into.

```go
type State struct {
    name   proton.Editor
    submit proton.Clickable
}

s := &State{}
a.Window("Form", 400, 300, func(win *proton.Win) {
    proton.Input(win, &s.name, "Your name")
    if proton.Button(win, &s.submit, "Go") {
        fmt.Println("Hello,", s.name.Text())
    }
})
```

## Layouts

Layouts take a `*Win` plus a list of `layout.Widget` functions.
Inside those functions, call `proton.Sub(win, gtx)` to get a scoped `*Win`.

```go
proton.Column(win,
    func(gtx layout.Context) layout.Dimensions {
        return proton.Label(proton.Sub(win, gtx), "top")
    },
    func(gtx layout.Context) layout.Dimensions {
        return proton.Label(proton.Sub(win, gtx), "bottom")
    },
)
```

For layouts that need one child to stretch:

```go
proton.Flex(win, layout.Horizontal, layout.SpaceBetween,
    proton.Rigid(func(gtx layout.Context) layout.Dimensions {
        return proton.Label(proton.Sub(win, gtx), "left")
    }),
    proton.Expand(func(gtx layout.Context) layout.Dimensions {
        return proton.Label(proton.Sub(win, gtx), "stretches")
    }),
)
```

### Available layouts

| Function | What it does |
|---|---|
| `Column(win, ...widgets)` | vertical stack |
| `Row(win, ...widgets)` | horizontal row |
| `RowSpread(win, ...widgets)` | horizontal, space between |
| `RowEnd(win, ...widgets)` | horizontal, pushed right |
| `Flex(win, axis, spacing, ...children)` | full control |
| `Rigid(w)` | child takes only what it needs |
| `Expand(w)` | child takes remaining space |
| `ExpandN(w, weight)` | weighted expansion |
| `Split(win, fraction, left, right)` | vertical split pane |
| `HSplit(win, fraction, top, bottom)` | horizontal split pane |
| `Pad(win, dp, w)` | uniform padding |
| `PadH / PadV` | horizontal or vertical padding |
| `PadSides(win, t, r, b, l, w)` | per-side padding |
| `Center(win, w)` | centered in available space |
| `Spacer(dp)` | blank gap (use inside Flex/Row/Column) |

## Widgets

| Function | Returns | Notes |
|---|---|---|
| `Label(win, text)` | dims | body1 style |
| `H1`–`H6(win, text)` | dims | heading sizes |
| `Body2, Caption` | dims | smaller text |
| `Text(win, s, size, color, bold)` | dims | custom text |
| `Button(win, &state, label)` | **bool** | true if clicked this frame |
| `StyledButton(win, &state, label, bg, fg)` | **bool** | custom-colored button |
| `OutlineButton(win, &state, label)` | **bool** | secondary style |
| `IconButton(win, &state, icon, desc)` | **bool** | icon-only button |
| `Clickable(win, &state, fn)` | **bool** | custom content as button |
| `Input(win, &state, hint)` | dims | single-line text field |
| `TextArea(win, &state, hint)` | dims | multi-line text field |
| `Checkbox(win, &state, label)` | **bool** | true if state changed |
| `RadioButton(win, &group, key, label)` | **bool** | true if selection changed |
| `Slider(win, &state)` | **float32** | returns current value 0–1 |
| `ProgressBar(win, progress)` | dims | 0–1 |
| `List(win, &scroll, n, fn)` | dims | virtual scrolling list |
| `HList(win, &scroll, n, fn)` | dims | horizontal list |
| `Divider(win)` | dims | horizontal rule |
| `Rect(win, color, w, h)` | dims | filled rectangle |
| `RoundRect(win, color, w, h, r)` | dims | rounded rectangle |
| `Card(win, bg, corner, pad, fn)` | dims | content in a rounded card |
| `Border(win, color, width, corner, fn)` | dims | content with a border stroke |
| `Badge(win, bg, fg, text)` | dims | small colored label chip |

## Theming

```go
a := proton.New("my app")
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

## Examples

```
go run ./examples/hello
go run ./examples/todo
go run ./examples/calculator
```

## License

MIT
