package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var level Level

type Level struct {
	cells     int
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

}

var level01 = Level{
	cells: 4,
	next:  &level02,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image {
				return generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 128})
			},
			x:     2,
			y:     1,
			tile:  true,
			color: color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc: func() *ebiten.Image {
				return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128})
			},
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc: func() *ebiten.Image {
				return generateMagnetImage()
			},
			x:        0,
			y:        2,
			moveable: true,
			magnetic: true,
		},
		{
			imgsrc: func() *ebiten.Image {
				return generateMagnetImage()
			},
			x:        1,
			y:        3,
			moveable: true,
			magnetic: true,
		},
	},
}

var level02 = Level{
	cells: 5,
	next:  nil,
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image {
				return generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 128})
			},
			x:     2,
			y:     1,
			tile:  true,
			color: color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc: func() *ebiten.Image {
				return generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 128})
			},
			x:        1,
			y:        2,
			moveable: true,
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 128},
		},
		{
			imgsrc: func() *ebiten.Image {
				return generateMagnetImage()
			},
			x:        0,
			y:        2,
			moveable: true,
			magnetic: true,
		},
		{
			imgsrc: func() *ebiten.Image {
				return generatePostImage()
			},
			x: 2,
			y: 2,
		},
	},
}
