package proton

import (
	"image/color"

	"gioui.org/unit"
)

// Palette holds the four colors that define an app's visual style.
type Palette struct {
	Bg        color.NRGBA // window background
	Fg        color.NRGBA // text and icons
	Primary   color.NRGBA // buttons, sliders, accents
	PrimaryFg color.NRGBA // text drawn on primary elements
}

// ApplyPalette pushes palette colors into the app theme.
// Call after New(), before Run().
func (a *App) ApplyPalette(p Palette) {
	if p.Bg != (color.NRGBA{}) {
		a.theme.Palette.Bg = p.Bg
	}
	if p.Fg != (color.NRGBA{}) {
		a.theme.Palette.Fg = p.Fg
	}
	if p.Primary != (color.NRGBA{}) {
		a.theme.Palette.ContrastBg = p.Primary
	}
	if p.PrimaryFg != (color.NRGBA{}) {
		a.theme.Palette.ContrastFg = p.PrimaryFg
	}
}

// SetFontScale multiplies the base text size. 1.0 is the default.
func (a *App) SetFontScale(scale float32) {
	a.theme.TextSize = unit.Sp(float32(a.theme.TextSize) * scale)
}

// MakePalette builds a Palette from four hex values.
// Saves typing color.NRGBA{} structs by hand.
//
//	a.ApplyPalette(proton.MakePalette(0x1e1e2e, 0xcdd6f4, 0x89b4fa, 0x1e1e2e))
func MakePalette(bg, fg, primary, primaryFg uint32) Palette {
	return Palette{
		Bg:        RGB(bg),
		Fg:        RGB(fg),
		Primary:   RGB(primary),
		PrimaryFg: RGB(primaryFg),
	}
}

var DarkPalette = MakePalette(0x121212, 0xdcdcdc, 0x6464f0, 0xffffff)

var NordPalette = MakePalette(0x2e3440, 0xd8dee9, 0x88c0d0, 0x2e3440)

var RosePinePalette = MakePalette(0x191724, 0xe0def4, 0xc4a7e7, 0x191724)

var CatppuccinPalette = MakePalette(0x1e1e2e, 0xcdd6f4, 0x89b4fa, 0x1e1e2e)

// DraculaPalette — the classic purple dark theme.
var DraculaPalette = MakePalette(0x282a36, 0xf8f8f2, 0xbd93f9, 0x282a36)

// GruvboxDarkPalette — warm earthy tones.
var GruvboxDarkPalette = MakePalette(0x282828, 0xebdbb2, 0xd79921, 0x282828)

// GruvboxLightPalette — same palette, light background.
var GruvboxLightPalette = MakePalette(0xfbf1c7, 0x3c3836, 0xd65d0e, 0xfbf1c7)

// TokyoNightPalette — dark blue-purple, very popular in editors.
var TokyoNightPalette = MakePalette(0x1a1b26, 0xc0caf5, 0x7aa2f7, 0x1a1b26)

// TokyoNightStormPalette — slightly lighter variant of Tokyo Night.
var TokyoNightStormPalette = MakePalette(0x24283b, 0xc0caf5, 0x7aa2f7, 0x24283b)

// MonokaiPalette — the Sublime Text classic.
var MonokaiPalette = MakePalette(0x272822, 0xf8f8f2, 0xa6e22e, 0x272822)

// SolarizedDarkPalette — the original tinted dark theme.
var SolarizedDarkPalette = MakePalette(0x002b36, 0x839496, 0x268bd2, 0x002b36)

// OneDarkPalette — Atom's One Dark, ported everywhere.
var OneDarkPalette = MakePalette(0x282c34, 0xabb2bf, 0x61afef, 0x282c34)

// MaterialDarkPalette — Material Design dark.
var MaterialDarkPalette = MakePalette(0x121212, 0xe1e1e1, 0xbb86fc, 0x121212)

