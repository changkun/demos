// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"golang.design/x/hotkey"
)

func main() {
	w := app.New().NewWindow("Hello")
	label := widget.NewLabel("Hello Fyne!")
	button := widget.NewButton("Hi!", func() { label.SetText("Welcome :)") })
	w.SetContent(container.NewVBox(label, button))

	// Use the provided CallOnMainThread if possible.
	ap, ok := fyne.CurrentApp().(interface{ CallOnMainThread(func()) })
	if !ok {
		panic("fyne: current driver does not support call on the main thread")
	}

	// Register a desired hotkey.
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	var err error
	if ap.CallOnMainThread(func() { err = hk.Register() }); err != nil {
		panic("hotkey registration failed")
	}

	// Start listen hotkey event whenever it is ready.
	go func() {
		for range hk.Keydown() {
			button.Tapped(&fyne.PointEvent{})
		}
	}()

	w.ShowAndRun()
}
