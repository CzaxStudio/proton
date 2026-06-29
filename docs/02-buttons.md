# Buttons

Buttons are how users tell your app to do things. Proton has three kinds,
plus a way to make literally anything clickable.

---

## The State Field Rule

Every button needs a `proton.Clickable` field in your state struct.
One button, one field. Don't share them between buttons — they'll both
respond to the same clicks, which is usually not what you want.

```go
type UI struct {
    save    proton.Clickable
    cancel  proton.Clickable
    delete  proton.Clickable
}
```

---

## Button — The Standard One

A filled, solid button. Use this for primary actions — the thing you most
want the user to do.

```go
var save proton.Clickable

if proton.Button(win, &save, "Save") {
    doSave()
}
```

`Button` returns `true` on the frame it gets clicked. Put your logic inside
the `if`. It won't fire continuously — just once per actual click.

**Signature:**
```go
proton.Button(win Context, state *proton.Clickable, label string) bool
```

---

## OutlineButton — The Secondary One

Same as Button but without the filled background. Use it for secondary
actions — things the user might want to do, but aren't the main thing.

```go
var save   proton.Clickable
var cancel proton.Clickable

proton.Row(win,
    func(win proton.Context) {
        if proton.Button(win, &save, "Save") {
            doSave()
        }
    },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) {
        if proton.OutlineButton(win, &cancel, "Cancel") {
            doCancel()
        }
    },
)
```

The visual hierarchy here is clear: Save is the primary action (filled),
Cancel is secondary (outline). Users know what to click.

**Signature:**
```go
proton.OutlineButton(win Context, state *proton.Clickable, label string) bool
```

---

## IconButton — Just an Icon

For when a label would take up too much space and the icon is obvious enough
on its own. Common in toolbars.

```go
// icon is a *proton.Icon — see gioui.org/widget for icon data, or use
// a pre-made icon set; wrap it with widget.NewIcon() then store as proton.Icon
var closeBtn proton.Clickable

if proton.IconButton(win, &closeBtn, closeIcon, "Close") {
    closeWindow()
}
```

The third argument is the icon (a `*widget.Icon`), the fourth is a text
description used for accessibility.

**Signature:**
```go
proton.IconButton(win Context, state *proton.Clickable, icon *widget.Icon, desc string) bool
```

---

## Tappable — Make Anything Clickable

Sometimes you want to click on something that isn't a standard button —
a card, a list row, a custom-drawn area. `Tappable` wraps any content in
a click handler.

```go
var rowClick proton.Clickable

if proton.Tappable(win, &rowClick, func(win proton.Context) {
    // draw whatever you want in here
    proton.Card(win, proton.RGB(0x2a2a2a), 8, 12, func(win proton.Context) {
        proton.Label(win, "Click anywhere on this card")
    })
}) {
    println("card was clicked!")
}
```

The whole area drawn inside the callback becomes the hit target.

**Signature:**
```go
proton.Tappable(win Context, state *proton.Clickable, content func(Context)) bool
```

---

## Buttons Inside Layouts

Buttons must be placed inside a layout helper (`Row`, `Pad`, `Column`, etc.)
for clicks to register properly. This is how Gio's input system works under
the hood — the layout pass is what establishes where hit areas are on screen.

```go
// Correct — button is inside Pad
proton.Pad(win, 8, func(win proton.Context) {
    if proton.Button(win, &btn, "Click me") {
        doThing()
    }
})

// Also correct — button is inside Row
proton.Row(win,
    func(win proton.Context) {
        if proton.Button(win, &btn, "Click me") {
            doThing()
        }
    },
)
```

At the top level of your draw function, buttons will display but their
clicks may not register. Always wrap them.

---

## Common Patterns

### Confirmation row (Save + Cancel)
```go
type UI struct {
    save   proton.Clickable
    cancel proton.Clickable
}

proton.RowEnd(win,
    func(win proton.Context) {
        if proton.OutlineButton(win, &u.cancel, "Cancel") {
            handleCancel()
        }
    },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) {
        if proton.Button(win, &u.save, "Save") {
            handleSave()
        }
    },
)
```

`RowEnd` pushes everything to the right edge — standard placement for
confirm/cancel pairs.

### Toolbar
```go
type UI struct {
    newFile  proton.Clickable
    openFile proton.Clickable
    saveFile proton.Clickable
}

proton.Row(win,
    func(win proton.Context) {
        if proton.Button(win, &u.newFile, "New") { handleNew() }
    },
    func(win proton.Context) { proton.Gap(win, 4) },
    func(win proton.Context) {
        if proton.Button(win, &u.openFile, "Open") { handleOpen() }
    },
    func(win proton.Context) { proton.Gap(win, 4) },
    func(win proton.Context) {
        if proton.Button(win, &u.saveFile, "Save") { handleSave() }
    },
)
```

### Clickable list row
```go
type UI struct {
    rows [5]proton.Clickable
}

items := []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
var scroll proton.Scrollable

proton.List(win, &scroll, len(items), func(win proton.Context, i int) {
    if proton.Tappable(win, &u.rows[i], func(win proton.Context) {
        proton.PadV(win, 10, func(win proton.Context) {
            proton.Label(win, items[i])
        })
    }) {
        fmt.Println("selected:", items[i])
    }
})
```
