package proton

import (
	"image"
	"sync"
	"time"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// ToastState holds notification state. Declare one in your UI struct.
//
//	u.toast.Show("Saved!", 2*time.Second)
//	proton.Toast(win, &u.toast)  // call last so it renders on top
type ToastState struct {
	msg   string
	until time.Time
	mu    sync.Mutex
}

func (t *ToastState) Show(msg string, duration time.Duration) {
	t.mu.Lock()
	t.msg = msg
	t.until = time.Now().Add(duration)
	t.mu.Unlock()
}

func (t *ToastState) active() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	if time.Now().Before(t.until) {
		return t.msg
	}
	return ""
}

// Toast draws a pill notification. Call last in your draw function.
func Toast(win Context, state *ToastState) {
	msg := state.active()
	if msg == "" {
		return
	}
	win.Invalidate()
	win.add(func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				w, h := gtx.Constraints.Min.X, gtx.Constraints.Min.Y
				r := gtx.Dp(unit.Dp(h/2 + 1))
				rrect := clip.RRect{Rect: image.Rect(0, 0, w, h), NW: r, NE: r, SE: r, SW: r}
				paint.FillShape(gtx.Ops, win.theme().Palette.ContrastBg, rrect.Op(gtx.Ops))
				return layout.Dimensions{Size: image.Pt(w, h)}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Body2(win.theme(), msg)
				lbl.Color = win.theme().Palette.ContrastFg
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, lbl.Layout)
			}),
		)
	})
}
