package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var nonbinaryColors [4]color.RGBA

func init() {
	nonbinaryColors = [4]color.RGBA{
		{R: 255, G: 248, B: 47, A: 255},
		{R: 255, G: 255, B: 255, A: 255},
		{R: 156, G: 77, B: 210, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
	}
}

func generateNonbinaryFlag(w, h, segments int) *ebiten.Image {

	if segments > len(nonbinaryColors) {
		segments = len(nonbinaryColors)
	}

	flag := ebiten.NewImage(w, h)
	flag.Fill(fg)

	stripe := ebiten.NewImage(w, flag.Bounds().Dy()/len(nonbinaryColors))
	stripeOpts := &ebiten.DrawImageOptions{}
	for i := 0; i < segments; i++ {
		stripe.Fill(nonbinaryColors[i])
		flag.DrawImage(stripe, stripeOpts)
		stripeOpts.GeoM.Translate(0, float64(stripe.Bounds().Dy()))
	}

	return flag
}

func getNonbinaryLevel(which int) Level {

	nonbinaryLevels := [2]Level{
		{
			label: "level 1: separate and multiple",
			moves: 12,
			cells: 5,
			flag:  generateFlagSplashScreen(2, generateNonbinaryFlag),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      4,
					y:      1,
					tile:   true,
					color:  nonbinaryColors[0],
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[1].color) },
					x:      3,
					y:      1,
					tile:   true,
					color:  white, // use custom value, otherwise piece is invisible
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[2].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    nonbinaryColors[0],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        2,
					y:        1,
					moveable: true,
					color:    white,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        1,
					y:        0,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      4,
					y:      2,
				},
			},
		},
		{
			label: "level 2: mix and without",
			moves: 1,
			cells: 3,
			flag:  generateFlagSplashScreen(4, generateNonbinaryFlag),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      2,
					y:      0,
					tile:   true,
					color:  nonbinaryColors[2],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        1,
					y:        0,
					moveable: true,
					color:    nonbinaryColors[2],
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[2].color) },
					x:      2,
					y:      1,
					tile:   true,
					color:  nonbinaryColors[3],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    nonbinaryColors[3],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        0,
					y:        0,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        0,
					y:        1,
					moveable: true,
					magnetic: true,
				},
			},
		},
	}

	if which > len(nonbinaryLevels)-1 {
		return last
	}
	return nonbinaryLevels[which]
}
