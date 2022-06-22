package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var progressiveColors [11]color.RGBA

func init() {
	progressiveColors = [11]color.RGBA{
		{R: 212, G: 29, B: 6, A: 255},
		{R: 238, G: 156, B: 0, A: 255},
		{R: 255, G: 255, B: 11, A: 255},
		{R: 5, G: 191, B: 0, A: 255},
		{R: 1, G: 26, B: 152, A: 255},
		{R: 118, G: 38, B: 127, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 96, G: 56, B: 19, A: 255},
		{R: 96, G: 207, B: 250, A: 255},
		{R: 244, G: 168, B: 186, A: 255},
		{R: 255, G: 255, B: 255, A: 255},
	}
}

func generateProgressiveFlag(w, h, segments int) *ebiten.Image {
	flag := ebiten.NewImage(w, h)
	flag.Fill(fg)

	stripe := ebiten.NewImage(w, flag.Bounds().Dy()/6)
	stripeOpts := &ebiten.DrawImageOptions{}
	for i := 0; i < 6; i++ {
		stripe.Fill(progressiveColors[i])
		flag.DrawImage(stripe, stripeOpts)
		stripeOpts.GeoM.Translate(0, float64(stripe.Bounds().Dy()))
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(-h/2))

	for i := 6; i < 11; i++ {
		dc := gg.NewContext(h, h*2)

		dc.MoveTo(0, 0)
		dc.LineTo(float64(h), float64(h))
		dc.LineTo(0, float64(h*2))
		dc.LineTo(0, 0)
		dc.SetColor(progressiveColors[i])
		dc.Fill()
		opts.GeoM.Translate(float64(-1*flag.Bounds().Dy()/6), 0)
		flag.DrawImage(ebiten.NewImageFromImage(dc.Image()), opts)
	}

	return flag
}

func getProgressiveLevel(which int) Level {

	return Level{
		label: "progressive",
		moves: 2,
		cells: 4,
		final: true,
		flag:  generateFlagSplashScreen(0, generateProgressiveFlag),
		pieces: []*Piece{
			{
				imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
				x:      2,
				y:      1,
				tile:   true,
				color:  rainbowColors[0],
			},
			{
				imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
				x:        1,
				y:        2,
				moveable: true,
				color:    rainbowColors[0],
			},
			{
				imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
				x:        0,
				y:        2,
				moveable: true,
				magnetic: true,
			},
			{
				imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
				x:        1,
				y:        3,
				moveable: true,
				magnetic: true,
			},
		},
	}
}
