# Getting Started with Proton

Welcome. You want to build a desktop GUI in Go. You've made a good choice.
Let's get you from zero to a running window in about 5 minutes.

---

## Prerequisites

You need Go 1.22 or newer. Check with:

```bash
go version
```

If that prints something old or nothing at all, go to [go.dev/dl](https://go.dev/dl) and fix that first.

**Linux users** — you also need three system packages. macOS and Windows users
can skip this and feel superior:

```bash
sudo apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

---

## Installing Proton

In your project folder:

```bash
go get github.com/CzaxStudio/proton
```

---

## Your First Window

Create a `main.go`:

```go
package main

import "github.com/CzaxStudio/proton"

func main() {
    a := proton.New("my app")
    a.Window("Hello World", 400, 300, func(win proton.Context) {
        proton.H3(win, "Hello from Proton!")
    })
    a.Run()
}
```

Run it:

```bash
go run .
```

A window appears. That's it. No XML, no interface implementations, no
forty lines of boilerplate. Just Go.

---

## A Slightly More Interesting Example

Let's add a text input and a button:

```go
package main

import (
    "fmt"
    "github.com/CzaxStudio/proton"
)

// UI holds all the widget state for the app.
// One struct, declared once, lives for the lifetime of the app.
type UI struct {
    name proton.Editor
    btn  proton.Clickable
}

func main() {
    u := &UI{}

    a := proton.New("greeter")
    a.Window("Greeter", 400, 200, func(win proton.Context) {
        proton.Input(win, &u.name, "Enter your name")
        proton.Gap(win, 8)
        if proton.Button(win, &u.btn, "Say Hello") {
            fmt.Println("Hello,", u.name.Text())
        }
    })
    a.Run()
}
```

Notice the pattern — state lives in your struct, you pass pointers to widgets.
The draw function runs every frame. When the button is clicked, `Button`
returns `true` and your code runs.

---

## How Proton Works (The Short Version)

Proton is built on [Gio](https://gioui.org), which is an immediate-mode GUI
toolkit. That means:

- Your draw function runs ~60 times per second
- You just call widget functions in order — they appear on screen in that order
- State (what was typed, what was clicked) lives in your own structs
- There are no hidden objects, no component trees, no virtual DOM

The mental model is simple: **every frame, describe what the UI should look
like right now**. Proton handles the rest.

---

## The State Struct Pattern

Every Proton app follows the same pattern:

```go
// 1. Define a struct that holds all your UI state
type UI struct {
    // one field per interactive widget
    name    proton.Editor
    submit  proton.Clickable
    checked proton.Bool
}

// 2. Create it once
u := &UI{}

// 3. Pass pointers to widgets in the draw function
a.Window("App", 400, 300, func(win proton.Context) {
    proton.Input(win, &u.name, "Name")
    proton.Checkbox(win, &u.checked, "Subscribe to newsletter")
    if proton.Button(win, &u.submit, "Submit") {
        // handle submission
    }
})
```

That's the whole pattern. You'll use it in every Proton app you write.

---

## Next Steps

- **[01-text.md](./01-text.md)** — all the text widgets
- **[02-buttons.md](./02-buttons.md)** — buttons and clickable areas
- **[03-inputs.md](./03-inputs.md)** — text fields, checkboxes, toggles, sliders
- **[04-layout.md](./04-layout.md)** — arranging things on screen
- **[05-lists.md](./05-lists.md)** — scrollable lists and scroll areas
- **[06-visuals.md](./06-visuals.md)** — shapes, cards, images, badges
- **[07-theming.md](./07-theming.md)** — colors, palettes, fonts
- **[08-advanced.md](./08-advanced.md)** — keyboard shortcuts, async updates, toast notifications
