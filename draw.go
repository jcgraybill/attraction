package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	screen.DrawImage(level.board, level.boardOpts)

	for _, p := range level.pieces {
		screen.DrawImage(p.image, p.opts)
	}
}
