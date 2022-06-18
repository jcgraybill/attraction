package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Piece struct {
	image                    *ebiten.Image
	imgsrc                   func() *ebiten.Image
	color                    color.Color
	currentX, currentY, x, y int
	magnetic, moveable, tile bool
	opts                     *ebiten.DrawImageOptions
}

func generateGemImage(c color.RGBA) *ebiten.Image {
	dc := gg.NewContext(bs/level.cells, bs/level.cells)
	dc.DrawCircle(float64(bs/(level.cells*2)), float64(bs/(level.cells*2)), float64((0.9*bs)/(level.cells*2)))
	dc.SetColor(c)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateMagnetImage() *ebiten.Image {
	dc := gg.NewContext(bs/level.cells, bs/level.cells)
	dc.DrawRoundedRectangle(float64(0.1*bs/level.cells), float64(0.1*bs/level.cells), float64((bs/level.cells)-(0.2*bs/level.cells)), float64((bs/level.cells)-(0.2*bs/level.cells)), float64(0.1*bs/level.cells))
	dc.SetRGBA255(128, 128, 128, 255)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateBlockImage() *ebiten.Image {
	dc := gg.NewContext(bs/level.cells, bs/level.cells)
	dc.DrawRoundedRectangle(float64(0.1*bs/level.cells), float64(0.1*bs/level.cells), float64((bs/level.cells)-(0.2*bs/level.cells)), float64((bs/level.cells)-(0.2*bs/level.cells)), float64(0.1*bs/level.cells))
	dc.SetRGBA255(64, 64, 64, 255)
	dc.Fill()
	dc.DrawRoundedRectangle(float64(0.2*bs/level.cells), float64(0.2*bs/level.cells), float64((bs/level.cells)-(0.4*bs/level.cells)), float64((bs/level.cells)-(0.4*bs/level.cells)), float64(0.1*bs/level.cells))
	dc.SetColor(bg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generatePostImage() *ebiten.Image {
	dc := gg.NewContext(bs/level.cells, bs/level.cells)
	dc.DrawRectangle(float64(0.1*bs/level.cells), float64(0.1*bs/level.cells), float64((bs/level.cells)-(0.2*bs/level.cells)), float64((bs/level.cells)-(0.2*bs/level.cells)))
	dc.SetRGBA255(64, 64, 64, 255)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateTargetImage(c color.RGBA) *ebiten.Image {
	dc := gg.NewContext(bs/level.cells, bs/level.cells)
	dc.DrawRectangle(0, 0, float64(bs/level.cells), float64(bs/level.cells))
	dc.SetColor(c)
	dc.Fill()
	dc.DrawCircle(float64(bs/(level.cells*2)), float64(bs/(level.cells*2)), float64((0.9*bs)/(level.cells*2)))
	dc.SetColor(bg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