// AyuDarkPalette — Ayu dark, clean and modern.
var AyuDarkPalette = MakePalette(0x0d1017, 0xbfbdb6, 0xe6b450, 0x0d1017)

// AyuMiragePalette — Ayu mirage, the middle-ground variant.
var AyuMiragePalette = MakePalette(0x1f2430, 0xcccac2, 0xffcc66, 0x1f2430)

// EverforestDarkPalette — muted green forest theme.
var EverforestDarkPalette = MakePalette(0x2d353b, 0xd3c6aa, 0xa7c080, 0x2d353b)

// KanagawaPalette — inspired by Hokusai's The Great Wave.
var KanagawaPalette = MakePalette(0x1f1f28, 0xdcd7ba, 0x7e9cd8, 0x1f1f28)

// VesperPalette — minimal warm dark theme.
var VesperPalette = MakePalette(0x101010, 0xffffff, 0xffc799, 0x101010)

// NightOwlPalette — designed for night coding sessions.
var NightOwlPalette = MakePalette(0x011627, 0xd6deeb, 0x82aaff, 0x011627)

// CarbonPalette — IBM Carbon Design dark.
var CarbonPalette = MakePalette(0x161616, 0xf4f4f4, 0x0f62fe, 0xffffff)

// MidnightPalette — deep navy, cool blue accent.
var MidnightPalette = MakePalette(0x0f172a, 0xf8fafc, 0x38bdf8, 0x0f172a)

// ObsidianPalette — dark green tinted, editor-inspired.
var ObsidianPalette = MakePalette(0x293134, 0xe0e2db, 0x93c763, 0x293134)

// HackerPalette — black terminal, green text.
var HackerPalette = MakePalette(0x000000, 0x00ff00, 0x008f11, 0x000000)

// CyberpunkPalette — neon pink on near-black, lime accent.
var CyberpunkPalette = MakePalette(0x1a0b0b, 0xff2a6d, 0xd1ff00, 0x000000)

// LightPalette — clean, neutral light theme.
var LightPalette = MakePalette(0xfafafa, 0x1a1a1a, 0x5b5bd6, 0xffffff)

// SolarizedLightPalette — warm cream, the light sibling of SolarizedDark.
var SolarizedLightPalette = MakePalette(0xfdf6e3, 0x657b83, 0x268bd2, 0xfdf6e3)

// RosePineDawnPalette — the light variant of Rose Pine.
var RosePineDawnPalette = MakePalette(0xfaf4ed, 0x575279, 0xb4637a, 0xfaf4ed)

// CatppuccinLattePalette — Catppuccin's lightest variant.
var CatppuccinLattePalette = MakePalette(0xeff1f5, 0x4c4f69, 0x1e66f5, 0xffffff)

// FluentLightPalette — Microsoft Fluent design, light.
var FluentLightPalette = MakePalette(0xffffff, 0x201f1e, 0x0078d4, 0xffffff)

// PaperPalette — warm off-white, ink-like text. Easy on the eyes.
var PaperPalette = MakePalette(0xf5f0e8, 0x2c2416, 0x8b4513, 0xf5f0e8)

// GithubLightPalette — GitHub's clean light theme.
var GithubLightPalette = MakePalette(0xffffff, 0x24292f, 0x0969da, 0xffffff)

// AyuLightPalette — Ayu light variant.
var AyuLightPalette = MakePalette(0xfafafa, 0x575f66, 0xff9940, 0xfafafa)

// EverforestLightPalette — soft green light theme.
var EverforestLightPalette = MakePalette(0xfdf6e3, 0x5c6a72, 0x8da101, 0xfdf6e3)

// NordLightPalette — Nord's lighter, warmer variant.
var NordLightPalette = MakePalette(0xeceff4, 0x2e3440, 0x5e81ac, 0xeceff4)

// RosePineMoonPalette — Rose Pine's dark moon variant.
var RosePineMoonPalette = MakePalette(0x232136, 0xe0def4, 0xc4a7e7, 0x232136)

