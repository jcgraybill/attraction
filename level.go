package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	label     string
	cells     int
	moves     int
	pieces    []*Piece
	board     *ebiten.Image
	boardOpts *ebiten.DrawImageOptions

	next *Level
}

func load() {
	level.board, level.boardOpts = generateBoardImage(level.cells)

	for _, p := range level.pieces {
		p.image = p.imgsrc()
		p.opts = &ebiten.DrawImageOptions{}
	}

	updateLocations()
	state = make([]*State, 0)
	pushState(0)
}

var level01 = Level{
	label: "level 1",
	moves: 2,
	cells: 4,
	next:  &level02,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:      2,
			y:      1,
			tile:   true,
			color:  color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
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

var level02 = Level{
	label: "level 2",
	moves: 6,
	cells: 4,
	next:  &level03,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:      2,
			y:      1,
			tile:   true,
			color:  color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
			x:        0,
			y:        2,
			moveable: true,
			magnetic: true,
		},
		{
			imgsrc: func() *ebiten.Image { return generatePostImage() },
			x:      2,
			y:      2,
		},
	},
}

var level03 = Level{
	label: "level 3",
	moves: 7,
	cells: 4,
	next:  &last,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:      1,
			y:      0,
			tile:   true,
			color:  color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(color.RGBA{R: 0, G: 255, B: 0, A: 128}) },
			x:      0,
			y:      0,
			tile:   true,
			color:  color.RGBA{R: 0, G: 255, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:        1,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 0, G: 255, B: 0, A: 128}) },
			x:        0,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 0, G: 255, B: 0, A: 128},
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
			y:        1,
			moveable: true,
			magnetic: true,
		},
		{
			imgsrc:   func() *ebiten.Image { return generateBlockImage() },
			x:        2,
			y:        2,
			moveable: true,
		},
	},
}

var last = Level{
	label: "thanks for playing",
	cells: 2,
	moves: 0,
	next:  nil,
	pieces: []*Piece{
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128}) },
			x:        0,
			y:        0,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 0, G: 255, B: 0, A: 128}) },
			x:        1,
			y:        0,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 0, G: 0, B: 255, A: 128}) },
			x:        0,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(color.RGBA{R: 255, G: 0, B: 255, A: 128}) },
			x:        1,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
	},
}
