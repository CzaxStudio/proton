package proton

import "gioui.org/widget"

// Re-exported state types so simple programs only need to import proton.

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
