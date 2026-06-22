# Lists and Scrolling

For displaying collections of things, and for making content areas scrollable.

---

## List — Vertical Scrollable List

The standard list. Renders items vertically, only drawing what's visible
on screen (so 10,000 items is fine — it doesn't render all 10,000).

```go
items := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
var scroll proton.Scrollable

proton.List(win, &scroll, len(items), func(win proton.Context, i int) {
    proton.Label(win, items[i])
})
```

The callback receives the `win` and the index `i`. Draw whatever you want
for each row.

**Signature:**
```go
proton.List(win Context, state *proton.Scrollable, length int, draw func(Context, int))
```

**Parameters:**
- `state` — declare one `proton.Scrollable` per list, it tracks scroll position
- `length` — total number of items
- `draw` — called once per visible item with the `Win` and the item index

---

## HList — Horizontal Scrollable List

Same as `List` but items go left to right and you scroll horizontally.
Good for image galleries, tab bars, carousels.

```go
var hscroll proton.Scrollable

proton.HList(win, &hscroll, len(thumbnails), func(win proton.Context, i int) {
    proton.Pad(win, 8, func(win proton.Context) {
        proton.Label(win, thumbnails[i].title)
    })
})
```

**Signature:**
```go
proton.HList(win Context, state *proton.Scrollable, length int, draw func(Context, int))
```

---

## Scroll — Scrollable Content Area

When you have a bunch of widgets that might not fit, wrap them in `Scroll`.
Unlike `List`, this is for arbitrary content, not indexed items.

```go
var scroll proton.Scrollable

proton.Scroll(win, &scroll, func(win proton.Context) {
    proton.H5(win, "A Very Long Page")
    proton.Gap(win, 8)
    proton.Label(win, "Paragraph one...")
    proton.Gap(win, 8)
    proton.Label(win, "Paragraph two...")
    proton.Gap(win, 8)
    // ... as many widgets as you want
})
```

**Signature:**
```go
proton.Scroll(win Context, state *proton.Scrollable, content func(Context))
```

---

## Making Lists Useful

A bare list of labels is rarely what you actually want. Here are the patterns
you'll actually use:

### Padded rows
```go
proton.List(win, &scroll, len(items), func(win proton.Context, i int) {
    proton.PadV(win, 8, func(win proton.Context) {
        proton.PadH(win, 12, func(win proton.Context) {
            proton.Label(win, items[i])
        })
    })
})
```

### Rows with two columns of text
```go
type Contact struct {
    Name  string
    Email string
}

proton.List(win, &scroll, len(contacts), func(win proton.Context, i int) {
    c := contacts[i]
    proton.PadV(win, 8, func(win proton.Context) {
        proton.Label(win, c.Name)
        proton.Gap(win, 2)
        proton.Caption(win, c.Email)
    })
})
```

### Clickable rows
```go
type UI struct {
    rows    [100]proton.Clickable
    selected int
}

proton.List(win, &scroll, len(items), func(win proton.Context, i int) {
    if proton.Tappable(win, &u.rows[i], func(win proton.Context) {
        proton.PadV(win, 10, func(win proton.Context) {
            proton.PadH(win, 12, func(win proton.Context) {
                proton.Label(win, items[i])
            })
        })
    }) {
        u.selected = i
    }
    proton.Divider(win)
})
```

### List inside a Card
```go
proton.Card(win, proton.RGB(0x1e1e2e), 10, 0, func(win proton.Context) {
    proton.List(win, &scroll, len(items), func(win proton.Context, i int) {
        proton.PadV(win, 8, func(win proton.Context) {
            proton.PadH(win, 12, func(win proton.Context) {
                proton.Label(win, items[i])
            })
        })
        if i < len(items)-1 {
            proton.Divider(win)
        }
    })
})
```

---

## The Scrollable State Field

Every `List`, `HList`, and `Scroll` call needs its own `proton.Scrollable`
in your state struct. It tracks how far the user has scrolled. Don't share
one between multiple lists — they'll fight over scroll position.

```go
type UI struct {
    contactScroll  proton.Scrollable
    messageScroll  proton.Scrollable
    sidebarScroll  proton.Scrollable
}
```

---

## Performance Note

`List` and `HList` use virtual rendering — they only call your draw function
for the items currently visible on screen. This means you can have a slice
with 100,000 items and the list will scroll smoothly. `Scroll` renders all
its content, so keep it for pages with a reasonable amount of content rather
than very long dynamic lists.
