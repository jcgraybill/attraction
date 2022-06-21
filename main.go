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
var white = color.RGBA{R: 200, G: 200, B: 200, A: 255}

var level Level
var state []*State
var menuImage *ebiten.Image

type Game struct {
	splashScreenImage     *ebiten.Image
	splashScreenCountdown int
	level                 int
	menu                  bool
}

func main() {
	ebiten.SetWindowTitle("m")
	ebiten.SetWindowSize(w, h)
	menuImage = generateMenuImage()

	// we're going to use the functions that generate
	// pieces for the help screen, and that requires a
	// level in order to know how big to make them
	level = Level{cells: 12}
	helpSplash = generateHelpSplashScreen()

	err := ebiten.RunGame(&Game{
		splashScreenImage:     generateStartingSplashScreen(),
		splashScreenCountdown: 240,
		menu:                  true,
	})
	if err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
