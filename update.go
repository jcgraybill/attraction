package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if handleKeyPress() {
		updateLocations()
		checkForPiecesOnTarget()
		checkForLevelComplete()
	}
	return nil
}

func handleKeyPress() bool {
	redraw := false
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		for _, p := range level.pieces {
			if p.magnetic {
				move(p, 1, 0)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		for _, p := range level.pieces {
			if p.magnetic {
				move(p, -1, 0)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		for _, p := range level.pieces {
			if p.magnetic {
				move(p, 0, 1)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		for _, p := range level.pieces {
			if p.magnetic {
				move(p, 0, -1)
				redraw = true
			}
		}
	}

	return redraw

}

func move(p *Piece, x, y int) bool {
	if !p.moveable {
		return false
	}

	if p.x = p.currentX + x; p.x < 0 || p.x >= level.cells {
		p.x = p.currentX
		return false
	}

	if p.y = p.currentY + y; p.y < 0 || p.y >= level.cells {
		p.y = p.currentY
		return false
	}

	for _, q := range level.pieces {
		if p != q {
			if !q.tile {
				if p.x == q.x && p.y == q.y {
					if move(q, x, y) {
						return true
					} else {
						p.x = p.currentX
						p.y = p.currentY
						return false
					}
				}
			}
		}
	}
	return true
}

func updateLocations() {
	for _, p := range level.pieces {
		p.currentX = p.x
		p.currentY = p.y
		p.opts.GeoM.Reset()
		p.opts.GeoM.Translate(float64((w-bs)/2+p.currentX*bs/level.cells), float64((h-bs)/2+p.currentY*bs/level.cells))
	}
}

func checkForPiecesOnTarget() {
	for _, p := range level.pieces {
		for _, q := range level.pieces {
			if p != q && p.currentX == q.currentX && p.currentY == q.currentY && p.color == q.color {
				p.moveable = false
			}
		}
	}
}

func checkForLevelComplete() {
	for _, p := range level.pieces {
		if p.tile && p.color != nil { // For now, this means it's a target
			for _, q := range level.pieces {
				if p != q && p.color == q.color { // it's the matching gem
					if p.currentX != q.currentX || p.currentY != q.currentY {
						return // At least one target doesn't have the matching gem on it
					}
				}
			}
		}
	}
	if level.next != nil {
		level = *level.next
		load()
	}
}
