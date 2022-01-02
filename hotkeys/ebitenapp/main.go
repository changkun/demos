package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.design/x/hotkey"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")

	go reghk()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func reghk() {
	// Register a desired hotkey.
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	if err := ebiten.RunOnMainThread(func() error { return hk.Register() }); err != nil {
		panic("hotkey registration failed")
	}

	// Unregister the hotkey when keydown event is triggered
	for range hk.Keydown() {
		hk.Unregister()
	}

	// Exist the app.
	os.Exit(0)
}
