package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.splashScreenCountdown > 0 {
		screen.DrawImage(g.splashScreenImage, nil)
	} else if g.menu {
		screen.DrawImage(menuImage, nil)
	} else {

		screen.Fill(bg)

		screen.DrawImage(level.board, level.boardOpts)

		for _, p := range level.pieces {
			screen.DrawImage(p.image, p.opts)
		}

		text.Draw(screen, level.label, *regular, (w-bs)/2, (h-bs)/2-10, fg)
		if len(state) == 1 {
			text.Draw(screen, fmt.Sprintf("moves remaining: %d\npress '?' for help", level.moves), *regular, (w-bs)/2, bs+(h-bs)/2+24, fg)
		} else {
			text.Draw(screen, fmt.Sprintf("moves remaining: %d\n'?': help, 'esc': undo", level.moves), *regular, (w-bs)/2, bs+(h-bs)/2+24, fg)
		}
	}

}
