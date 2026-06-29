# Visual Widgets

Shapes, cards, images, badges, and dividers.
The things that make your app look like more than a form.

---

## Divider

A thin horizontal line. Use it to separate sections visually.

```go
proton.H5(win, "Section One")
proton.Gap(win, 8)
proton.Label(win, "Some content...")
proton.Gap(win, 12)
proton.Divider(win)
proton.Gap(win, 12)
proton.H5(win, "Section Two")
```

The divider uses the theme's foreground color at 1dp height.

**Signature:**
```go
proton.Divider(win Context)
```

---

## Rect — Filled Rectangle

Draws a solid colored rectangle. Pass `0` for width or height to fill
the available space on that axis.

```go
// 100dp wide, 4dp tall red bar
proton.Rect(win, proton.RGB(0xff4444), 100, 4)

// fill full available width, 2dp tall
proton.Rect(win, proton.RGB(0x333333), 0, 2)

// fill all available space (both 0)
proton.Rect(win, proton.RGB(0x1a1a2e), 0, 0)
```

**Signature:**
```go
proton.Rect(win Context, c color.NRGBA, widthDp, heightDp float32)
```

---

## RoundRect — Rounded Rectangle

Same as Rect but with rounded corners. The `radiusDp` parameter controls
how round the corners are.

```go
// 200dp wide, 60dp tall, 12dp corner radius
proton.RoundRect(win, proton.RGB(0x2a2a3e), 200, 60, 12)

// fill width, 80dp tall, pill-shaped (radius = height/2)
proton.RoundRect(win, proton.RGB(0x4c566a), 0, 40, 20)
```

**Signature:**
```go
proton.RoundRect(win Context, c color.NRGBA, widthDp, heightDp, radiusDp float32)
```

---

## Card

A content container with a colored rounded-rectangle background.
Everything drawn inside the callback sits on top of the background.

```go
proton.Card(win, proton.RGB(0x2a2a3e), 12, 16, func(win proton.Context) {
    proton.H6(win, "Card Title")
    proton.Gap(win, 4)
    proton.Label(win, "This content sits on the card background.")
    proton.Gap(win, 8)
    if proton.Button(win, &u.btn, "Card Button") {
        doThing()
    }
})
```

**Signature:**
```go
proton.Card(win Context, bg color.NRGBA, cornerDp, padDp float32, content func(Context))
```

**Parameters:**
- `bg` — background color of the card
- `cornerDp` — corner radius in dp (8–16 looks good for most cards)
- `padDp` — padding between the card edge and the content
- `content` — draw function for what goes inside the card

### Cards as List Items

Cards work really well as list items:

```go
type Project struct {
    Name        string
    Description string
    Status      string
}

proton.List(win, &scroll, len(projects), func(win proton.Context, i int) {
    p := projects[i]
    proton.PadV(win, 4, func(win proton.Context) {
        proton.Card(win, proton.RGB(0x1e1e2e), 8, 12, func(win proton.Context) {
            proton.Label(win, p.Name)
            proton.Gap(win, 2)
            proton.Caption(win, p.Description)
            proton.Gap(win, 6)
            proton.Badge(win, proton.RGB(0x4c566a), proton.RGB(0xeceff4), p.Status)
        })
    })
})
```

---

## Badge

A small rounded chip with text inside. Good for status labels, tags, counts.

```go
// background color, text color, text
proton.Badge(win, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "Active")
proton.Badge(win, proton.RGB(0xa3be8c), proton.RGB(0x2e3440), "Done")
proton.Badge(win, proton.RGB(0xbf616a), proton.RGB(0xeceff4), "Error")
```

**Signature:**
```go
proton.Badge(win Context, bg color.NRGBA, fg color.NRGBA, text string)
```

Badges with `Row` for a row of tags:

```go
proton.Row(win,
    func(win proton.Context) { proton.Badge(win, proton.RGB(0x5e81ac), proton.RGB(0xeceff4), "Go") },
    func(win proton.Context) { proton.Gap(win, 4) },
    func(win proton.Context) { proton.Badge(win, proton.RGB(0xa3be8c), proton.RGB(0x2e3440), "GUI") },
    func(win proton.Context) { proton.Gap(win, 4) },
    func(win proton.Context) { proton.Badge(win, proton.RGB(0xebcb8b), proton.RGB(0x2e3440), "v1.0") },
)
```

---

## Images

Load once at startup, draw every frame.

```go
// at startup — not in the draw function
img, err := proton.LoadImage("logo.png")
if err != nil {
    log.Fatal("couldn't load logo:", err)
}

// in the draw function — as many times as you want
proton.Image(win, img, 200, 150)   // 200dp wide, 150dp tall

// use the image's natural pixel size
proton.Image(win, img, 0, 0)
```

**Signatures:**
```go
proton.LoadImage(path string) (proton.ImageOp, error)
proton.Image(win Context, img proton.ImageOp, widthDp, heightDp float32)
```

Supported formats: PNG and JPEG.

`LoadImage` does disk I/O and memory allocation — only call it once, not
every frame. Store the returned `proton.ImageOp` in your app state.

---

## Size Constraints

### MinSize

Forces content to be at least `widthDp` × `heightDp`. Pass `0` on either
axis to leave it unconstrained.

```go
// buttons that are always at least 120dp wide and 40dp tall
proton.MinSize(win, 120, 40, func(win proton.Context) {
    if proton.Button(win, &u.btn, "OK") { handleOK() }
})
```

### MaxWidth

Caps content width at `widthDp`. Good for keeping forms readable on wide windows.

```go
// form never wider than 400dp, even on a widescreen
proton.MaxWidth(win, 400, func(win proton.Context) {
    proton.Input(win, &u.email, "Email address")
    proton.Gap(win, 8)
    proton.Input(win, &u.password, "Password")
})
```

**Signatures:**
```go
proton.MinSize(win Context, widthDp, heightDp float32, fn func(Context))
proton.MaxWidth(win Context, widthDp float32, fn func(Context))
```

---

## Colors Quick Reference

```go
// from hex
proton.RGB(0x1e1e2e)   // dark background
proton.RGB(0xcdd6f4)   // light text
proton.RGB(0x89b4fa)   // blue accent
proton.RGB(0xa6e3a1)   // green
proton.RGB(0xf38ba8)   // red/pink
proton.RGB(0xfab387)   // orange

// with alpha
proton.RGBA(0, 0, 0, 128)       // 50% transparent black overlay
proton.RGBA(255, 255, 255, 20)  // very subtle white highlight
```
