# Examples

Working code you can copy, paste, and modify.
These are complete programs — they compile and run.

---

## Hello World

The absolute minimum. Open a window, show some text.

```go
package main

import "github.com/CzaxStudio/proton"

func main() {
    a := proton.New("hello")
    a.Window("Hello", 400, 200, func(win proton.Context) {
        proton.H3(win, "Hello from Proton!")
    })
    a.Run()
}
```

---

## Counter

A number that goes up when you click a button.

```go
package main

import (
    "fmt"
    "github.com/CzaxStudio/proton"
)

type UI struct {
    count int
    inc   proton.Clickable
    dec   proton.Clickable
}

func main() {
    u := &UI{}
    a := proton.New("counter")
    a.Window("Counter", 300, 200, func(win proton.Context) {
        proton.Center(win, func(win proton.Context) {
            proton.H2(win, fmt.Sprintf("%d", u.count))
            proton.Gap(win, 16)
            proton.Row(win,
                func(win proton.Context) {
                    proton.Pad(win, 8, func(win proton.Context) {
                        if proton.OutlineButton(win, &u.dec, "−") {
                            u.count--
                        }
                    })
                },
                func(win proton.Context) { proton.Gap(win, 8) },
                func(win proton.Context) {
                    proton.Pad(win, 8, func(win proton.Context) {
                        if proton.Button(win, &u.inc, "+") {
                            u.count++
                        }
                    })
                },
            )
        })
    })
    a.Run()
}
```

---

## Todo List

The required demo for every UI framework in existence.

```go
package main

import "github.com/CzaxStudio/proton"

type item struct {
    text string
    done proton.Bool
}

type UI struct {
    input  proton.Editor
    add    proton.Clickable
    items  []item
    scroll proton.Scrollable
}

func main() {
    u := &UI{}
    a := proton.New("todo")
    a.ApplyPalette(proton.NordPalette)
    a.Window("Todo", 420, 560, func(win proton.Context) {
        draw(win, u)
    })
    a.Run()
}

func draw(win proton.Context, u *UI) {
    proton.H4(win, "Todo")
    proton.Gap(win, 12)

    // input row
    proton.GrowRow(win,
        proton.GrowItem(win, func(win proton.Context) {
            proton.Input(win, &u.input, "What needs doing?")
        }),
        proton.FixedItem(win, func(win proton.Context) {
            proton.Gap(win, 8)
        }),
        proton.FixedItem(win, func(win proton.Context) {
            proton.Pad(win, 4, func(win proton.Context) {
                if proton.Button(win, &u.add, "Add") {
                    t := u.input.Text()
                    if t != "" {
                        u.items = append(u.items, item{text: t})
                        u.input.SetText("")
                    }
                }
            })
        }),
    )

    proton.Gap(win, 8)
    proton.Divider(win)
    proton.Gap(win, 8)

    // item list
    if len(u.items) == 0 {
        proton.Center(win, func(win proton.Context) {
            proton.Caption(win, "Nothing to do. Suspicious.")
        })
        return
    }

    proton.List(win, &u.scroll, len(u.items), func(win proton.Context, i int) {
        t := &u.items[i]
        proton.PadV(win, 6, func(win proton.Context) {
            proton.Checkbox(win, &t.done, t.text)
        })
    })
}
```

---

## Login Form

A form with validation feedback.

```go
package main

import (
    "github.com/CzaxStudio/proton"
    "strings"
)

type UI struct {
    email    proton.Editor
    password proton.Editor
    submit   proton.Clickable
    errMsg   string
}

func main() {
    u := &UI{}
    a := proton.New("login")
    a.ApplyPalette(proton.DarkPalette)
    a.Window("Sign In", 400, 320, func(win proton.Context) {
        draw(win, u)
    })
    a.Run()
}

func draw(win proton.Context, u *UI) {
    proton.Center(win, func(win proton.Context) {
        proton.MaxWidth(win, 320, func(win proton.Context) {
            proton.H4(win, "Sign In")
            proton.Gap(win, 24)

            proton.Label(win, "Email")
            proton.Gap(win, 4)
            proton.Pad(win, 4, func(win proton.Context) {
                proton.Input(win, &u.email, "you@example.com")
            })
            proton.Gap(win, 12)

            proton.Label(win, "Password")
            proton.Gap(win, 4)
            proton.Pad(win, 4, func(win proton.Context) {
                proton.Input(win, &u.password, "••••••••")
            })
            proton.Gap(win, 8)

            if u.errMsg != "" {
                proton.Text(win, u.errMsg, 13, proton.RGB(0xf38ba8), false)
                proton.Gap(win, 8)
            }

            proton.Pad(win, 4, func(win proton.Context) {
                if proton.Button(win, &u.submit, "Sign In") {
                    if err := validate(u); err != "" {
                        u.errMsg = err
                    } else {
                        u.errMsg = ""
                        handleLogin(u.email.Text(), u.password.Text())
                    }
                }
            })
        })
    })
}

func validate(u *UI) string {
    if !strings.Contains(u.email.Text(), "@") {
        return "Please enter a valid email address."
    }
    if len(u.password.Text()) < 8 {
        return "Password must be at least 8 characters."
    }
    return ""
}

func handleLogin(email, password string) {
    println("Logging in:", email)
}
```

