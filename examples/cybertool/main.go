package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/CzaxStudio/proton"
)

type CyberTool struct {
	// Scanner State
	targetInput proton.Editor
	scanBtn     proton.Clickable
	scanning    bool
	progress    float32
	results     []string
	resList     proton.Scrollable

	// Hash State
	hashInput  proton.Editor
	hashResult string

	// Base64 State
	b64Input  proton.Editor
	b64Result proton.Editor
	encBtn    proton.Clickable
	decBtn    proton.Clickable
}

func main() {
	u := &CyberTool{
		results: []string{"Ready to scan..."},
	}

	a := proton.New("Proton CyberTool")
	a.ApplyPalette(proton.NordPalette)
	a.SetBackground(proton.RGB(0x008F11))

	a.Window("Proton CyberTool v1.0", 500, 700, func(win *proton.Win) {
		// Header
		proton.Pad(win, 16, func(win *proton.Win) {
			proton.H4(win, "⚡ Proton CyberTool")
			proton.Caption(win, "Lightweight security utility toolkit")
		})

		proton.Divider(win)

		// 1. Port Scanner Section (Simulated)
		section(win, "Port Scanner", func(win *proton.Win) {
			proton.Input(win, &u.targetInput, "Target (e.g. 127.0.0.1)")
			proton.Gap(win, 8)

			if !u.scanning {
				if proton.Button(win, &u.scanBtn, "Start Port Scan") {
					u.scanning = true
					u.progress = 0
					u.results = []string{"Scanning " + u.targetInput.Text() + "..."}
					go u.runFakeScan(win)
				}
			} else {
				proton.ProgressBar(win, u.progress)
			}

			proton.Gap(win, 12)
			proton.MinSize(win, 0, 150, func(win *proton.Win) {
				// Nord secondary background color
				proton.Card(win, proton.RGB(0x3B4252), 4, 8, func(win *proton.Win) {
					proton.List(win, &u.resList, len(u.results), func(win *proton.Win, i int) {
						proton.Body2(win, u.results[i])
					})
				})
			})
		})

		// 2. Hash Generator (SHA-256)
		section(win, "SHA-256 Hasher", func(win *proton.Win) {
			proton.Input(win, &u.hashInput, "Text to hash...")
			txt := u.hashInput.Text()
			if txt != "" {
				sum := sha256.Sum256([]byte(txt))
				u.hashResult = fmt.Sprintf("%x", sum)
			} else {
				u.hashResult = ""
			}

			if u.hashResult != "" {
				proton.Gap(win, 8)
				proton.Card(win, proton.RGB(0x3B4252), 4, 8, func(win *proton.Win) {
					proton.Text(win, u.hashResult, 12, proton.RGB(0x88C0D0), false) // Nord frost blue
				})
			}
		})

		// 3. Base64 Encoder/Decoder
		section(win, "Base64 Tool", func(win *proton.Win) {
			proton.TextArea(win, &u.b64Input, "Data...")
			proton.Gap(win, 8)

			proton.Row(win, func(win *proton.Win) {
				if proton.Button(win, &u.encBtn, "Encode") {
					u.b64Result.SetText(base64.StdEncoding.EncodeToString([]byte(u.b64Input.Text())))
				}
				proton.Gap(win, 8)
				if proton.Button(win, &u.decBtn, "Decode") {
					data, err := base64.StdEncoding.DecodeString(u.b64Input.Text())
					if err != nil {
						u.b64Result.SetText("Error: " + err.Error())
					} else {
						u.b64Result.SetText(string(data))
					}
				}
			})

			if u.b64Result.Text() != "" {
				proton.Gap(win, 8)
				proton.TextArea(win, &u.b64Result, "Result")
			}
		})

		proton.Gap(win, 20)
	})

	a.Run()
}

func section(win *proton.Win, title string, content func(*proton.Win)) {
	proton.Pad(win, 16, func(win *proton.Win) {
		proton.H6(win, title)
		proton.Gap(win, 8)
		content(win)
	})
}

func (u *CyberTool) runFakeScan(win *proton.Win) {
	ports := []int{21, 22, 53, 80, 443, 3306, 5432, 8080}
	for i, p := range ports {
		time.Sleep(400 * time.Millisecond)
		u.progress = float32(i+1) / float32(len(ports))
		status := "CLOSED"
		if p == 80 || p == 443 || p == 22 {
			status = "OPEN"
		}
		u.results = append(u.results, fmt.Sprintf("Port %d: %s", p, status))
		win.Invalidate()
	}
	u.scanning = false
	u.results = append(u.results, "Scan complete.")
	win.Invalidate()
}
