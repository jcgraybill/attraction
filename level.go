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
	label: "level 1: life",
	moves: 2,
	cells: 4,
	next:  &level02,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
			x:      2,
			y:      1,
			tile:   true,
			color:  color.RGBA{R: 228, G: 3, B: 3, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 228, G: 3, B: 3, A: 255},
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
	label: "level 2: healing",
	moves: 6,
	cells: 4,
	next:  &level03,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
			x:      2,
			y:      1,
			tile:   true,
			color:  color.RGBA{R: 255, G: 140, B: 0, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 255, G: 140, B: 0, A: 255},
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
	label: "level 3: sunlight and nature",
	moves: 6,
	cells: 4,
	next:  &level04,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
			x:      1,
			y:      0,
			tile:   true,
			color:  color.RGBA{R: 255, G: 237, B: 3, A: 255},
		},
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[1].color) },
			x:      0,
			y:      0,
			tile:   true,
			color:  color.RGBA{R: 0, G: 128, B: 38, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[2].color) },
			x:        1,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 255, G: 237, B: 3, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
			x:        0,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 0, G: 128, B: 38, A: 255},
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

var level04 = Level{
	label: "level 4: magic and spirit",
	moves: 11,
	cells: 4,
	next:  &last,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
			x:      3,
			y:      2,
			tile:   true,
			color:  color.RGBA{R: 0, G: 77, B: 255, A: 255},
		},
		{
			imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[1].color) },
			x:      2,
			y:      1,
			tile:   true,
			color:  color.RGBA{R: 117, G: 7, B: 135, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[2].color) },
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 0, G: 77, B: 255, A: 255},
		},
		{
			imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
			x:        1,
			y:        1,
			moveable: true,
			color:    color.RGBA{R: 117, G: 7, B: 135, A: 255},
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
			x:        2,
			y:        3,
			moveable: true,
			magnetic: true,
		},
		{
			imgsrc:   func() *ebiten.Image { return generateBlockImage() },
			x:        2,
			y:        2,
			moveable: true,
		},
		{
			imgsrc:   func() *ebiten.Image { return generateBlockImage() },
			x:        0,
			y:        1,
			moveable: true,
		},
		{
			imgsrc:   func() *ebiten.Image { return generateBlockImage() },
			x:        3,
			y:        1,
			moveable: true,
		},
		{
			imgsrc: func() *ebiten.Image { return generatePostImage() },
			x:      2,
			y:      0,
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
