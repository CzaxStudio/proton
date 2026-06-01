package proton

import (
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
)

// OnKey fires fn when the given key+modifiers are pressed.
//
//	proton.OnKey(win, key.ModCtrl, "S", func() { save() })
//	proton.OnKey(win, 0, key.NameEscape, func() { closeDialog() })
func OnKey(win *Win, modifiers key.Modifiers, name key.Name, fn func()) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		tag := &struct{ key.Name }{name}
		event.Op(gtx.Ops, tag)
		for {
			ev, ok := gtx.Source.Event(key.Filter{Name: name, Required: modifiers})
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

// FrameTag is a stable pointer for use as a Gio event tag.
type FrameTag struct{}

// FocusArea registers a UI region as a key event receiver.
func FocusArea(win *Win, tag event.Tag, filter key.Filter, content func(*Win)) {
	win.add(func(gtx layout.Context) layout.Dimensions {
		event.Op(gtx.Ops, tag)
		for {
			_, ok := gtx.Source.Event(filter)
			if !ok {
				break
			}
		}
		return child(win, content)(gtx)
	})
}
