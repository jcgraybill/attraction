package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func main() {
	ebiten.SetWindowTitle("m")
	ebiten.SetWindowSize(w, w)

	level = level01
	load()

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, w
}
