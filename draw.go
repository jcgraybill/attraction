package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	screen.DrawImage(level.board, level.boardOpts)

	for _, p := range level.pieces {
		screen.DrawImage(p.image, p.opts)
	}

	text.Draw(screen, level.label, *regular, (w-bs)/2, (h-bs)/2-10, fg)
	text.Draw(screen, fmt.Sprintf("moves remaining: %d\npress 'esc' to undo", level.moves), *regular, (w-bs)/2, bs+(h-bs)/2+24, fg)

}
