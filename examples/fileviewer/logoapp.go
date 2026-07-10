package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/CzaxStudio/proton"
)

// Embed a PNG at compile time so the binary is fully self-contained.
// Drop any PNG/JPEG into the same folder as main.go, rename it logo.png,
// and this line will include it in the binary automatically.
//
//go:embed logo.png
var logoBytes []byte

type ui struct {
	count  int
	btn    proton.Clickable
	outBtn proton.Clickable
	name   proton.Editor
	toast  proton.ToastState
}

var gApp *proton.App

func main() {
	u := &ui{}

	gApp = proton.New("logoapp")
	gApp.ApplyPalette(proton.CarbonPalette)

	// Load the embedded logo once. If it fails (e.g. logo.png is missing),
	// the app still runs — Logo() just draws nothing.
	if err := gApp.SetLogoBytes(logoBytes); err != nil {
		// in production you'd log this, not panic
		println("logo load failed:", err.Error())
	}

	gApp.Window("LogoIntroduced!", 480, 520, func(ctx proton.Context) {
		draw(ctx, u)
	})
	gApp.Run()
}

func draw(ctx proton.Context, u *ui) {
	// centered layout
	proton.Center(ctx, func(ctx proton.Context) {
		proton.MaxWidth(ctx, 340, func(ctx proton.Context) {

			// logo + app name side by side
			proton.Row(ctx,
				func(ctx proton.Context) {
					// draws the logo at 52×52 dp
					// does nothing if SetLogoBytes failed
					proton.Logo(ctx, 52, 52)
				},
				func(ctx proton.Context) { proton.Gap(ctx, 14) },
				func(ctx proton.Context) {
					proton.Column(ctx,
						func(ctx proton.Context) { proton.H4(ctx, "Click it") },
						func(ctx proton.Context) { proton.Gap(ctx, 2) },
						func(ctx proton.Context) { proton.Muted(ctx, "Built with Proton") },
					)
				},
			)

			proton.Gap(ctx, 32)
			proton.Divider(ctx)
			proton.Gap(ctx, 24)

			// simple counter
			proton.H2(ctx, fmt.Sprintf("%d", u.count))
			proton.Gap(ctx, 4)
			proton.Muted(ctx, "clicks so far")
			proton.Gap(ctx, 16)

			proton.Row(ctx,
				func(ctx proton.Context) {
					proton.Pad(ctx, 4, func(ctx proton.Context) {
						if proton.Button(ctx, &u.btn, "Click me") {
							u.count++
							if u.count%10 == 0 {
								u.toast.Show(fmt.Sprintf("%d clicks — nice.", u.count), 2*time.Second)
							}
						}
					})
				},
				func(ctx proton.Context) { proton.Gap(ctx, 8) },
				func(ctx proton.Context) {
					proton.Pad(ctx, 4, func(ctx proton.Context) {
						if proton.OutlineButton(ctx, &u.outBtn, "Reset") {
							u.count = 0
							u.toast.Show("Reset.", time.Second)
						}
					})
				},
			)

			proton.Gap(ctx, 24)
			proton.Divider(ctx)
			proton.Gap(ctx, 16)

			proton.Input(ctx, &u.name, "Your name")
			proton.Gap(ctx, 6)
			if u.name.Text() != "" {
				proton.Label(ctx, "Hello, "+u.name.Text()+"!")
			}
		})
	})

	proton.Toast(ctx, &u.toast)
}
