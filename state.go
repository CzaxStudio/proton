package proton

import "gioui.org/widget"

// Re-exported state types so simple programs only need to import proton.
//
// These are deliberate type aliases, not wrapped types — unlike Context,
// which fully hides Gio behind an interface. State types are stable,
// low-churn structs (Gio's widget.Clickable etc. haven't broken in years),
// and aliasing them lets you write "var btn proton.Clickable" with a normal
// zero value, no constructor needed. The API-immunity guarantee applies to
// everything you pass a Context into — widget functions, layout helpers —
// which is where Gio's actual churn happens (layout/event internals).

// Clickable tracks clicks on a button or tappable area.
type Clickable = widget.Clickable

// Editor holds state for a text input or textarea.
type Editor = widget.Editor

// Bool holds the checked state of a checkbox.
type Bool = widget.Bool

// Enum holds the selected key in a radio group.
type Enum = widget.Enum

// Float holds a slider value 0.0–1.0.
type Float = widget.Float

// Scrollable tracks scroll position for List and HList.
type Scrollable = widget.List

// Drag tracks drag gesture state.
type Drag = widget.Draggable

// Icon holds vector icon data for IconButton.
type Icon = widget.Icon

// NumberState (number.go), ResizeSplitState (scrollable_split.go),
// TabState / AccordionState / SpinnerState / SelectBoxState / OverlayState
// (extra.go), and AlertKind (alert.go) are Proton-native types — not Gio
// aliases — declared in their respective files.
