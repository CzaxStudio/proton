# Advanced Topics

Keyboard shortcuts, async updates, toast notifications, tooltips, and
multiple windows. The stuff you'll need once your app grows beyond a single form.

---

## Toast Notifications

A timed message that appears, stays for a few seconds, and disappears.
No dialog box, no blocking the user. Just a small notification.

### Setup

Declare a `proton.ToastState` in your state struct:

```go
type UI struct {
    toast proton.ToastState
}
```

### Showing a Toast

Call `Show` from anywhere — including goroutines. It's goroutine-safe.

```go
// show for 2 seconds
u.toast.Show("File saved!", 2*time.Second)

// show for longer
u.toast.Show("Processing complete", 4*time.Second)
```

### Drawing the Toast

Call `proton.Toast` at the **end** of your draw function. This ensures it
renders on top of everything else.

```go
func draw(win proton.Context, u *UI) {
    // ... all your other widgets ...

    // always last
    proton.Toast(win, &u.toast)
}
```

If there's no active toast, `Toast` draws nothing. No need to check.

### Full Example

```go
type UI struct {
    saveBtn proton.Clickable
    toast   proton.ToastState
}

func draw(win proton.Context, u *UI) {
    proton.Pad(win, 16, func(win proton.Context) {
        if proton.Button(win, &u.saveBtn, "Save") {
            go func() {
                saveToFile()
                u.toast.Show("Saved!", 2*time.Second)
                win.Invalidate()
            }()
        }
    })

    proton.Toast(win, &u.toast)
}
```

---

## Keyboard Shortcuts

Register a function to run when a key combination is pressed.

```go
// Ctrl+S
proton.OnKey(win, proton.ModCtrl, "S", func() {
    save()
})

// Ctrl+Z
proton.OnKey(win, proton.ModCtrl, "Z", func() {
    undo()
})

// Just Escape (no modifier)
proton.OnKey(win, proton.ModNone, proton.KeyEscape, func() {
    closeDialog()
})

// Ctrl+Shift+N
proton.OnKey(win, proton.ModCtrl|proton.ModShift, "N", func() {
    newWindow()
})
```

Call `OnKey` inside your draw function. It registers the shortcut for
that frame. Since the draw function runs every frame, the shortcut is
always active while the window is open.

**Signature:**
```go
proton.OnKey(win proton.Context, modifiers proton.Modifier, keyName string, fn func())
```

**Modifiers:** `proton.ModNone`, `proton.ModCtrl`, `proton.ModShift`, `proton.ModAlt`.
Combine with `|`, e.g. `proton.ModCtrl|proton.ModShift`.

**Common key names** — letter keys are just their uppercase string ("S", "Z", "N").
For everything else, use the Key* constants:
- `proton.KeyEscape` — Escape
- `proton.KeyReturn` — Enter
- `proton.KeyBackspace` — Backspace
- `proton.KeyDelete` — Delete
- `proton.KeyTab`, `proton.KeySpace`, `proton.KeyUp`, `proton.KeyDown`, `proton.KeyLeft`, `proton.KeyRight`
- Letter keys: just the letter as a string — `"S"`, `"Z"`, `"N"`

---

## Async Updates and Goroutines

Your draw function runs on the main thread. Goroutines run elsewhere.
When a goroutine finishes work and changes your state, you need to tell
Proton to redraw. Use `win.Invalidate()`.

```go
type UI struct {
    loading bool
    result  string
}

func draw(win proton.Context, u *UI) {
    if u.loading {
        proton.Label(win, "Loading...")
        proton.ProgressBar(win, 0.5)   // indeterminate look
        return
    }

    if u.result != "" {
        proton.Label(win, u.result)
        return
    }

    proton.Pad(win, 16, func(win proton.Context) {
        if proton.Button(win, &u.fetchBtn, "Fetch Data") {
            u.loading = true
            go func() {
                data := fetchFromAPI()       // takes a while
                u.result = data
                u.loading = false
                win.Invalidate()             // ask for a redraw
            }()
        }
    })
}
```

Without `win.Invalidate()`, the window won't redraw until the user moves
the mouse or interacts with it. Always call it after changing state from
a goroutine.

**Signature:**
```go
win.Invalidate()
```

---

## Tooltip

Shows a small label when the user hovers over something. Useful for
explaining what a button does without cluttering the UI.

```go
type UI struct {
    saveHover proton.Clickable
    saveBtn   proton.Clickable
}

proton.Tooltip(win, &u.saveHover, "Saves your work to disk (Ctrl+S)", func(win proton.Context) {
    if proton.Button(win, &u.saveBtn, "Save") {
        doSave()
    }
})
```

The tooltip appears below the content when hovered. It disappears when
the user moves the mouse away.

**Signature:**
```go
proton.Tooltip(win Context, state *proton.Clickable, tip string, content func(Context))
```

Note: `Tooltip` uses a `proton.Clickable` for hover tracking. This is separate
from any button state inside the content. Declare a dedicated `Clickable`
for the tooltip itself.

---

## Multiple Windows

Register multiple windows with `a.Window()`. They all open when `a.Run()` is
called. The app stays alive until all windows are closed.

```go
func main() {
    u := &UI{}

    a := proton.New("my app")
    a.ApplyPalette(proton.NordPalette)

    a.Window("Main Window", 800, 600, func(win proton.Context) {
        drawMain(win, u)
    })

    a.Window("Settings", 400, 300, func(win proton.Context) {
        drawSettings(win, u)
    })

    a.Run()
}
```

Both windows share the same `u` state struct, so changes in one window
are immediately visible in the other on the next frame.

---

## Window Options

For windows that need special behavior, use `WindowEx`:

```go
a.WindowEx("My App", 800, 600, []proton.WindowOption{
    proton.Fullscreen(),
}, draw)
```

`WindowEx` accepts Proton's own option constructors:
- `proton.Fullscreen()` — start fullscreen
- `proton.Maximized()` — start maximized

---

## Keeping the Frame Rate Reasonable

Proton only redraws when there's user input or you call `win.Invalidate()`.
It doesn't spin at 60fps when nothing is happening — Gio is smart about this.

If you're animating something (a progress bar that fills over time, a loading
spinner, etc.), call `win.Invalidate()` at the end of each frame to keep
the redraws going:

```go
func draw(win proton.Context, u *UI) {
    if u.animating {
        u.progress += 0.016   // roughly one frame at 60fps
        if u.progress >= 1.0 {
            u.progress = 0
            u.animating = false
        }
        proton.ProgressBar(win, u.progress)
        win.Invalidate()   // keep drawing next frame
    }
}
```

When `u.animating` becomes false, `Invalidate` stops being called and
Proton goes back to only redrawing on user input.

---

## FocusArea — Scoped Key Handling

If you need keyboard events only in a specific part of the UI (not globally),
use `FocusArea`. It registers a region as a key event receiver.

```go
type UI struct {
    editorTag proton.FrameTag
}

proton.FocusArea(win, &u.editorTag, "A", func(win proton.Context) {
    proton.TextArea(win, &u.text, "Type here...")
})
```

For most apps `OnKey` is enough. `FocusArea` is for when you have multiple
panels and only want keyboard shortcuts active in the focused one.

---

## Why Proton Hides Gio Completely

Every function that takes a `proton.Context` never receives or returns
a raw Gio type. This is deliberate: Gio's layout and event APIs have
changed more than once as the library matures, and code written against
`proton.Context` keeps compiling across those changes without edits.

If you ever hit something Proton genuinely doesn't expose, that's worth
reporting — Proton is meant to cover the full surface you need without
reaching past it.