// CatppuccinFrappePalette — medium-dark Catppuccin variant.
var CatppuccinFrappePalette = MakePalette(0x303446, 0xc6d0f5, 0x8caaee, 0x303446)

// CatppuccinMacchiatoPalette — slightly darker Catppuccin.
var CatppuccinMacchiatoPalette = MakePalette(0x24273a, 0xcad3f5, 0x8aadf4, 0x24273a)

// GruvboxMaterialDarkPalette — Gruvbox with softened colors.
var GruvboxMaterialDarkPalette = MakePalette(0x292828, 0xdfbf8e, 0xa9b665, 0x292828)

// TokyoNightDayPalette — Tokyo Night but light.
var TokyoNightDayPalette = MakePalette(0xe1e2e7, 0x3760bf, 0x2e7de9, 0xffffff)

// OceanicNextPalette — cool ocean blues.
var OceanicNextPalette = MakePalette(0x1b2b34, 0xd8dee9, 0x6699cc, 0x1b2b34)

// IcebergPalette — cold blue-grey.
var IcebergPalette = MakePalette(0x161821, 0xc6c8d1, 0x84a0c6, 0x161821)

// SynthwavePalette — 80s retro neon.
var SynthwavePalette = MakePalette(0x262335, 0xffffff, 0xf92aad, 0x262335)

// OleDarkPalette — warm brown-dark, like old paper under lamplight.
var OleDarkPalette = MakePalette(0x1c1917, 0xe7e5e4, 0xf97316, 0x1c1917)

// SlackPalette — Slack's sidebar dark palette.
var SlackPalette = MakePalette(0x3f0e40, 0xffffff, 0xe8a723, 0x3f0e40)

// TerminalGreenPalette — classic green phosphor terminal.
var TerminalGreenPalette = MakePalette(0x001100, 0x33ff33, 0x00cc00, 0x001100)

// TerminalAmberPalette — amber phosphor terminal.
var TerminalAmberPalette = MakePalette(0x0d0800, 0xffb000, 0xff8c00, 0x0d0800)

// AllPalettes is a slice of all built-in palettes with their names.
// Useful for building a theme picker UI.
//
//	for _, p := range proton.AllPalettes {
//	    if proton.Button(win, &btns[i], p.Name) {
//	        a.ApplyPalette(p.Palette)
//	    }
//	}
var AllPalettes = []NamedPalette{
	{"Dark", DarkPalette},
	{"Nord", NordPalette},
	{"Rose Pine", RosePinePalette},
	{"Rose Pine Moon", RosePineMoonPalette},
	{"Rose Pine Dawn", RosePineDawnPalette},
	{"Catppuccin Mocha", CatppuccinPalette},
	{"Catppuccin Frappé", CatppuccinFrappePalette},
	{"Catppuccin Macchiato", CatppuccinMacchiatoPalette},
	{"Catppuccin Latte", CatppuccinLattePalette},
	{"Dracula", DraculaPalette},
	{"Gruvbox Dark", GruvboxDarkPalette},
	{"Gruvbox Light", GruvboxLightPalette},
	{"Gruvbox Material", GruvboxMaterialDarkPalette},
	{"Tokyo Night", TokyoNightPalette},
	{"Tokyo Night Storm", TokyoNightStormPalette},
	{"Tokyo Night Day", TokyoNightDayPalette},
	{"Monokai", MonokaiPalette},
	{"Solarized Dark", SolarizedDarkPalette},
	{"Solarized Light", SolarizedLightPalette},
	{"One Dark", OneDarkPalette},
	{"Material Dark", MaterialDarkPalette},
	{"Ayu Dark", AyuDarkPalette},
	{"Ayu Mirage", AyuMiragePalette},
	{"Ayu Light", AyuLightPalette},
	{"Everforest Dark", EverforestDarkPalette},
	{"Everforest Light", EverforestLightPalette},
	{"Kanagawa", KanagawaPalette},
	{"Vesper", VesperPalette},
	{"Night Owl", NightOwlPalette},
	{"Carbon", CarbonPalette},
	{"Midnight", MidnightPalette},
	{"Obsidian", ObsidianPalette},
	{"Hacker", HackerPalette},
	{"Cyberpunk", CyberpunkPalette},
	{"Light", LightPalette},
	{"Paper", PaperPalette},
	{"GitHub Light", GithubLightPalette},
	{"Nord Light", NordLightPalette},
	{"Fluent Light", FluentLightPalette},
	{"Oceanic Next", OceanicNextPalette},
	{"Iceberg", IcebergPalette},
	{"Synthwave", SynthwavePalette},
	{"Ole Dark", OleDarkPalette},
	{"Slack", SlackPalette},
	{"Terminal Green", TerminalGreenPalette},
	{"Terminal Amber", TerminalAmberPalette},
}

