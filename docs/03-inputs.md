# Inputs

Text fields, checkboxes, toggles, radio buttons, and sliders.
Everything the user types or toggles.

---

## Text Input

A single-line text field. The `hint` is the placeholder text shown when
the field is empty.

```go
type UI struct {
    name proton.Editor
}

proton.Input(win, &u.name, "Your name")

// read the current value any time
fmt.Println(u.name.Text())
```

**Signature:**
```go
proton.Input(win Context, state *proton.Editor, hint string)
```

The state type is `proton.Editor` — declare one per text field.

---

## Text Area

Like Input but multi-line. The user can press Enter to add new lines.
Good for messages, descriptions, or any longer text.

```go
type UI struct {
    bio proton.Editor
}

proton.TextArea(win, &u.bio, "Write a short bio...")

fmt.Println(u.bio.Text())   // same as Input
```

**Signature:**
```go
proton.TextArea(win Context, state *proton.Editor, hint string)
```

---

## Checkbox

A checkbox with a label. Returns `true` on the frame the user changes it.

```go
type UI struct {
    agreed proton.Bool
}

if proton.Checkbox(win, &u.agreed, "I agree to the terms and conditions") {
    // user just toggled it — u.agreed.Value is the new state
    fmt.Println("now checked:", u.agreed.Value)
}

// read the current value without caring about change events
if u.agreed.Value {
    proton.Label(win, "Thanks for agreeing (nobody read it)")
}
```

**Signature:**
```go
proton.Checkbox(win Context, state *proton.Bool, label string) bool
```

The state type is `proton.Bool`. `.Value` is `true` when checked.

---

## Toggle (Switch)

A material-style on/off switch. Same API as Checkbox, different look.
Use it when you're toggling a setting that takes effect immediately.

```go
type UI struct {
    darkMode proton.Bool
}

if proton.Toggle(win, &u.darkMode, "Dark mode") {
    if u.darkMode.Value {
        applyDarkTheme()
    } else {
        applyLightTheme()
    }
}
```

**Signature:**
```go
proton.Toggle(win Context, state *proton.Bool, label string) bool
```

---

## RadioButton

Radio buttons for when the user needs to pick exactly one option from a group.
All buttons in the same group share one `proton.Enum` state field.

```go
type UI struct {
    size proton.Enum   // one field for the whole group
}

proton.RadioButton(win, &u.size, "sm", "Small")
proton.RadioButton(win, &u.size, "md", "Medium")
proton.RadioButton(win, &u.size, "lg", "Large")

// read the selected key
fmt.Println("selected:", u.size.Value)   // "sm", "md", or "lg"
```

The second argument is the `key` — what gets stored in `.Value` when that
option is selected. The third is the label shown next to the button.

**Signature:**
```go
proton.RadioButton(win Context, group *proton.Enum, key string, label string) bool
```

Returns `true` on the frame the selection changes.

### Horizontal radio buttons

Wrap them in `Row` to lay them out side by side:

```go
proton.Row(win,
    func(win proton.Context) { proton.RadioButton(win, &u.size, "sm", "S") },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) { proton.RadioButton(win, &u.size, "md", "M") },
    func(win proton.Context) { proton.Gap(win, 8) },
    func(win proton.Context) { proton.RadioButton(win, &u.size, "lg", "L") },
)
```

---

## Slider

A horizontal drag handle for choosing a value in a range.
The value is always between `0.0` and `1.0` — scale it yourself.

```go
type UI struct {
    vol proton.Float
}

v := proton.Slider(win, &u.vol)

// v is 0.0–1.0, scale to whatever you need
volume := int(v * 100)
proton.Caption(win, fmt.Sprintf("Volume: %d%%", volume))
```

You can also read the value directly from the state:

```go
proton.Slider(win, &u.vol)
fmt.Println(u.vol.Value)   // 0.0 to 1.0
```

**Signature:**
```go
proton.Slider(win Context, state *proton.Float) float32
```

The state type is `proton.Float`. Returns the current value.

---

## ProgressBar

Not interactive — just shows progress. Pass a `float32` between `0.0` and `1.0`.

```go
proton.ProgressBar(win, 0.65)    // 65% done
proton.ProgressBar(win, 1.0)     // done!
proton.ProgressBar(win, progress) // from a variable
```

**Signature:**
```go
proton.ProgressBar(win Context, progress float32)
```

---

## Full Form Example

Here's a realistic form with most of these widgets:

```go
type SettingsUI struct {
    username  proton.Editor
    bio       proton.Editor
    notify    proton.Bool
    darkMode  proton.Bool
    plan      proton.Enum
    volume    proton.Float
    save      proton.Clickable
}

func drawSettings(win proton.Context, s *SettingsUI) {
    proton.H4(win, "Settings")
    proton.Gap(win, 16)

    proton.Label(win, "Username")
    proton.Gap(win, 4)
    proton.Input(win, &s.username, "your_username")
    proton.Gap(win, 12)

    proton.Label(win, "Bio")
    proton.Gap(win, 4)
    proton.TextArea(win, &s.bio, "Tell us something...")
    proton.Gap(win, 16)

    proton.Toggle(win, &s.darkMode, "Dark mode")
    proton.Gap(win, 8)
    proton.Checkbox(win, &s.notify, "Email notifications")
    proton.Gap(win, 16)

    proton.Label(win, "Plan")
    proton.Gap(win, 4)
    proton.RadioButton(win, &s.plan, "free", "Free")
    proton.RadioButton(win, &s.plan, "pro", "Pro ($9/mo)")
    proton.RadioButton(win, &s.plan, "team", "Team ($29/mo)")
    proton.Gap(win, 16)

    proton.Label(win, fmt.Sprintf("Volume: %.0f%%", s.volume.Value*100))
    proton.Gap(win, 4)
    proton.Slider(win, &s.volume)
    proton.Gap(win, 24)

    proton.Pad(win, 0, func(win proton.Context) {
        if proton.Button(win, &s.save, "Save Settings") {
            handleSave(s)
        }
    })
}
```
