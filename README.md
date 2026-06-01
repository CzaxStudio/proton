# Proton

A GUI library for Go. Built on [Gio](https://gioui.org). No C deps, pure Go.

# Example apps (made using Proton)
*Note --> These are very basic, you can make even better apps.
<img width="813" height="508" alt="GUI demo" src="https://github.com/user-attachments/assets/c8e48374-7e98-41c5-9d46-4427a007b02b" />

<img width="683" height="360" alt="Example app 2" src="https://github.com/user-attachments/assets/d536a863-d010-4dc4-b223-b790b00c8478" />

# Logo

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

### Built-in Palettes
```go
a.ApplyPalette(proton.DarkPalette)
a.ApplyPalette(proton.NordPalette)
a.ApplyPalette(proton.RosePinePalette)
a.ApplyPalette(proton.CatppuccinPalette)
```

### Background Color
You can set a custom background color directly (infinite options):
```go
a.SetBackground(proton.RGB(0x1a1b26)) // Hex code
a.SetBackground(proton.RGBA(20, 20, 20, 255)) // RGBA values
```

### Custom Palette Properties
The `proton.Palette` struct controls these areas:
| Property | Description |
|---|---|
| `Bg` | Main window background color |
| `Fg` | Default text and icon color |
| `Primary` | Accent color (Buttons, Sliders, Progress) |
| `PrimaryFg` | Text color inside primary elements |

### Copy-Paste Theme Presets

**Hacker Green (Matrix Style)**
```go
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x000000), // Black
    Fg:        proton.RGB(0x00FF00), // Hacker Green
    Primary:   proton.RGB(0x008F11), // Dark Green
    PrimaryFg: proton.RGB(0x000000),
})
```

**Midnight Ocean**
```go
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x0f172a), // Navy
    Fg:        proton.RGB(0xf8fafc), // Off-white
    Primary:   proton.RGB(0x38bdf8), // Sky Blue
    PrimaryFg: proton.RGB(0x0f172a),
})
```

**Cyberpunk Red**
```go
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x1a0b0b), // Deep Red-Black
    Fg:        proton.RGB(0xff2a6d), // Neon Pink
    Primary:   proton.RGB(0xd1ff00), // Neon Lime
    PrimaryFg: proton.RGB(0x000000),
})
```

**Font Scaling**
```go
a.SetFontScale(1.1) 
```

## State types

Declare these in your UI state struct — no imports beyond `proton` needed:

```go
type UI struct {
    btn     proton.Clickable   
    name    proton.Editor      
    checked proton.Bool        
    choice  proton.Enum       
    vol     proton.Float       
    scroll  proton.Scrollable  
}
```

## Examples

```
go run ./examples/hello
go run ./examples/todo
go run ./examples/calculator
go run ./examples/showcase
go run ./examples/cybertool
```

## Building a Self-Contained App

To build a standalone executable for your OS, run:

```bash
# Windows
go build -o CyberTool.exe ./examples/cybertool/main.go

# Linux/macOS
go build -o CyberTool ./examples/cybertool/main.go
```

For cross-platform distribution, you can use the provided `build.sh` or `build.bat` scripts to generate binaries for Windows, macOS, and Linux simultaneously.

## User Installation Guide (From Scratch)

If you are a new user and want to run Proton apps, follow these steps:

### 1. Install Go
Download and install Go (1.22+) from [go.dev](https://go.dev/dl/).

### 2. Install System Dependencies
Only required for **Linux** users:
```bash
sudo apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

### 3. Clone and Run
```bash
# Clone the repository
git clone https://github.com/CzaxStudio/proton.git
cd proton

# Download dependencies
go mod tidy

# Run the CyberTool example
go run ./examples/cybertool
```

### 4. Direct Install (Into your own project)
To use Proton in your own project:
```bash
mkdir myapp && cd myapp
go mod init myapp
go get github.com/CzaxStudio/proton
```
Copy any example code into `main.go` and run `go run .`.

## Proton CLI

Proton includes a CLI tool for embedding assets.

### Install

```bash
go install github.com/CzaxStudio/proton/cmd/Proton@latest
```

### Adding a logo

To add a logo to your project:

```bash
Proton logo path/to/image.png
```

This copies the image and creates `logo_gen.go`. Usage in code:

```go
func main() {
    a := proton.New("My App")
    a.SetLogo(Logo_data) // loads image once at startup

    a.Window("My App", 400, 300, func(win *proton.Win) {
        proton.Logo(win, 64, 64) // uses cached logo
        proton.H3(win, "Welcome")
    })
    a.Run()
}
```

## License

MIT
