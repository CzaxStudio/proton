package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/CzaxStudio/proton"
)

type ui struct {
	target   proton.Editor
	scanBtn  proton.Clickable
	pingBtn  proton.Clickable
	portBtn  proton.Clickable
	clearBtn proton.Clickable
	logs     []string
	scroll   proton.Scrollable
	mu       sync.Mutex
}

func main() {
	rand.Seed(time.Now().UnixNano())

	u := &ui{}

	a := proton.New("CyberTool")
	// local palette with pure black background
	blackPalette := proton.MakePalette(0x000000, 0xffffff, 0xd1ff00, 0x000000)
	a.ApplyPalette(blackPalette)
	a.SetFontScale(1.05)
	a.Window("CyberTool — Cybersecurity Demo", 760, 480, func(win *proton.Win) {
		draw(win, u)
	})
	a.Run()
}

func draw(win *proton.Win, u *ui) {
	proton.Column(win,
		func(win *proton.Win) {
			proton.Row(win,
				func(win *proton.Win) {
					proton.H3(win, "CyberTool — Live Demo")
				},
				func(win *proton.Win) { proton.Gap(win, 8) },
			)
		},
		func(win *proton.Win) {
			proton.Row(win,
				func(win *proton.Win) {
					proton.Input(win, &u.target, "target (ip or hostname)")
				},
				func(win *proton.Win) {
					if proton.Button(win, &u.scanBtn, "Quick Scan") {
						appendLog(win, u, "Starting quick scan")
						simulateQuickScan(win, u)
					}
				},
				func(win *proton.Win) {
					if proton.Button(win, &u.pingBtn, "Ping") {
						appendLog(win, u, "Pinging target")
						simulatePing(win, u)
					}
				},
				func(win *proton.Win) {
					if proton.Button(win, &u.portBtn, "Port Scan") {
						appendLog(win, u, "Starting port scan")
						simulatePortScan(win, u)
					}
				},
				func(win *proton.Win) {
					if proton.Button(win, &u.clearBtn, "Clear") {
						u.mu.Lock()
						u.logs = nil
						u.mu.Unlock()
					}
				},
			)
		},
		func(win *proton.Win) {
			proton.Divider(win)
		},
		func(win *proton.Win) {
			proton.List(win, &u.scroll, len(u.logs), func(win *proton.Win, i int) {
				u.mu.Lock()
				txt := u.logs[i]
				u.mu.Unlock()
				proton.Label(win, txt)
			})
		},
	)
}

func appendLog(win *proton.Win, u *ui, s string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.logs = append(u.logs, fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), s))
	win.Invalidate()
}

func simulateQuickScan(win *proton.Win, u *ui) {
	go func() {
		time.Sleep(time.Millisecond * 300)
		appendLog(win, u, "Found open endpoint: /status")
		time.Sleep(time.Millisecond * 500)
		appendLog(win, u, "Found weak header: Server: tiny-http")
		time.Sleep(time.Millisecond * 400)
		appendLog(win, u, "Quick scan complete")
	}()
}

func simulatePing(win *proton.Win, u *ui) {
	go func() {
		for i := 0; i < 4; i++ {
			delay := rand.Intn(100) + 20
			time.Sleep(time.Millisecond * time.Duration(delay))
			appendLog(win, u, fmt.Sprintf("icmp_seq=%d time=%dms", i+1, delay))
		}
		appendLog(win, u, "ping complete")
	}()
}

func simulatePortScan(win *proton.Win, u *ui) {
	go func() {
		ports := []int{22, 80, 443, 8080, 3306}
		for _, p := range ports {
			time.Sleep(time.Millisecond * 350)
			state := "closed"
			if rand.Intn(2) == 0 {
				state = "open"
			}
			appendLog(win, u, fmt.Sprintf("port %d: %s", p, state))
		}
		appendLog(win, u, "port scan complete")
	}()
}
