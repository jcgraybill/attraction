package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if movePiece() {
		block.opts.GeoM.Reset()
		block.opts.GeoM.Translate(float64(w*1/12+(block.cellX*w*1/6)), float64(w*1/12+(block.cellY*w*1/6)))
		gem.opts.GeoM.Reset()
		gem.opts.GeoM.Translate(float64(w*1/12+(gem.cellX*w*1/6)), float64(w*1/12+(gem.cellY*w*1/6)))
	}
	return nil
}

func movePiece() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && block.cellX < 4 {
		block.cellX += 1
		if gem.cellX == block.cellX && gem.cellY == block.cellY {
			if gem.cellX >= 4 {
				block.cellX -= 1
				return false
			} else {
				gem.cellX += 1
			}
		}
		return true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && block.cellX > 0 {
		block.cellX -= 1
		if gem.cellX == block.cellX && gem.cellY == block.cellY {
			if gem.cellX <= 0 {
				block.cellX += 1
				return false
			} else {
				gem.cellX -= 1
			}
		}
		return true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && block.cellY > 0 {
		block.cellY -= 1
		if gem.cellX == block.cellX && gem.cellY == block.cellY {
			if gem.cellY <= 0 {
				block.cellY += 1
				return false
			} else {
				gem.cellY -= 1
			}
		}
		return true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && block.cellY < 4 {
		block.cellY += 1
		if gem.cellX == block.cellX && gem.cellY == block.cellY {
			if gem.cellY >= 4 {
				block.cellY -= 1
				return false
			} else {
				gem.cellY += 1
			}
		}
		return true
	}
	return false
}
