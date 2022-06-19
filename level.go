package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	label     string
	final     bool
	cells     int
	moves     int
	pieces    []*Piece
	board     *ebiten.Image
	boardOpts *ebiten.DrawImageOptions

	flag *ebiten.Image

	next *Level
}

func initializeLevel() {
	level.board, level.boardOpts = generateBoardImage(level.cells)

	for _, p := range level.pieces {
		p.image = p.imgsrc()
		p.opts = &ebiten.DrawImageOptions{}
	}

	updatePieceLocations()
	state = make([]*State, 0)
	pushState(0)
}

var last = Level{
	label: "thanks for playing",
	cells: 2,
	moves: 0,
	final: true,
	flag:  ebiten.NewImage(w, h),
	pieces: []*Piece{
		{
			imgsrc: func() *ebiten.Image { return generateGemImage(rainbowColors[0]) },
			x:      0,
			y:      0,
		},
		{
			imgsrc: func() *ebiten.Image { return generateGemImage(rainbowColors[1]) },
			x:      1,
			y:      0,
		},
		{
			imgsrc: func() *ebiten.Image { return generateGemImage(rainbowColors[4]) },
			x:      0,
			y:      1,
		},
		{
			imgsrc: func() *ebiten.Image { return generateGemImage(rainbowColors[5]) },
			x:      1,
			y:      1,
		},
	},
}
