# Text

Proton has nine text functions. They cover every size from "I want people to
see this from across the room" to "please nobody read this".

---

## Headings — H1 through H6

```go
proton.H1(win, "I am enormous")
proton.H2(win, "I am large")
proton.H3(win, "I am medium-large")
proton.H4(win, "I am medium")
proton.H5(win, "I am medium-small")
proton.H6(win, "I am small but still a heading, I have dignity")
```

These follow the same size scale as HTML headings. H1 is the biggest,
H6 is the smallest. Use H3 or H4 for most app titles and section headers.
H1 is for when you really want to make a statement.

**Signature:**
```go
proton.H1(win Context, text string)
// same for H2, H3, H4, H5, H6
```

---

## Label

Regular body text. This is what you use for most of the words in your UI.

```go
proton.Label(win, "This is a label. It's medium-sized and readable.")
```

**Signature:**
```go
proton.Label(win Context, text string)
```

---

## Body2

Slightly smaller than Label. Good for secondary information that's important
but shouldn't compete with the main content.

```go
proton.Label(win, "Main price: $29.99")
proton.Body2(win, "Taxes and fees may apply. Or not. We'll see.")
```

**Signature:**
```go
proton.Body2(win Context, text string)
```

---

## Caption

The smallest text. Use it for hints, fine print, timestamps, or anything
you want to communicate while maintaining plausible deniability.

```go
proton.Caption(win, "Last updated: 3 years ago")
proton.Caption(win, "* Results may vary")
```

**Signature:**
```go
proton.Caption(win Context, text string)
```

---

## Text — Custom Styling

When the preset sizes don't cut it, `Text` lets you set the size, color,
and weight yourself.

```go
// size in sp (scaled pixels), color, bold true/false
proton.Text(win, "BIG RED BOLD", 32, proton.RGB(0xff4444), true)

// pass a zero color to use the theme's default text color
proton.Text(win, "just bigger", 24, proton.NRGBA{}, false)
```

**Signature:**
```go
proton.Text(win Context, s string, size float32, c color.NRGBA, bold bool)
```

**Parameters:**
- `s` — the string to display
- `size` — font size in sp (scaled pixels). The default body size is around 14sp
- `c` — text color as `color.NRGBA`. Pass `color.NRGBA{}` (zero value) to use the theme's foreground color
- `bold` — `true` makes it bold, `false` keeps it regular weight

---

## Colors for Text

Proton gives you two helpers for building colors:

```go
// RGB from a hex value — the # from CSS, but as a Go integer
proton.RGB(0xff6b6b)   // a nice coral red
proton.RGB(0x4ecdc4)   // a nice teal
proton.RGB(0xffffff)   // white
proton.RGB(0x000000)   // black

// RGBA with explicit alpha (0 = transparent, 255 = fully opaque)
proton.RGBA(255, 107, 107, 255)   // same coral red, fully opaque
proton.RGBA(0, 0, 0, 128)         // semi-transparent black
```

---

## Wrapping

Long text wraps automatically within the available width. You don't need
to do anything special — just write your string and Proton handles it.

```go
proton.Label(win, "This is a very long piece of text that will wrap "+
    "onto multiple lines when the window is too narrow to fit it all.")
```

---

## Putting It Together

A typical header section might look like:

```go
proton.H4(win, "Account Settings")
proton.Gap(win, 4)
proton.Body2(win, "Manage your profile and preferences.")
proton.Gap(win, 16)
proton.Divider(win)
proton.Gap(win, 16)
```

The `Gap` calls add vertical space between elements.
See [04-layout.md](./04-layout.md) for more on spacing.
