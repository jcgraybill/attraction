package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var board *ebiten.Image
var boardOpts *ebiten.DrawImageOptions

func init() {
	board, boardOpts = generateBoardImage()
}

func generateBoardImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	dc := gg.NewContext(w*5/6, w*5/6)
	dc.SetColor(fg)

	for x := 0; x <= 6; x += 1 {
		dc.DrawLine(float64(x*w/6), 0, float64(x*w/6), float64(w*5/6))
		dc.DrawLine(0, float64(x*w/6), float64(w*5/6), float64(x*w/6))
	}

	dc.Stroke()
	bo := &ebiten.DrawImageOptions{}
	bo.GeoM.Translate(float64(w*1/12), float64(w*1/12))
	return ebiten.NewImageFromImage(dc.Image()), bo
}
