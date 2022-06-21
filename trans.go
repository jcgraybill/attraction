package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var transColors [5]color.RGBA

func init() {
	transColors = [5]color.RGBA{
		{R: 96, G: 207, B: 250, A: 255},
		{R: 244, G: 168, B: 186, A: 255},
		{R: 255, G: 255, B: 255, A: 255},
		{R: 244, G: 168, B: 186, A: 255},
		{R: 96, G: 207, B: 250, A: 255},
	}
}

func generateTransFlag(w, h, segments int) *ebiten.Image {

	if segments > len(transColors) {
		segments = len(transColors)
	}

	flag := ebiten.NewImage(w, h)
	flag.Fill(fg)

	stripe := ebiten.NewImage(w, flag.Bounds().Dy()/len(transColors))
	stripeOpts := &ebiten.DrawImageOptions{}
	for i := 0; i < segments; i++ {
		stripe.Fill(transColors[i])
		flag.DrawImage(stripe, stripeOpts)
		stripeOpts.GeoM.Translate(0, float64(stripe.Bounds().Dy()))
	}

	return flag
}

func getTransLevel(which int) Level {

	transLevels := [3]Level{
		{
			label: "level 1: boys and girls",
			moves: 15,
			cells: 6,
			flag:  generateFlagSplashScreen(2, generateTransFlag),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      0,
					y:      1,
					tile:   true,
					color:  transColors[0],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    transColors[0],
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[2].color) },
					x:      1,
					y:      5,
					tile:   true,
					color:  transColors[1],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        1,
					y:        4,
					moveable: true,
					color:    transColors[1],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        5,
					y:        1,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        3,
					y:        4,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        3,
					y:        1,
					moveable: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      3,
					y:      0,
				},

				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      5,
					y:      0,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      4,
					y:      2,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      0,
					y:      3,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      3,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      3,
					y:      3,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      4,
					y:      4,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      0,
					y:      5,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      5,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      3,
					y:      5,
				},
			},
		},

		{
			label: "level 2: transitioning",
			moves: 15,
			cells: 6,
			flag:  generateFlagSplashScreen(3, generateTransFlag),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      1,
					y:      2,
					tile:   true,
					color:  color.RGBA{R: 200, G: 200, B: 200, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        3,
					y:        2,
					moveable: true,
					color:    color.RGBA{R: 200, G: 200, B: 200, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        3,
					y:        0,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        3,
					y:        3,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        3,
					y:        1,
					moveable: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      1,
					y:      0,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      2,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      4,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      1,
					y:      5,
				},
			},
		},
		{
			label: "level 3: girls and boys",
			moves: 20,
			cells: 4,
			flag:  generateFlagSplashScreen(5, generateTransFlag),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      3,
					y:      3,
					tile:   true,
					color:  transColors[3],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    transColors[3],
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[2].color) },
					x:      0,
					y:      0,
					tile:   true,
					color:  transColors[4],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        2,
					y:        2,
					moveable: true,
					color:    transColors[4],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        1,
					y:        0,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        2,
					y:        3,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      0,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      1,
					y:      3,
				},
			},
		},
	}

	if which > len(transLevels)-1 {
		return last
	}
	return transLevels[which]
}