// NamedPalette pairs a palette with its display name.
type NamedPalette struct {
	Name    string
	Palette Palette
}

// ThemePickerState tracks the picker's scroll and selection.
//
//	type UI struct {
//	    picker proton.ThemePickerState
//	}
//
//	proton.ThemePicker(win, &u.picker, a)
type ThemePickerState struct {
	scroll   Scrollable
	selected int
	btns     []Clickable
}

// ThemePicker draws a scrollable list of all built-in palettes.
// Clicking one applies it immediately to the app.
// Place it in a settings panel or a dedicated theme window.
//
//	proton.ThemePicker(win, &u.picker, a)
func ThemePicker(win Context, state *ThemePickerState, a *App) {
	for len(state.btns) < len(AllPalettes) {
		state.btns = append(state.btns, Clickable{})
	}

	List(win, &state.scroll, len(AllPalettes), func(win Context, i int) {
		p := AllPalettes[i]
		if Tappable(win, &state.btns[i], func(win Context) {
			PadV(win, 6, func(win Context) {
				Row(win,
					func(win Context) {
						// four color swatches
						Row(win,
							func(win Context) { Rect(win, p.Palette.Bg, 14, 14) },
							func(win Context) { Gap(win, 2) },
							func(win Context) { Rect(win, p.Palette.Fg, 14, 14) },
							func(win Context) { Gap(win, 2) },
							func(win Context) { Rect(win, p.Palette.Primary, 14, 14) },
							func(win Context) { Gap(win, 2) },
							func(win Context) { Rect(win, p.Palette.PrimaryFg, 14, 14) },
						)
					},
					func(win Context) { Gap(win, 10) },
					func(win Context) {
						label := p.Name
						if i == state.selected {
							label = "• " + label
						}
						Label(win, label)
					},
				)
			})
		}) {
			state.selected = i
			a.ApplyPalette(p.Palette)
		}
	})
}

// Theme is a chainable builder for custom palettes using CSS hex color codes.
// Get one from a.ThemeBuilder(), set the colors you want, then call Apply().
//
//	a.ThemeBuilder().
//	    Bg("#1e1e2e").
//	    Fg("#cdd6f4").
//	    Primary("#89b4fa").
//	    PrimaryFg("#1e1e2e").
//	    Apply()
//
// Or use ColorCode to patch a single slot on an existing palette:
//
//	a.ApplyPalette(proton.NordPalette)
//	a.ThemeBuilder().Primary("#ff6b6b").Apply()  // override just the primary color
type ThemeBuilder struct {
	app *App
	p   Palette
}

// ThemeBuilder returns a builder pre-loaded with the current palette colors.
func (a *App) ThemeBuilder() *ThemeBuilder {
	return &ThemeBuilder{
		app: a,
		p: Palette{
			Bg:        a.theme.Palette.Bg,
			Fg:        a.theme.Palette.Fg,
			Primary:   a.theme.Palette.ContrastBg,
			PrimaryFg: a.theme.Palette.ContrastFg,
		},
	}
}

