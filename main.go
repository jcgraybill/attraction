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

type Game struct{}

func main() {
	ebiten.SetWindowTitle("m")
	ebiten.SetWindowSize(w, h)

	level = level01
	load()

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
