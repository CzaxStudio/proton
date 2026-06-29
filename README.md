# Proton

A GUI library for Go that doesn't make you want to switch to web dev.

### Currently under development for v0.2.5
### Proton win has changed to Context
[![Go Report Card](https://goreportcard.com/badge/github.com/CzaxStudio/proton)](https://goreportcard.com/report/github.com/CzaxStudio/proton) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

## Our First Stargazers

Thank you to the early adopters who supported Proton from the very beginning!

<table>
  <tr>
    <td align="center"><a href="https://github.com/VioGrafu"><img src="https://github.com/VioGrafu.png?size=100" width="100px;" alt=""/><br /><sub><b>@VioGrafu(First Stargazer)</b></sub></a></td>
    <td align="center"><a href="https://github.com/bigwhite"><img src="https://github.com/bigwhite.png?size=100" width="100px;" alt=""/><br /><sub><b>@bigwhite</b></sub></a></td>
    <td align="center"><a href="https://github.com/TanmayCzax"><img src="https://github.com/TanmayCzax.png?size=100" width="100px;" alt=""/><br /><sub><b>@TanmayCzax</b></sub></a></td>
    <td align="center"><a href="https://github.com/aurax"><img src="https://github.com/aurax.png?size=100" width="100px;" alt=""/><br /><sub><b>@aurax</b></sub></a></td>
    <td align="center"><a href="https://github.com/DemonK1"><img src="https://github.com/DemonK1.png?size=100" width="100px;" alt=""/><br /><sub><b>@DemonK1</b></sub></a></td>
    <td align="center"><a href="https://github.com/pekim"><img src="https://github.com/pekim.png?size=100" width="100px;" alt=""/><br /><sub><b>@pekim</b></sub></a>
</td>
</td>
    <td align="center"><a href="https://github.com/fbaube"><img src="https://github.com/fbaube.png?size=100" width="100px;" alt=""/><br /><sub><b>@fbaube</b></sub></a>
</td>
 </td>
    <td align="center"><a href="https://github.com/gorilacrocodille"><img src="https://github.com/gorilacrocodille.png?size=100" width="100px;" alt=""/><br /><sub><b>@gorilacrocodille</b></sub></a></td>   
</td>
    <td align="center"><a href="https://github.com/alanmsant2"><img src="https://github.com/alanmsant2.png?size=100" width="100px;" alt=""/><br /><sub><b>@alanmsant2</b></sub></a></td>
    
</tr>
</table>

## Documentation 

#### https://github.com/CzaxStudio/proton-documentation

## Example apps (made using Proton)

### Note: I have created basic apps, you can create even better apps with Proton.


<img width="813" height="508" alt="GUI demo" src="https://github.com/user-attachments/assets/c8e48374-7e98-41c5-9d46-4427a007b02b" />

<img width="813" height="508" alt="Demo2" src="https://github.com/user-attachments/assets/af7552c4-107f-4760-835a-1ae736f49358" />



## Logo

<img width="1254" height="1254" alt="Proton" src="https://github.com/user-attachments/assets/e044d0b8-a96f-4bc2-9df9-725a41a99ed2" />

# Getting started

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

---

## How it works

Your draw function runs every frame. Call widget functions in order — they stack vertically by default. State lives in your own struct using Proton's re-exported types, so you only ever need one import.

```go
type UI struct {
    btn     proton.Clickable    // button
    name    proton.Editor       // text input
    checked proton.Bool         // checkbox / toggle
    choice  proton.Enum         // radio group
    vol     proton.Float        // slider
    scroll  proton.Scrollable   // list / scroll area
}
```

---

## Widgets

### Text
`Label` `H1`–`H6` `Body2` `Caption` `Text` `Muted` `ColoredText` `ErrorText` `SuccessText` `WarningText`

### Buttons
`Button` `OutlineButton` `IconButton` `Tappable` `Link` `LinkSmall`

### Inputs
`Input` `TextArea` `Checkbox` `Toggle` `RadioButton` `Slider` `ProgressBar` `NumberInput` `SelectBox`

### Lists
`List` `HList` `Scroll` `TextView` `LogView`

### Layout
`Row` `Column` `RowSpread` `RowEnd` `GrowRow` `GrowColumn` `GrowItem` `FixedItem` `FlexSpacer`
`Split` `HSplit` `ResizeSplit` `ResizeHSplit` `Center` `ZStack`
`Pad` `PadH` `PadV` `PadSides` `Gap` `Grid` `MinSize` `MaxWidth`

### Visual
`Divider` `LabeledDivider` `Rect` `RoundRect` `Card` `HoverCard` `Badge` `StatusDot`
`Image` `CodeBlock` `ShortcutHint` `ColorSwatch`

### Feedback
`Toast` `Alert` `AlertDismissable` `Tooltip` `Spinner`

### Overlays & Dialogs
`Overlay` `Tabs` `Accordion` `ContextMenu`

### Utilities
`If` `OnKey` `FocusArea`

---

## Layout

Widgets stack vertically by default. Use `Row` or `Column` to group them differently.

```go
// side by side
proton.Row(win,
    func(win *proton.Win) { proton.Label(win, "left") },
    func(win *proton.Win) { proton.Label(win, "right") },
)

// one child fills remaining space
proton.GrowRow(win,
    proton.FixedItem(win, func(win *proton.Win) { proton.Label(win, "Search:") }),
    proton.GrowItem(win, func(win *proton.Win) { proton.Input(win, &e, "") }),
    proton.FixedItem(win, func(win *proton.Win) { proton.Button(win, &b, "Go") }),
)

// split pane (draggable)
proton.ResizeSplit(win, &u.split, 0.35, leftFn, rightFn)

// padding
proton.Pad(win, 16, func(win *proton.Win) { ... })
proton.PadSides(win, 8, 16, 8, 16, func(win *proton.Win) { ... })

// blank gap
proton.Gap(win, 12)
```

---

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

---

## Alerts and Feedback

```go
proton.Alert(win, proton.AlertInfo,    "Informational message.")
proton.Alert(win, proton.AlertSuccess, "Operation completed.")
proton.Alert(win, proton.AlertWarning, "Proceed with caution.")
proton.Alert(win, proton.AlertError,   "Something went wrong.")

// dismissable
if proton.AlertDismissable(win, &u.closeBtn, proton.AlertInfo, "Click × to close") {
    u.showAlert = false
}

// toast — call last in your draw function
u.toast.Show("Saved!", 2*time.Second)
proton.Toast(win, &u.toast)
```

---

## Async updates

```go
go func() {
    result := fetchFromAPI()
    u.data = result
    win.Invalidate() // ask for a redraw
}()
```

---

## Keyboard shortcuts

```go
proton.OnKey(win, key.ModCtrl, "S", func() { save() })
proton.OnKey(win, 0, key.NameEscape, func() { closeDialog() })
```

---

## Examples

```bash
go run ./examples/hello        # 7 lines, one window
go run ./examples/todo         # classic todo list
go run ./examples/calculator   # buttons and state
go run ./examples/showcase     # layout and theming demo
go run ./examples/kitchen      # every widget in one place
```

---

## Docs

See the [`docs`]((https://github.com/CzaxStudio/proton-documentation)) repo for detailed per-topic guides:

- [Getting started]
- [Text]
- [Buttons]
- [Inputs]
- [Layout]
- [Lists]
- [Visuals]
- [Theming]
- [Advanced]
- [Examples]

---

## License

MIT