// ColorCode parses a CSS hex color string and sets the background color.
// Accepts "#rrggbb", "#rgb", "rrggbb", or "rgb" — with or without the #.
// This is the main entry point for the hex color code API.
//
// Which slot to set is determined by the method you chain it from:
//
//	a.ThemeBuilder().Bg("#1e1e2e").Fg("#cdd6f4").Primary("#89b4fa").Apply()
//
// Or if you want to patch just one color on the current theme:
//
//	a.ThemeBuilder().Primary("#ff0000").Apply()
func (a *App) ColorCode(slot, code string) {
	b := a.ThemeBuilder()
	c := parseHex(code)
	switch slot {
	case "bg", "background":
		b.p.Bg = c
	case "fg", "foreground", "text":
		b.p.Fg = c
	case "primary", "accent":
		b.p.Primary = c
	case "primaryfg", "primarytext":
		b.p.PrimaryFg = c
	}
	b.Apply()
}

// Bg sets the background color from a hex code string.
// Accepts "#rrggbb", "rrggbb", "#rgb", or "rgb".
func (tb *ThemeBuilder) Bg(code string) *ThemeBuilder {
	tb.p.Bg = parseHex(code)
	return tb
}

// Fg sets the foreground (text) color from a hex code string.
func (tb *ThemeBuilder) Fg(code string) *ThemeBuilder {
	tb.p.Fg = parseHex(code)
	return tb
}

// Primary sets the accent/button color from a hex code string.
func (tb *ThemeBuilder) Primary(code string) *ThemeBuilder {
	tb.p.Primary = parseHex(code)
	return tb
}

// PrimaryFg sets the text-on-primary color from a hex code string.
func (tb *ThemeBuilder) PrimaryFg(code string) *ThemeBuilder {
	tb.p.PrimaryFg = parseHex(code)
	return tb
}

// Apply pushes the built palette into the app.
func (tb *ThemeBuilder) Apply() {
	tb.app.ApplyPalette(tb.p)
}

// Palette returns the built Palette without applying it.
// Useful if you want to store it and apply it later.
func (tb *ThemeBuilder) Palette() Palette {
	return tb.p
}

// parseHex converts a CSS hex color string to color.NRGBA.
// Handles "#rrggbb", "rrggbb", "#rgb", "rgb", "#rrggbbaa", "rrggbbaa".
// Returns transparent black on invalid input rather than panicking.
func parseHex(s string) color.NRGBA {
	// strip leading #
	if len(s) > 0 && s[0] == '#' {
		s = s[1:]
	}

	switch len(s) {
	case 3: // rgb shorthand — expand to rrggbb
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
		fallthrough
	case 6: // rrggbb
		v, ok := hexToUint32(s)
		if !ok {
			return color.NRGBA{}
		}
		return color.NRGBA{
			R: uint8(v >> 16),
			G: uint8(v >> 8),
			B: uint8(v),
			A: 0xff,
		}
	case 4: // rgba shorthand
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2], s[3], s[3]})
		fallthrough
	case 8: // rrggbbaa
		v, ok := hexToUint32(s[:6])
		a, aok := hexToUint32("00" + s[6:8])
		if !ok || !aok {
			return color.NRGBA{}
		}
		return color.NRGBA{
			R: uint8(v >> 16),
			G: uint8(v >> 8),
			B: uint8(v),
			A: uint8(a),
		}
	}
	return color.NRGBA{}
}

// hexToUint32 converts up to 8 hex chars to a uint32.
func hexToUint32(s string) (uint32, bool) {
	var v uint32
	for _, c := range s {
		v <<= 4
		switch {
		case c >= '0' && c <= '9':
			v |= uint32(c - '0')
		case c >= 'a' && c <= 'f':
			v |= uint32(c-'a') + 10
		case c >= 'A' && c <= 'F':
			v |= uint32(c-'A') + 10
		default:
			return 0, false
		}
	}
	return v, true
}
