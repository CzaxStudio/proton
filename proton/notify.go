package proton

// Notify is a stub for OS-level notifications.
// Gio's notification support lives in gioui.org/x/notify which has
// platform-specific setup requirements (Android manifest entries, macOS
// entitlements, etc.). Rather than pulling that dependency in by default,
// Proton exposes this no-op stub so your code compiles on all platforms.
//
// To send real OS notifications, add gioui.org/x/notify to your go.mod
// and call its API directly from your app code.
//
//	// go get gioui.org/x/notify
//	// then in your app:
//	manager, _ := notify.NewManager()
//	manager.CreateNotification("Title", "Body")
func (a *App) Notify(title, body string) {
	// no-op stub — see package comment above
}
