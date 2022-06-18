package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func generateBoardImage(cells int) (*ebiten.Image, *ebiten.DrawImageOptions) {
	dc := gg.NewContext(bs, bs)
	dc.SetColor(fg)

	for x := 0; x <= cells+1; x += 1 {
		dc.DrawLine(float64(x*bs/cells), 0, float64(x*bs/cells), float64(bs))
		dc.DrawLine(0, float64(x*bs/cells), float64(bs), float64(x*bs/cells))
	}

	dc.Stroke()
	bo := &ebiten.DrawImageOptions{}
	bo.GeoM.Translate(float64((w-bs)/2), float64((w-bs)/2))
	return ebiten.NewImageFromImage(dc.Image()), bo
}
