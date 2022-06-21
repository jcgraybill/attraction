package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func generateRightArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(0, 0)
	dc.LineTo(float64(textHeight), float64(textHeight/2))
	dc.LineTo(0, float64(textHeight))
	dc.LineTo(0, 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateLeftArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(float64(textHeight), 0)
	dc.LineTo(float64(textHeight), float64(textHeight))
	dc.LineTo(0, float64(textHeight/2))
	dc.MoveTo(float64(textHeight), 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateDownArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(0, 0)
	dc.LineTo(float64(textHeight), 0)
	dc.LineTo(float64(textHeight/2), float64(textHeight))
	dc.LineTo(0, 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateUpArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(float64(textHeight/2), 0)
	dc.LineTo(float64(textHeight), float64(textHeight))
	dc.LineTo(0, float64(textHeight))
	dc.LineTo(float64(textHeight/2), 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
