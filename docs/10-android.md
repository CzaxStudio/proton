# Android

Proton apps run on Android through Gio's native Android support.
The same code that runs on your desktop runs on Android — no rewrites,
no separate UI layer.

---

## What works

Every Proton widget, layout, and theme works on Android exactly as it
does on desktop. Touch events map to pointer events. The soft keyboard
integrates with `Input` and `TextArea`. `Invalidate()` works correctly
on mobile. The one restriction from Gio's side: Android supports only
one window per app.

---

## Install the build tool

You need `gogio`, Gio's cross-compilation tool:

```bash
go install gioui.org/cmd/gogio@latest
```

You also need the Android SDK and NDK. The easiest path is Android Studio:
[developer.android.com/studio](https://developer.android.com/studio)

Set the environment variables so gogio can find the SDK:

```bash
export ANDROID_HOME=$HOME/Android/Sdk
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/<version>
```

---

## Build an APK

From your project root:

```bash
gogio -target android -appid com.yourname.yourapp .
```

This produces `yourapp.apk`. Install it on a connected device:

```bash
adb install yourapp.apk
```

Or run directly on a connected device:

```bash
gogio -target android -appid com.yourname.yourapp -run .
```

---

## Build an AAR (embed in an existing Android project)

```bash
gogio -target android -buildmode archive -appid com.yourname.yourapp .
```

Then include the `.aar` in your Android project's `libs/` folder and
add it to `build.gradle`:

```groovy
dependencies {
    implementation fileTree(dir: 'libs', include: ['*.aar'])
}
```

Declare the Gio activity in `AndroidManifest.xml`:

```xml
<activity
    android:name="org.gioui.GioActivity"
    android:theme="@style/Theme.GioApp"
    android:configChanges="orientation|keyboardHidden"
    android:windowSoftInputMode="adjustResize">
</activity>
```

---

## Android 15+ (16kB page sizes)

Google Play requires 16kB page-size compatible APKs from November 2025.
`gogio` handles this automatically — just keep it up to date:

```bash
go install gioui.org/cmd/gogio@latest
```

---

## Logo / app icon

Set your app icon through the standard Android mechanism (in your
`AndroidManifest.xml` via `android:icon`), or use Proton's logo feature
to draw it inside the UI itself:

```go
//go:embed assets/logo.png
var logoBytes []byte

func main() {
    a := proton.New("myapp")
    a.SetLogoBytes(logoBytes) // load once

    a.Window("My App", 480, 800, func(ctx proton.Context) {
        proton.Logo(ctx, 64, 64) // draw it in the layout
        proton.H4(ctx, "My App")
    })
    a.Run()
}
```

---

## Minimum SDK

Gio supports Android SDK 16+ (Android 4.1, Jelly Bean). In practice,
anything below SDK 21 (Android 5.0) is under 1% of active devices,
so targeting 21 is a reasonable minimum.

Set it in your `AndroidManifest.xml`:

```xml
<uses-sdk
    android:minSdkVersion="21"
    android:targetSdkVersion="35" />
```

---

## A full example

```go
package main

import "github.com/CzaxStudio/proton"

type UI struct {
    btn proton.Clickable
    count int
}

func main() {
    u := &UI{}
    a := proton.New("counter")
    a.ApplyPalette(proton.NordPalette)
    a.Window("Counter", 480, 800, func(ctx proton.Context) {
        proton.Center(ctx, func(ctx proton.Context) {
            proton.H2(ctx, fmt.Sprintf("%d", u.count))
            proton.Gap(ctx, 24)
            proton.Pad(ctx, 8, func(ctx proton.Context) {
                if proton.Button(ctx, &u.btn, "Tap me") {
                    u.count++
                }
            })
        })
    })
    a.Run()
}
```

Build for Android:

```bash
gogio -target android -appid com.example.counter .
adb install counter.apk
```

That's it. The same binary you run with `go run .` on your laptop becomes
an Android APK with one command.
