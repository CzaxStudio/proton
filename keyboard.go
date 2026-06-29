package proton

import (
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
)

// Modifier represents a keyboard modifier key (Ctrl, Shift, Alt).
// Combine with bitwise OR: proton.ModCtrl | proton.ModShift.
type Modifier int

const (
	ModNone  Modifier = 0
	ModCtrl  Modifier = 1 << 0
	ModShift Modifier = 1 << 1
	ModAlt   Modifier = 1 << 2
)

func (m Modifier) toGio() key.Modifiers {
	var g key.Modifiers
	if m&ModCtrl != 0 {
		g |= key.ModCtrl
	}
	if m&ModShift != 0 {
		g |= key.ModShift
	}
	if m&ModAlt != 0 {
		g |= key.ModAlt
	}
	return g
}

// Common key names for use with OnKey. Letter keys are just their
// uppercase string, e.g. "S", "Z", "N" — these constants cover the rest.
const (
	KeyEscape    = string(key.NameEscape)
	KeyReturn    = string(key.NameReturn)
	KeyBackspace = string(key.NameDeleteBackward)
	KeyDelete    = string(key.NameDeleteForward)
	KeyTab       = string(key.NameTab)
	KeySpace     = " "
	KeyUp        = string(key.NameUpArrow)
	KeyDown      = string(key.NameDownArrow)
	KeyLeft      = string(key.NameLeftArrow)
	KeyRight     = string(key.NameRightArrow)
)

// OnKey fires fn when the given key is pressed, with the given modifiers held.
//
//	proton.OnKey(win, proton.ModCtrl, "S", func() { save() })
//	proton.OnKey(win, proton.ModNone, proton.KeyEscape, func() { closeDialog() })
//	proton.OnKey(win, proton.ModCtrl|proton.ModShift, "N", func() { newWindow() })
func OnKey(win Context, modifiers Modifier, keyName string, fn func()) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		name := key.Name(keyName)
		tag := &struct{ key.Name }{name}
		event.Op(gtx.Ops, tag)
		for {
			ev, ok := gtx.Source.Event(key.Filter{Name: name, Required: modifiers.toGio()})
			if !ok {
				break
			}
			if e, ok := ev.(key.Event); ok && e.State == key.Press {
				fn()
			}
		}
		return layout.Dimensions{}
	})
}

// FrameTag is a stable identity token for registering a focusable region.
// Declare one per focusable area in your state struct.
//
//	type UI struct {
//	    editorTag proton.FrameTag
//	}
type FrameTag struct{}

// FocusArea registers a region of the UI as a key event receiver.
// keyName works like OnKey — pass a letter or one of the Key* constants.
//
//	proton.FocusArea(win, &ui.editorTag, "A", func(win proton.Context) {
//	    proton.TextArea(win, &ui.text, "Type here...")
//	})
func FocusArea(win Context, tag *FrameTag, keyName string, content func(Context)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		event.Op(gtx.Ops, tag)
		for {
			_, ok := gtx.Source.Event(key.Filter{Name: key.Name(keyName)})
			if !ok {
				break
			}
		}
		return child(win, content)(gtx)
	})
}
