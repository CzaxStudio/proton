# Proton Docs

Everything you need to build desktop apps with Proton.
Pick a topic or read them in order — both work fine.

---

| File | What's in it |
|---|---|
| [00-getting-started.md](./00-getting-started.md) | install, first window, the state struct pattern |
| [01-text.md](./01-text.md) | Label, H1–H6, Body2, Caption, custom Text |
| [02-buttons.md](./02-buttons.md) | Button, OutlineButton, IconButton, Tappable |
| [03-inputs.md](./03-inputs.md) | Input, TextArea, Checkbox, Toggle, RadioButton, Slider, ProgressBar |
| [04-layout.md](./04-layout.md) | Row, Column, Split, Pad, Gap, Grid, GrowRow, Center |
| [05-lists.md](./05-lists.md) | List, HList, Scroll |
| [06-visuals.md](./06-visuals.md) | Divider, Rect, RoundRect, Card, Badge, Image, MinSize, MaxWidth |
| [07-theming.md](./07-theming.md) | palettes, custom colors, font scale |
| [08-advanced.md](./08-advanced.md) | Toast, OnKey, goroutines, Tooltip, multiple windows |
| [09-examples.md](./09-examples.md) | complete copy-paste examples |

---

## The One Thing to Know

Proton is immediate mode. Your draw function runs every frame. You call
widget functions, they appear on screen in that order. State lives in
your own struct. That's it.

```go
type UI struct {
    btn proton.Clickable
}

u := &UI{}

a.Window("App", 400, 300, func(win proton.Context) {
    proton.Label(win, "Click the button.")
    proton.Gap(win, 8)
    proton.Pad(win, 8, func(win proton.Context) {
        if proton.Button(win, &u.btn, "Hello") {
            println("hello!")
        }
    })
})
```

Everything else in the docs is details.
