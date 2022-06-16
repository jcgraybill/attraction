package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Piece struct {
	img   *ebiten.Image
	cellX int
	cellY int
	opts  *ebiten.DrawImageOptions
}

var gem Piece
var block Piece

func init() {
	gem.img = generateGemImage()
	gem.cellX, gem.cellY = 2, 3
	gem.opts = &ebiten.DrawImageOptions{}
	gem.opts.GeoM.Translate(float64(w*1/12+(gem.cellX*w*1/6)), float64(w*1/12+(gem.cellY*w*1/6)))

	block.img = generateBlockImage()
	block.cellX, block.cellY = 3, 3
	block.opts = &ebiten.DrawImageOptions{}
	block.opts.GeoM.Translate(float64(w*1/12+(block.cellX*w*1/6)), float64(w*1/12+(block.cellY*w*1/6)))

}

func generateGemImage() *ebiten.Image {
	dc := gg.NewContext(w*1/6, w*1/6)
	dc.DrawCircle(w*1/12, w*1/12, 0.9*w*1/12)
	dc.SetRGBA255(255, 0, 0, 128)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateBlockImage() *ebiten.Image {
	dc := gg.NewContext(w*1/6, w*1/6)
	dc.DrawRoundedRectangle(float64(0.1*w*1/6), float64(0.1*w*1/6), float64((w*1/6)-(0.2*w*1/6)), float64((w*1/6)-(0.2*w*1/6)), 10)
	dc.SetRGBA255(128, 128, 128, 255)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
