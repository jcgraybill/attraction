package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	pieces = make([]*Piece, 5)

	pieces[0] = &Piece{}
	pieces[0].img = generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128})
	pieces[0].destX, pieces[0].destY = 2, 3
	pieces[0].opts = &ebiten.DrawImageOptions{}
	pieces[0].moveable = true

	pieces[1] = &Piece{}
	pieces[1].img = generateMagnetImage()
	pieces[1].destX, pieces[1].destY = 3, 3
	pieces[1].opts = &ebiten.DrawImageOptions{}
	pieces[1].magnetic = true
	pieces[1].moveable = true

	pieces[2] = &Piece{}
	pieces[2].img = generateMagnetImage()
	pieces[2].destX, pieces[2].destY = 4, 1
	pieces[2].opts = &ebiten.DrawImageOptions{}
	pieces[2].magnetic = true
	pieces[2].moveable = true

	pieces[3] = &Piece{}
	pieces[3].img = generateGemImage(color.RGBA{R: 0, G: 255, B: 0, A: 128})
	pieces[3].destX, pieces[3].destY = 1, 4
	pieces[3].opts = &ebiten.DrawImageOptions{}
	pieces[3].moveable = true

	pieces[4] = &Piece{}
	pieces[4].img = generateBlockImage()
	pieces[4].destX, pieces[4].destY = 1, 3
	pieces[4].opts = &ebiten.DrawImageOptions{}

	updateLocations()
}
