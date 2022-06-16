package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	screen.DrawImage(board, boardOpts)
	screen.DrawImage(gem.img, gem.opts)
	screen.DrawImage(block.img, block.opts)
}
