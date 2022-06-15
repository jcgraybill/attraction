package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

const w = 400

func main() {
	ebiten.SetWindowSize(w, w)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(int, int) (int, int) {
	return w, w
}

func (g *Game) Update() error {
	return nil
}
