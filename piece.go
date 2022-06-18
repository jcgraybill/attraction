package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Piece struct {
	img                        *ebiten.Image
	cellX, cellY, destX, destY int
	magnetic, moveable         bool
	opts                       *ebiten.DrawImageOptions
}

var pieces []*Piece

func generateGemImage(c color.RGBA) *ebiten.Image {
	dc := gg.NewContext(bs/cells, bs/cells)
	dc.DrawCircle(bs/(cells*2), bs/(cells*2), (0.9*bs)/(cells*2))
	dc.SetColor(c)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateMagnetImage() *ebiten.Image {
	dc := gg.NewContext(bs/cells, bs/cells)
	dc.DrawRoundedRectangle(float64(0.1*bs/cells), float64(0.1*bs/cells), float64((bs/cells)-(0.2*bs/cells)), float64((bs/cells)-(0.2*bs/cells)), 0.1*bs/cells)
	dc.SetRGBA255(128, 128, 128, 255)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateBlockImage() *ebiten.Image {
	dc := gg.NewContext(bs/cells, bs/cells)
	dc.DrawRectangle(float64(0.1*bs/cells), float64(0.1*bs/cells), float64((bs/cells)-(0.2*bs/cells)), float64((bs/cells)-(0.2*bs/cells)))
	dc.SetRGBA255(64, 64, 64, 255)
	dc.Fill()
	dc.DrawRectangle(float64(0.2*bs/cells), float64(0.2*bs/cells), float64((bs/cells)-(0.4*bs/cells)), float64((bs/cells)-(0.4*bs/cells)))
	dc.SetColor(bg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generatePostImage() *ebiten.Image {
	dc := gg.NewContext(bs/cells, bs/cells)
	dc.DrawRectangle(float64(0.1*bs/cells), float64(0.1*bs/cells), float64((bs/cells)-(0.2*bs/cells)), float64((bs/cells)-(0.2*bs/cells)))
	dc.SetRGBA255(64, 64, 64, 255)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
