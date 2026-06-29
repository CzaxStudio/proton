# Layout

By default, every widget you call in your draw function stacks vertically —
top to bottom, one after another. For anything more interesting, Proton has
a full set of layout helpers.

---

## Gap — Space Between Things

The most-used layout function. Inserts a blank space of `dp` density-independent
pixels. Use it constantly.

```go
proton.H4(win, "Section Title")
proton.Gap(win, 8)    // a little breathing room
proton.Label(win, "Section content here")
proton.Gap(win, 24)   // a bigger gap before the next section
proton.H4(win, "Next Section")
```

**Signature:**
```go
proton.Gap(win Context, dp float32)
```

A `dp` value of `8` is a small gap, `16` is medium, `24` is large.
Those three cover most cases.

---

## Row — Side by Side

Places widgets horizontally, left to right. Each widget takes only as much
width as it needs.

```go
proton.Row(win,
    func(win proton.Context) { proton.Label(win, "Name:") },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) { proton.Label(win, "Alice") },
)
```

Every child is a `func(Context)`. Inside it, call whatever widgets you want.

**Signature:**
```go
proton.Row(win Context, widgets ...func(Context))
```

---

## Column — Explicit Vertical Group

Stacks widgets vertically, but as a named group. You usually don't need
this at the top level (widgets stack automatically), but it's useful inside
layouts like `Row` or `Split` where you want the right side to be a vertical
stack of things.

```go
proton.Row(win,
    func(win proton.Context) {
        // left side: just a label
        proton.Label(win, "Left")
    },
    func(win proton.Context) {
        // right side: a vertical stack
        proton.Column(win,
            func(win proton.Context) { proton.Label(win, "Right Top") },
            func(win proton.Context) { proton.Label(win, "Right Bottom") },
        )
    },
)
```

**Signature:**
```go
proton.Column(win Context, widgets ...func(Context))
```

---

## RowSpread — Space Between Items

Like `Row` but puts any leftover horizontal space between the children,
pushing the first item to the left edge and the last to the right.

```go
// title on the left, "v1.0" on the right
proton.RowSpread(win,
    func(win proton.Context) { proton.H5(win, "My App") },
    func(win proton.Context) { proton.Caption(win, "v1.0") },
)
```

**Signature:**
```go
proton.RowSpread(win Context, widgets ...func(Context))
```

---

## RowEnd — Everything on the Right

Pushes all children to the right edge. Classic placement for action buttons.

```go
proton.RowEnd(win,
    func(win proton.Context) {
        if proton.OutlineButton(win, &u.cancel, "Cancel") { handleCancel() }
    },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) {
        if proton.Button(win, &u.save, "Save") { handleSave() }
    },
)
```

**Signature:**
```go
proton.RowEnd(win Context, widgets ...func(Context))
```

---

## GrowRow and GrowColumn — Stretchy Layouts

When you need one child to fill all the remaining space and the others to
stay their natural size, use `GrowRow` (horizontal) or `GrowColumn` (vertical)
with `GrowItem` and `FixedItem`.

```go
// input stretches to fill space, label and button stay fixed size
proton.GrowRow(win,
    proton.FixedItem(win, func(win proton.Context) {
        proton.Label(win, "Search:")
    }),
    proton.GrowItem(win, func(win proton.Context) {
        proton.Input(win, &u.search, "Type to search...")
    }),
    proton.FixedItem(win, func(win proton.Context) {
        if proton.Button(win, &u.searchBtn, "Go") { doSearch() }
    }),
)
```

`GrowItem` takes all remaining space. `FixedItem` takes only what it needs.
You can have multiple `GrowItem`s — they split the remaining space evenly.

**Signatures:**
```go
proton.GrowRow(win proton.Context, children ...proton.FlexItem)
proton.GrowColumn(win proton.Context, children ...proton.FlexItem)
proton.GrowItem(win proton.Context, fn func(proton.Context)) proton.FlexItem
proton.FixedItem(win proton.Context, fn func(proton.Context)) proton.FlexItem
```

