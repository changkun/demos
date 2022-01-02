// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() { mainthread.Init(fn) }
func fn() {
	hk1 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyS)
	go func() {
		println("register")
		var err error
		if mainthread.Call(func() { err = hk1.Register() }); err != nil {
			panic(err)
		}
		for range hk1.Keydown() {
			println("hotkey ctrl+s is triggered")
		}
	}()

	<-time.After(5 * time.Second)
	hk1.Unregister()
	println("unregistered")

	<-time.After(5 * time.Second)
	go func() {
		println("register again")
		var err error
		if mainthread.Call(func() { err = hk1.Register() }); err != nil {
			panic(err)
		}
		for range hk1.Keydown() {
			println("hotkey ctrl+s is triggered")
		}
	}()

	<-time.After(5 * time.Second)
	hk1.Unregister()
	println("unregistered again")

	println("done")
}
