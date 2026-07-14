package proton

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// logoState holds a decoded, GPU-ready logo image cached on the App.
type logoState struct {
	img   ImageOp
	ready bool
}

// SetLogo loads a logo from a file path. Call once at startup.
// Supports PNG and JPEG. The image is decoded and cached — not re-read every frame.
//
//	err := a.SetLogo("assets/logo.png")
//
// With go:embed:
//
//	//go:embed assets/logo.png
//	var logoBytes []byte
//
//	err := a.SetLogoBytes(logoBytes)
func (a *App) SetLogo(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	op, err := toImageOp(img)
	if err != nil {
		return err
	}
	a.logo = &logoState{img: op, ready: true}
	return nil
}

// SetLogoBytes loads a logo from raw bytes in memory.
// Useful with go:embed — embed the file and pass the byte slice here.
func (a *App) SetLogoBytes(data []byte) error {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}
	op, err := toImageOp(img)
	if err != nil {
		return err
	}
	a.logo = &logoState{img: op, ready: true}
	return nil
}

// toImageOp converts any image.Image to a Proton ImageOp.
func toImageOp(img image.Image) (ImageOp, error) {
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		bounds := img.Bounds()
		nrgba = image.NewNRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				nrgba.Set(x, y, img.At(x, y))
			}
		}
	}
	return ImageOp{op: paint.NewImageOp(nrgba), sz: nrgba.Bounds().Size()}, nil
}

// Logo draws the app logo at the given size in dp.
// Pass 0 for either dimension to use the image's natural pixel size.
//
//	proton.Logo(ctx, 48, 48)
//	proton.Logo(ctx, 0, 0) // natural size
func Logo(win Context, widthDp, heightDp float32) {
	logo := win.appLogo()
	if logo == nil || !logo.ready {
		return
	}
	win.add(func(gtx layout.Context) layout.Dimensions {
		w, h := logo.img.sz.X, logo.img.sz.Y
		if widthDp > 0 {
			w = gtx.Dp(unit.Dp(widthDp))
		}
		if heightDp > 0 {
			h = gtx.Dp(unit.Dp(heightDp))
		}
		sz := image.Pt(w, h)
		stack := clip.Rect{Max: sz}.Push(gtx.Ops)
		logo.img.op.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		stack.Pop()
		return layout.Dimensions{Size: sz}
	})
}

// HasLogo returns true if the app has a logo set via SetLogo or SetLogoBytes.
func HasLogo(win Context) bool {
	l := win.appLogo()
	return l != nil && l.ready
}
