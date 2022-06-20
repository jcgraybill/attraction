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
