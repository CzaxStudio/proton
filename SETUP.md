# Setup

After downloading or cloning, run this once from the project root:

```bash
go mod tidy
```

This downloads Gio's transitive dependencies (`gioui.org/shader`,
`golang.org/x/exp`, `golang.org/x/image`, `golang.org/x/sys`,
`golang.org/x/text`, `github.com/go-text/typesetting`, etc.) and writes
them into `go.sum`. These are dependencies Gio itself needs — `go.mod`
only has to declare `gioui.org` directly, and `go mod tidy` resolves
the rest automatically.

If VS Code's Go extension still shows red squiggles after `go mod tidy`,
reload the window (`Ctrl+Shift+P` → "Developer: Reload Window") so
gopls picks up the new go.sum.

## Linux only

```bash
sudo apt install libwayland-dev libxkbcommon-dev libvulkan-dev
```

## Then run any example

```bash
go run ./examples/hello
go run ./examples/kitchen
```
