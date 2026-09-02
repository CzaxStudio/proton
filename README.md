# Proton

A GUI library for Go that doesn't make you want to switch to web dev.

[![Go Report Card](https://goreportcard.com/badge/github.com/CzaxStudio/proton)](https://goreportcard.com/report/github.com/CzaxStudio/proton) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

## We Need Contributors
We are open to accepting new contributors! If you would like to join the project, please open a pull request. We will review your GitHub profile and projects before making a decision

## Why Proton?
**Zero CGO Required: Cross-compile to Windows and macOS flawlessly from any machine without fighting external compiler toolchains.**

**Pure Go Ecosystem: Built on top of Gio, maintaining a 100% Go native development experience.**

**API Immunity: No raw Gio types pollute your code. If the underlying rendering engine updates, your app's codebase remains completely untouched.**

## Our First Stargazers

Thank you to the early adopters who supported Proton from the very beginning!

<table>
  <tr>
    <td align="center"><a href="https://github.com/VioGrafu"><img src="https://github.com/VioGrafu.png?size=100" width="100px;" alt=""/><br /><sub><b>@VioGrafu</b><br />(First Stargazer)</sub></a></td>
    <td align="center"><a href="https://github.com/bigwhite"><img src="https://github.com/bigwhite.png?size=100" width="100px;" alt=""/><br /><sub><b>@bigwhite</b></sub></a></td>
    <td align="center"><a href="https://github.com/TanmayCzax"><img src="https://github.com/TanmayCzax.png?size=100" width="100px;" alt=""/><br /><sub><b>@TanmayCzax</b></sub></a></td>
    <td align="center"><a href="https://github.com/aurax"><img src="https://github.com/aurax.png?size=100" width="100px;" alt=""/><br /><sub><b>@aurax</b></sub></a></td>
    <td align="center"><a href="https://github.com/DemonK1"><img src="https://github.com/DemonK1.png?size=100" width="100px;" alt=""/><br /><sub><b>@DemonK1</b></sub></a></td>
    <td align="center"><a href="https://github.com/pekim"><img src="https://github.com/pekim.png?size=100" width="100px;" alt=""/><br /><sub><b>@pekim</b></sub></a></td>
    <td align="center"><a href="https://github.com/fbaube"><img src="https://github.com/fbaube.png?size=100" width="100px;" alt=""/><br /><sub><b>@fbaube</b></sub></a></td>
    <td align="center"><a href="https://github.com/gorilacrocodille"><img src="https://github.com/gorilacrocodille.png?size=100" width="100px;" alt=""/><br /><sub><b>@gorilacrocodille</b></sub></a></td>
    <td align="center"><a href="https://github.com/alanmsant2"><img src="https://github.com/alanmsant2.png?size=100" width="100px;" alt=""/><br /><sub><b>@alanmsant2</b></sub></a></td>
  </tr>
</table>

### Special Thanks

<td align="center"><a href="https://github.com/diamondosas"><img src="https://github.com/diamondosas.png?size=100" width="100px;" alt=""/><br /><sub><b>@diamondosas(Contributor)</b></sub></a></td>

## Documentation

**[Official Docs For Proton](https://nexus-65.gitbook.io/proton)**

**[Docs Repo](https://github.com/CzaxStudio/proton-documentation)**

---

## Example apps

### GIF

**You can create even better! Proton is not just for Todo or Calculator apps**

<img width="813" height="508" alt="GUI demo" src="https://github.com/user-attachments/assets/c8e48374-7e98-41c5-9d46-4427a007b02b" />

<img width="813" height="508" alt="Demo2" src="https://github.com/user-attachments/assets/af7552c4-107f-4760-835a-1ae736f49358" />

### Pictures

<img width="471" height="542" alt="Sample2" src="https://github.com/user-attachments/assets/77bc7ba2-c06b-4c58-805f-574d1952cd57" />

<img width="1366" height="729" alt="Sample1" src="https://github.com/user-attachments/assets/13e6880b-e378-457e-b717-9b3953e6ea06" />

**[Code](https://github.com/CzaxStudio/proton/blob/main/examples/dashboard/main.go)**


## Logo

<img width="1254" height="1254" alt="ChatGPT Image Jul 10, 2026, 01_09_19 PM" src="https://github.com/user-attachments/assets/750a0671-1527-488e-851f-8c1f38f39128" />


---

# Getting Started

```go
package main

import "github.com/CzaxStudio/proton"

func main() {
    a := proton.New("hello")
    a.Window("Hello", 400, 200, func(ctx proton.Context) {
        proton.H3(ctx, "Hello from Proton!")
    })
    a.Run()
}
```

## Install

```
go get github.com/CzaxStudio/proton
```

Then run once to pull Gio's dependencies:

```
go mod tidy
```

**Linux** — three system packages required:
```bash
sudo apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

macOS and Windows need nothing extra.

---

## How it works

Your draw function runs every frame. Call widget functions in order — they stack vertically by default. State lives in your own struct. No `setState`, no component trees, no XML.

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
`Input` `TextArea` `Checkbox` `Toggle` `RadioButton` `Slider` `ProgressBar` `NumberInput` `SelectBox` `SearchInput`

### Lists
`List` `HList` `Scroll` `TextView` `LogView`

### Layout
`Row` `Column` `RowSpread` `RowEnd` `GrowRow` `GrowColumn` `GrowItem` `FixedItem` `FlexSpacer`
`Split` `HSplit` `ResizeSplit` `ResizeHSplit` `Center` `ZStack`
`Pad` `PadH` `PadV` `PadSides` `Gap` `Grid` `MinSize` `MaxWidth`

### Visual
`Divider` `LabeledDivider` `Rect` `RoundRect` `Card` `HoverCard` `Badge` `StatusDot`
`Avatar` `Tag` `Image` `Logo` `CodeBlock` `ShortcutHint` `ColorSwatch`

### Data
`Table` `ProgressRing` `Stepper`

### Feedback
`Toast` `Alert` `AlertDismissable` `Tooltip` `Spinner`

### Overlays & Dialogs
`Overlay` `Tabs` `Accordion` `ContextMenu`

### Utilities
`If` `OnKey` `FocusArea`

---

## Layout

```go
// side by side
proton.Row(ctx,
    func(ctx proton.Context) { proton.Label(ctx, "left") },
    func(ctx proton.Context) { proton.Label(ctx, "right") },
)

// one child fills remaining space
proton.GrowRow(ctx,
    proton.FixedItem(ctx, func(ctx proton.Context) { proton.Label(ctx, "Search:") }),
    proton.GrowItem(ctx, func(ctx proton.Context) { proton.Input(ctx, &e, "") }),
    proton.FixedItem(ctx, func(ctx proton.Context) { proton.Button(ctx, &b, "Go") }),
)

// draggable split pane
proton.ResizeSplit(ctx, &u.split, 0.35, leftFn, rightFn)

// padding
proton.Pad(ctx, 16, func(ctx proton.Context) { ... })
proton.PadSides(ctx, 8, 16, 8, 16, func(ctx proton.Context) { ... })

// blank space
proton.Gap(ctx, 12)
```

---

## Theming

46 built-in palettes. One line to apply any of them.

```go
a.ApplyPalette(proton.NordPalette)
a.ApplyPalette(proton.CatppuccinPalette)
a.ApplyPalette(proton.DraculaPalette)
a.ApplyPalette(proton.TokyoNightPalette)
a.ApplyPalette(proton.GruvboxDarkPalette)
a.ApplyPalette(proton.RosePinePalette)
// ... 40 more
```

### Custom palette

```go
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x1e1e2e),
    Fg:        proton.RGB(0xcdd6f4),
    Primary:   proton.RGB(0x89b4fa),
    PrimaryFg: proton.RGB(0x1e1e2e),
})
```

### Hex color codes

No structs needed — just pass the hex string:

```go
a.ThemeBuilder().
    Bg("#1e1e2e").
    Fg("#cdd6f4").
    Primary("#89b4fa").
    PrimaryFg("#1e1e2e").
    Apply()

// patch one color on the current theme
a.ColorCode("primary", "#ff6b6b")
a.ColorCode("bg", "#0d1117")
```

Accepted formats: `"#rrggbb"`, `"rrggbb"`, `"#rgb"`, `"#rrggbbaa"`.

### Background colors

```go
a.SetBackgroundCode("#1a1b26")
a.SetBackgroundRGB(26, 27, 38)
a.SetBackgroundGradient("#1a1b26", "#2d1b69", "vertical")
a.SetBackgroundRainbow() // animated full-spectrum gradient
```

### Font scale

```go
a.SetFontScale(1.1)
```

### Live theme picker

Drop into a settings window to let users switch themes at runtime:

```go
type UI struct {
    picker proton.ThemePickerState
}

proton.ThemePicker(ctx, &u.picker, a)
```

---

## Logo

Load once at startup, draw anywhere.

```go
//go:embed assets/logo.png
var logoBytes []byte

func main() {
    a := proton.New("myapp")
    a.SetLogoBytes(logoBytes)

    a.Window("My App", 480, 300, func(ctx proton.Context) {
        proton.Row(ctx,
            func(ctx proton.Context) { proton.Logo(ctx, 40, 40) },
            func(ctx proton.Context) { proton.Gap(ctx, 10) },
            func(ctx proton.Context) { proton.H5(ctx, "My App") },
        )
    })
    a.Run()
}
```

Or load from a file path:

```go
a.SetLogo("assets/logo.png")
```

PNG and JPEG both work. The image is decoded once and cached — not re-read per frame.

---

## Alerts and Feedback

```go
proton.Alert(ctx, proton.AlertInfo,    "Informational message.")
proton.Alert(ctx, proton.AlertSuccess, "Operation completed.")
proton.Alert(ctx, proton.AlertWarning, "Proceed with caution.")
proton.Alert(ctx, proton.AlertError,   "Something went wrong.")

// dismissable
if proton.AlertDismissable(ctx, &u.closeBtn, proton.AlertInfo, "Click x to close") {
    u.showAlert = false
}

// toast — call last in your draw function
u.toast.Show("Saved!", 2*time.Second)
proton.Toast(ctx, &u.toast)
```

---

## New in v0.2.x

**Table**
```go
proton.Table(ctx,
    []string{"Name", "Status", "Score"},
    []proton.TableRow{
        {"Alice", "Active", "98"},
        {"Bob",   "Away",   "74"},
    },
)
```

**Stepper**
```go
proton.Stepper(ctx, currentStep, []string{"Build", "Test", "Stage", "Deploy"})
```

**ProgressRing**
```go
proton.ProgressRing(ctx, 0.72, 48, 5, proton.RGB(0x88c0d0))
```

**SearchInput**
```go
q := proton.SearchInput(ctx, &u.search, "Search notes...")
```

**Avatar**
```go
proton.Avatar(ctx, "AJ", proton.RGB(0x5e81ac), proton.RGB(0xeceff4), 40)
```

**NumberInput**
```go
qty := proton.NumberInput(ctx, &u.qty, 1, 99, 1)
```

**Overlay / modal**
```go
proton.Overlay(ctx, &u.modal, func(ctx proton.Context) {
    proton.Card(ctx, proton.RGB(0x2e3440), 12, 24, func(ctx proton.Context) {
        proton.H5(ctx, "Confirm?")
        proton.Gap(ctx, 16)
        proton.Pad(ctx, 4, func(ctx proton.Context) {
            if proton.Button(ctx, &u.closeBtn, "Close") {
                u.modal.Hide()
            }
        })
    })
})
```

---

## Async Updates

```go
go func() {
    result := fetchFromAPI()
    u.data = result
    ctx.Invalidate()
}()
```

---

## Keyboard Shortcuts

```go
proton.OnKey(ctx, proton.ModCtrl, "S", func() { save() })
proton.OnKey(ctx, proton.ModNone, proton.KeyEscape, func() { closeDialog() })
proton.OnKey(ctx, proton.ModCtrl|proton.ModShift, "N", func() { newWindow() })
```

---

## Window Options

```go
a.WindowEx("App", 800, 600, []proton.WindowOption{
    proton.Fullscreen(),
}, draw)
```

---

## Android

Same code runs on Android. No rewrites.

```bash
go install gioui.org/cmd/gogio@latest
gogio -target android -appid com.yourname.yourapp .
adb install yourapp.apk
```

Full guide: [docs/10-android.md](https://github.com/CzaxStudio/proton-documentation/blob/main/docs/10-android.md)

---

## API Immunity

Every draw function takes `proton.Context` — an interface. No Gio types appear in the public API. If Gio's internals change in a future version, only Proton's implementation updates. Your code keeps compiling unchanged.

```go
func (app *MyApp) draw(ctx proton.Context) {
    proton.Button(ctx, &app.btn, "Click")
}
```

---

## Examples

```bash
go run ./examples/hello        # minimal — one window, one label
go run ./examples/todo         # todo list
go run ./examples/calculator   # grid of buttons
go run ./examples/notes        # note-taking app with search and split pane
go run ./examples/dashboard    # dev dashboard with charts, logs, and tables
go run ./examples/showcase     # every widget in one place
go run ./examples/themes       # live theme picker
go run ./examples/logoapp      # custom logo with go:embed
go run ./examples/kitchen      # stress test for all features
```

---

## License

MIT

**Currently under development for v0.4.0**