---

## Split — Side-by-Side Panes

Divides the available width between two sections. `leftFraction` is the
proportion the left side gets (0.0–1.0).

```go
// left gets 30%, right gets 70%
proton.Split(win, 0.3,
    func(win proton.Context) {
        // sidebar
        proton.Label(win, "Navigation")
    },
    func(win proton.Context) {
        // main content
        proton.Label(win, "Content area")
    },
)
```

**Signature:**
```go
proton.Split(win Context, leftFraction float32, left func(Context), right func(Context))
```

### HSplit — Top and Bottom

Same idea but vertical. `topFraction` is the proportion the top gets.

```go
proton.HSplit(win, 0.7,
    func(win proton.Context) { proton.Label(win, "Main content") },
    func(win proton.Context) { proton.Label(win, "Status bar") },
)
```

**Signature:**
```go
proton.HSplit(win Context, topFraction float32, top func(Context), bottom func(Context))
```

---

## Center — Put Something in the Middle

Centers its content in the available space.

```go
proton.Center(win, func(win proton.Context) {
    proton.H3(win, "Nothing to see here yet")
})
```

Great for empty states — when a list is empty, a dashboard has no data, etc.

**Signature:**
```go
proton.Center(win Context, fn func(Context))
```

---

## Padding

### Pad — All Sides

Adds padding on all four sides.

```go
proton.Pad(win, 16, func(win proton.Context) {
    proton.Label(win, "16dp of space around me")
})
```

### PadH — Left and Right Only

```go
proton.PadH(win, 24, func(win proton.Context) {
    proton.Label(win, "24dp of horizontal padding")
})
```

### PadV — Top and Bottom Only

```go
proton.PadV(win, 12, func(win proton.Context) {
    proton.Label(win, "12dp of vertical padding")
})
```

### PadSides — Each Edge Individually

Arguments are top, right, bottom, left — same order as CSS.

```go
proton.PadSides(win, 8, 16, 8, 16, func(win proton.Context) {
    proton.Label(win, "8 top/bottom, 16 left/right")
})
```

**Signatures:**
```go
proton.Pad(win Context, dp float32, fn func(Context))
proton.PadH(win Context, dp float32, fn func(Context))
proton.PadV(win Context, dp float32, fn func(Context))
proton.PadSides(win Context, top, right, bottom, left float32, fn func(Context))
```

---

## Grid — Fixed-Column Grid

Arranges children in a grid with a fixed number of columns.
Each cell gets an equal share of the width.

```go
proton.Grid(win, 3, 8,   // 3 columns, 8dp gap
    func(win proton.Context) { proton.Label(win, "Cell 1") },
    func(win proton.Context) { proton.Label(win, "Cell 2") },
    func(win proton.Context) { proton.Label(win, "Cell 3") },
    func(win proton.Context) { proton.Label(win, "Cell 4") },
    func(win proton.Context) { proton.Label(win, "Cell 5") },
)
```

The cells wrap onto new rows automatically. If the last row has fewer
than `cols` cells, the remaining cells are just empty.

**Signature:**
```go
proton.Grid(win Context, cols int, gapDp float32, cells ...func(Context))
```

---

## Putting It All Together

A typical two-column app layout:

```go
func draw(win proton.Context, u *UI) {
    // header
    proton.PadSides(win, 0, 0, 12, 0, func(win proton.Context) {
        proton.RowSpread(win,
            func(win proton.Context) { proton.H5(win, "My App") },
            func(win proton.Context) { proton.Caption(win, "v1.0") },
        )
    })
    proton.Divider(win)
    proton.Gap(win, 16)

    // two-column body
    proton.Split(win, 0.35,
        func(win proton.Context) { drawSidebar(win, u) },
        func(win proton.Context) {
            proton.PadH(win, 16, func(win proton.Context) {
                drawContent(win, u)
            })
        },
    )
}
```
