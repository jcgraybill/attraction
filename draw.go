package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	screen.DrawImage(board, boardOpts)

	for _, p := range pieces {
		screen.DrawImage(p.img, p.opts)
	}
}
