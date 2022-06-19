package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const w = 400
const h = 520
const bs = 360

var bg = color.White
var fg = color.Black

var level Level
var state []*State

type Game struct {
	splash *ebiten.Image
	timer  int
	level  int
}

func main() {
	ebiten.SetWindowTitle("m")
	ebiten.SetWindowSize(w, h)
	level = getRainbowLevel(0)
	initializeLevel()

	err := ebiten.RunGame(&Game{
		splash: generateStartingSplashScreen(),
		timer:  240,
		level:  0,
	})
	if err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
