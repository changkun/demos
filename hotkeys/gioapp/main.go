package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"golang.design/x/hotkey"
)

func main() {
	go fn()
	app.Main()
}

func fn() {
	w := app.NewWindow(app.Size(unit.Dp(200), unit.Dp(200)))

	go reghk()

	for range w.Events() {
	}
}

func reghk() {
	// Register a desired hotkey.
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	var err error
	if app.CallOnMainThread(func() { err = hk.Register() }); err != nil {
		panic("hotkey registration failed")
	}

	// Unregister the hotkey when keydown event is triggered
	for range hk.Keydown() {
		hk.Unregister()
	}

	// Exist the app.
	os.Exit(0)
}
