package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)
	text.Draw(screen, level.label, *regular, (w-bs)/2, (h-bs)/2-10, fg)
	screen.DrawImage(level.board, level.boardOpts)

	for _, p := range level.pieces {
		screen.DrawImage(p.image, p.opts)
	}

}
