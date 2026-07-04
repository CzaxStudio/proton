# Proton

[![Go Report Card](https://goreportcard.com/badge/github.com/CzaxStudio/proton)](https://goreportcard.com/report/github.com/CzaxStudio/proton) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

##### (Currently under development for v0.2.5)
A GUI library for Go that doesn't make you want to switch to web dev. 

Built on [Gio](https://gioui.org). No C dependencies. Pure Go. Works on Linux, macOS, and Windows.

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
    a.ApplyPalette(proton.NordPalette)
    a.Window("Hello", 480, 300, func(win proton.Context) {
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
## Our First Stargazers 
Thank you to the early adopters who supported Proton from the very beginning! 
<table>  <tr>    
    <td align="center"><a href="https://github.com/VioGrafu"><img src="https://github.com/VioGrafu.png?size=100" width="100px;" alt=""/><br /><sub><b>@VioGrafu</b><br />(First Stargazer)</sub></a></td>    <td align="center"><a href="https://github.com/bigwhite"><img src="https://github.com/bigwhite.png?size=100" width="100px;" alt=""/><br /><sub><b>@bigwhite</b></sub></a></td>    <td align="center"><a href="https://github.com/TanmayCzax"><img src="https://github.com/TanmayCzax.png?size=100" width="100px;" alt=""/><br /><sub><b>@TanmayCzax</b></sub></a></td>    <td align="center"><a href="https://github.com/aurax"><img src="https://github.com/aurax.png?size=100" width="100px;" alt=""/><br /><sub><b>@aurax</b></sub></a></td>    <td align="center"><a href="https://github.com/DemonK1"><img src="https://github.com/DemonK1.png?size=100" width="100px;" alt=""/><br /><sub><b>@DemonK1</b></sub></a></td>    <td align="center"><a href="https://github.com/pekim"><img src="https://github.com/pekim.png?size=100" width="100px;" alt=""/><br /><sub><b>@pekim</b></sub></a></td>    <td align="center"><a href="https://github.com/fbaube"><img src="https://github.com/fbaube.png?size=100" width="100px;" alt=""/><br /><sub><b>@fbaube</b></sub></a></td>    <td align="center"><a href="https://github.com/gorilacrocodille"><img src="https://github.com/gorilacrocodille.png?size=100" width="100px;" alt=""/><br /><sub><b>@gorilacrocodille</b></sub></a></td>    <td align="center"><a href="https://github.com/alanmsant2"><img src="https://github.com/alanmsant2.png?size=100" width="100px;" alt=""/><br /><sub><b>@alanmsant2</b></sub></a></td>  </tr></table>
---

## Install

```bash
go get github.com/CzaxStudio/proton
```

**Linux** — three extra packages required:
```bash
sudo apt install libwayland-dev libxkbcommon-dev libvulkan-dev
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
    func(win proton.Context) { proton.Label(win, "left") },
    func(win proton.Context) { proton.Label(win, "right") },
)

// one child fills remaining space
proton.GrowRow(win,
    proton.FixedItem(win, func(win proton.Context) { proton.Label(win, "Search:") }),
    proton.GrowItem(win, func(win proton.Context) { proton.Input(win, &e, "") }),
    proton.FixedItem(win, func(win proton.Context) { proton.Button(win, &b, "Go") }),
)

// split pane (draggable)
proton.ResizeSplit(win, &u.split, 0.35, leftFn, rightFn)

// padding
proton.Pad(win, 16, func(win proton.Context) { ... })
proton.PadSides(win, 8, 16, 8, 16, func(win proton.Context) { ... })

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
proton.OnKey(win, proton.ModCtrl, "S", func() { save() })
proton.OnKey(win, proton.ModNone, proton.KeyEscape, func() { closeDialog() })
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

See the Docs for Detailed topics

Getting started
- Text
- Buttons
- Inputs
- Layout
- Lists
- Visuals
- Theming
- Advanced
- Examples
- Android

---

## License

MIT

---

## Logo

Load a logo once at startup and draw it anywhere:

```go
//go:embed assets/logo.png
var logoBytes []byte

func main() {
    a := proton.New("myapp")
    a.SetLogoBytes(logoBytes)

    a.Window("My App", 480, 300, func(ctx proton.Context) {
        proton.Logo(ctx, 48, 48)
        proton.Gap(ctx, 8)
        proton.H4(ctx, "My App")
    })
    a.Run()
}
```

Or load from a file path:

```go
a.SetLogo("assets/logo.png")
```

Both PNG and JPEG work. The image is decoded once and cached — never re-read per frame.

---

## Android

Proton apps run on Android through Gio's native support. Same code, no rewrites.

Install the build tool:

```bash
go install gioui.org/cmd/gogio@latest
```

Build an APK:

```bash
gogio -target android -appid com.yourname.yourapp .
adb install yourapp.apk
```

Full setup guide: 10-android.md

---
