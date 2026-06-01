package proton

import (
	"image/color"

	"gioui.org/unit"
)

// Palette holds the colors you want to customize.
// Zero fields fall back to material defaults.
type Palette struct {
	Bg        color.NRGBA
	Fg        color.NRGBA
	Primary   color.NRGBA
	PrimaryFg color.NRGBA
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

func (a *App) SetBackground(c color.NRGBA) {
	a.theme.Palette.Bg = c
}

// SetFontScale multiplies the base text size. 1.0 is default.
func (a *App) SetFontScale(scale float32) {
	a.theme.TextSize = unit.Sp(float32(a.theme.TextSize) * scale)
}

var DarkPalette = Palette{
	Bg:        color.NRGBA{R: 18, G: 18, B: 18, A: 255},
	Fg:        color.NRGBA{R: 220, G: 220, B: 220, A: 255},
	Primary:   color.NRGBA{R: 100, G: 149, B: 237, A: 255},
	PrimaryFg: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
}

var NordPalette = Palette{
	Bg:        color.NRGBA{R: 46, G: 52, B: 64, A: 255},
	Fg:        color.NRGBA{R: 216, G: 222, B: 233, A: 255},
	Primary:   color.NRGBA{R: 136, G: 192, B: 208, A: 255},
	PrimaryFg: color.NRGBA{R: 46, G: 52, B: 64, A: 255},
}

var RosePinePalette = Palette{
	Bg:        color.NRGBA{R: 25, G: 23, B: 36, A: 255},
	Fg:        color.NRGBA{R: 224, G: 222, B: 244, A: 255},
	Primary:   color.NRGBA{R: 196, G: 167, B: 231, A: 255},
	PrimaryFg: color.NRGBA{R: 25, G: 23, B: 36, A: 255},
}

var CatppuccinPalette = Palette{
	Bg:        color.NRGBA{R: 30, G: 30, B: 46, A: 255},
	Fg:        color.NRGBA{R: 205, G: 214, B: 244, A: 255},
	Primary:   color.NRGBA{R: 137, G: 180, B: 250, A: 255},
	PrimaryFg: color.NRGBA{R: 30, G: 30, B: 46, A: 255},
}
