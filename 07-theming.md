# Theming

Proton has three ways to set colors. Pick whichever fits your workflow.

---

## 1. Built-in palettes

One line. Done.

```go
a.ApplyPalette(proton.NordPalette)
```

46 palettes are included. See the full list at the bottom of this page.

---

## 2. Hex color codes

Use `a.ThemeBuilder()` to set colors with CSS hex strings — no structs,
no imports, just the hex codes you already know.

```go
// full custom theme
a.ThemeBuilder().
    Bg("#1e1e2e").
    Fg("#cdd6f4").
    Primary("#89b4fa").
    PrimaryFg("#1e1e2e").
    Apply()

// override one color on an existing palette
a.ApplyPalette(proton.NordPalette)
a.ThemeBuilder().Primary("#ff6b6b").Apply()
```

Supported formats:
- `"#rrggbb"` — standard hex with hash
- `"rrggbb"`  — hex without hash
- `"#rgb"`    — shorthand (expands to `rrggbbr`)
- `"#rrggbbaa"` — with alpha channel

Or use `a.ColorCode(slot, code)` for a single one-liner:

```go
a.ColorCode("bg",      "#0d1117")
a.ColorCode("fg",      "#e6edf3")
a.ColorCode("primary", "#1f6feb")
a.ColorCode("primaryfg", "#ffffff")
```

Slot names: `"bg"`, `"background"`, `"fg"`, `"foreground"`, `"text"`,
`"primary"`, `"accent"`, `"primaryfg"`, `"primarytext"`.

---

## 3. Palette struct with RGB helpers

For when you want to define the whole thing explicitly:

```go
a.ApplyPalette(proton.Palette{
    Bg:        proton.RGB(0x1e1e2e),
    Fg:        proton.RGB(0xcdd6f4),
    Primary:   proton.RGB(0x89b4fa),
    PrimaryFg: proton.RGB(0x1e1e2e),
})
```

Or use `MakePalette` for slightly less typing:

```go
a.ApplyPalette(proton.MakePalette(0x1e1e2e, 0xcdd6f4, 0x89b4fa, 0x1e1e2e))
```

---

## Font scale

```go
a.SetFontScale(1.1)   // 10% bigger globally
a.SetFontScale(0.9)   // a bit smaller
```

---

## Live theme picker widget

Drop this into a settings window to let users pick their own theme at runtime:

```go
type UI struct {
    picker proton.ThemePickerState
}

proton.ThemePicker(win, &u.picker, a)
```

The picker shows all 46 built-in palettes with color swatches. Clicking one
applies it to the running app immediately.

---

## All built-in palettes

### Dark themes

| Variable | Description |
|---|---|
| `DarkPalette` | Clean neutral dark |
| `NordPalette` | Arctic blue-grey |
| `RosePinePalette` | Warm muted purple |
| `RosePineMoonPalette` | Rose Pine dark moon variant |
| `CatppuccinPalette` | Catppuccin Mocha (darkest) |
| `CatppuccinFrappePalette` | Catppuccin Frappé |
| `CatppuccinMacchiatoPalette` | Catppuccin Macchiato |
| `DraculaPalette` | Classic purple |
| `GruvboxDarkPalette` | Warm earthy retro |
| `GruvboxMaterialDarkPalette` | Gruvbox with softer colors |
| `TokyoNightPalette` | Deep blue-purple |
| `TokyoNightStormPalette` | Slightly lighter Tokyo Night |
| `MonokaiPalette` | Sublime Text classic |
| `SolarizedDarkPalette` | Tinted dark, high contrast |
| `OneDarkPalette` | Atom One Dark |
| `MaterialDarkPalette` | Material Design dark |
| `AyuDarkPalette` | Ayu dark, clean modern |
| `AyuMiragePalette` | Ayu mirage, warm medium |
| `EverforestDarkPalette` | Muted green forest |
| `KanagawaPalette` | Inspired by The Great Wave |
| `VesperPalette` | Minimal warm dark |
| `NightOwlPalette` | Designed for night coding |
| `CarbonPalette` | IBM Carbon dark |
| `MidnightPalette` | Deep navy, sky blue |
| `ObsidianPalette` | Green-tinted editor dark |
| `HackerPalette` | Terminal green on black |
| `CyberpunkPalette` | Neon pink and lime |
| `OleDarkPalette` | Warm lamplight brown |
| `SlackPalette` | Slack sidebar purple |
| `TerminalGreenPalette` | Phosphor green CRT |
| `TerminalAmberPalette` | Phosphor amber CRT |
| `OceanicNextPalette` | Cool ocean blues |
| `IcebergPalette` | Cold blue-grey |
| `SynthwavePalette` | 80s retro neon |

### Light themes

| Variable | Description |
|---|---|
| `LightPalette` | Clean neutral light |
| `SolarizedLightPalette` | Warm cream |
| `RosePineDawnPalette` | Rose Pine light |
| `CatppuccinLattePalette` | Catppuccin lightest |
| `FluentLightPalette` | Microsoft Fluent |
| `PaperPalette` | Warm off-white, ink text |
| `GithubLightPalette` | GitHub's clean light |
| `AyuLightPalette` | Ayu light |
| `EverforestLightPalette` | Soft green light |
| `NordLightPalette` | Nord warm light |
| `GruvboxLightPalette` | Gruvbox cream |
| `TokyoNightDayPalette` | Tokyo Night light |

---

## Using colors in custom widgets

```go
// fixed color
proton.Rect(win, proton.RGB(0xff6b6b), 0, 4)
```

Proton doesn't expose the live theme's colors back to your code by
design — that would mean handing back a Gio type, which breaks the
API-immunity guarantee. If you need a color to match your palette,
store it yourself when you set the palette:

```go
var primaryColor = proton.RGB(0x89b4fa)

a.ApplyPalette(proton.Palette{
    Bg: proton.RGB(0x1e1e2e), Fg: proton.RGB(0xcdd6f4),
    Primary: primaryColor, PrimaryFg: proton.RGB(0x1e1e2e),
})

// later, anywhere in your draw code:
proton.Rect(win, primaryColor, 0, 4)
```

---

## Background Colors

Beyond palettes, you can set the window background directly — solid colors,
gradients, or an animated rainbow. These override the palette's `Bg` color.

### Solid color

```go
a.SetBackground(proton.RGB(0x1a1b26))
a.SetBackgroundCode("#1a1b26")
a.SetBackgroundRGB(26, 27, 38)
```

All three do the same thing — pick whichever input format is convenient.

### Gradient

```go
a.SetBackgroundGradient("#1a1b26", "#2d1b69", "vertical")
```

`dir` is one of `"horizontal"`, `"vertical"`, `"diagonal"`, or `"radial"`.

### Rainbow

```go
a.SetBackgroundRainbow()
```

An animated full-spectrum gradient that shifts slowly over time. Good for
demos, novelty apps, or just because you can.

### Precedence

If you call more than one of these, the most recent call wins. Calling
`ApplyPalette` after setting a background does not override it — backgrounds
are independent of the palette system. Call background setters last if you
want them to take priority.