---

## Settings Panel

Toggles, radio buttons, and a slider in a clean settings layout.

```go
package main

import (
    "fmt"
    "github.com/CzaxStudio/proton"
)

type Settings struct {
    notifications proton.Bool
    darkMode      proton.Bool
    language      proton.Enum
    fontSize      proton.Float
    save          proton.Clickable
}

func main() {
    s := &Settings{}
    s.fontSize.Value = 0.5   // default to middle of range

    a := proton.New("settings")
    a.ApplyPalette(proton.CatppuccinPalette)
    a.Window("Settings", 480, 500, func(win proton.Context) {
        drawSettings(win, s)
    })
    a.Run()
}

func drawSettings(win proton.Context, s *Settings) {
    proton.H4(win, "Settings")
    proton.Gap(win, 20)

    // appearance section
    proton.H6(win, "Appearance")
    proton.Gap(win, 8)
    proton.Pad(win, 4, func(win proton.Context) {
        proton.Toggle(win, &s.darkMode, "Dark mode")
    })
    proton.Gap(win, 8)
    proton.Label(win, fmt.Sprintf("Font size: %.0f%%", 80+s.fontSize.Value*40))
    proton.Gap(win, 4)
    proton.Slider(win, &s.fontSize)
    proton.Gap(win, 20)

    // notifications section
    proton.H6(win, "Notifications")
    proton.Gap(win, 8)
    proton.Pad(win, 4, func(win proton.Context) {
        proton.Checkbox(win, &s.notifications, "Email notifications")
    })
    proton.Gap(win, 20)

    // language section
    proton.H6(win, "Language")
    proton.Gap(win, 8)
    proton.Pad(win, 4, func(win proton.Context) {
        proton.RadioButton(win, &s.language, "en", "English")
        proton.Gap(win, 4)
        proton.RadioButton(win, &s.language, "es", "Español")
        proton.Gap(win, 4)
        proton.RadioButton(win, &s.language, "fr", "Français")
    })
    proton.Gap(win, 24)

    // save button
    proton.RowEnd(win,
        func(win proton.Context) {
            proton.Pad(win, 4, func(win proton.Context) {
                if proton.Button(win, &s.save, "Save Settings") {
                    println("saved")
                }
            })
        },
    )
}
```

---

## Dashboard with Cards

A two-column layout with stat cards.

```go
package main

import (
    "fmt"
    "github.com/CzaxStudio/proton"
)

type Stat struct {
    label string
    value string
    color uint32
}

func main() {
    stats := []Stat{
        {"Users", "1,204", 0x5e81ac},
        {"Revenue", "$4,820", 0xa3be8c},
        {"Errors", "3", 0xbf616a},
        {"Uptime", "99.9%", 0x88c0d0},
    }

    a := proton.New("dashboard")
    a.ApplyPalette(proton.NordPalette)
    a.Window("Dashboard", 640, 400, func(win proton.Context) {
        proton.H4(win, "Dashboard")
        proton.Gap(win, 16)

        proton.Grid(win, 2, 12,
            statCards(stats)...,
        )
    })
    a.Run()
}

func statCards(stats []Stat) []func(proton.Context) {
    fns := make([]func(proton.Context), len(stats))
    for i, s := range stats {
        s := s
        fns[i] = func(win proton.Context) {
            proton.Card(win, proton.RGB(0x3b4252), 10, 16, func(win proton.Context) {
                proton.Badge(win, proton.RGB(s.color), proton.RGB(0x2e3440), s.label)
                proton.Gap(win, 8)
                proton.H4(win, s.value)
            })
        }
    }
    return fns
}

// helper — Grid takes []func(Context) not ...func(Context) directly here
// so we unpack: proton.Grid(win, 2, 12, statCards(stats)...)
```

---

## Where to Go From Here

- Read the individual doc pages for functions you want to use
- Look at the examples in the `examples/` folder of the repo
- When something isn't working, check [04-layout.md](./04-layout.md) —
  most issues come down to missing layout wrappers around interactive widgets
