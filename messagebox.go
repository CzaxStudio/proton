package proton

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// DialogResult represents the button clicked in a MessageBox.
type DialogResult int

const (
	MsgNone DialogResult = iota
	MsgOk
	MsgCancel
)

// MessageBoxState holds state for a modal dialog.
type MessageBoxState struct {
	title   string
	msg     string
	visible bool
	result  DialogResult

	btnOk     Clickable
	btnCancel Clickable
}

// Show activates the message box with the given title and message.
func (m *MessageBoxState) Show(title, msg string) {
	m.title = title
	m.msg = msg
	m.visible = true
	m.result = MsgNone
}

// MessageBox draws a modal dialog on top of the UI.
// Call it last in your draw function.
func MessageBox(win Context, state *MessageBoxState) DialogResult {
	if !state.visible {
		return state.result
	}

	if clickResults[&state.btnOk] {
		state.visible = false
		state.result = MsgOk
	}
	if clickResults[&state.btnCancel] {
		state.visible = false
		state.result = MsgCancel
	}

	win.add(func(gtx layout.Context) layout.Dimensions {
		// Scrim / Background overlay
		paint.FillShape(gtx.Ops, RGBA(0, 0, 0, 150), clip.Rect{Max: gtx.Constraints.Max}.Op())

		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return child(win, func(win Context) {
				Card(win, win.theme().Palette.Bg, 8, 16, func(win Context) {
					MinSize(win, 300, 0, func(win Context) {
						H6(win, state.title)
						Gap(win, 8)
						Body2(win, state.msg)
						Gap(win, 16)
						RowEnd(win, func(win Context) {
							OutlineButton(win, &state.btnCancel, "Cancel")
							Gap(win, 8)
							Button(win, &state.btnOk, "OK")
						})
					})
				})
			})(gtx)
		})
	})

	return state.result
}
