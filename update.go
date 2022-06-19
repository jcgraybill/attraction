package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {

	if g.timer > 0 {
		g.timer -= 1

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.timer = 0
		}

	} else if handleKeyPress() {
		updatePieceLocations()
		checkForPiecesOnTarget()
		checkForLevelComplete(g)
	}
	return nil
}

func handleKeyPress() bool {
	redraw := false
	var keypress ebiten.Key

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		popState(1)
		return true
	}

	if level.moves == 0 {
		return false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		for _, p := range level.pieces {
			if p.magnetic {
				movePiece(p, 1, 0)
				keypress = ebiten.KeyArrowRight
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		for _, p := range level.pieces {
			if p.magnetic {
				movePiece(p, -1, 0)
				keypress = ebiten.KeyArrowLeft
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		for _, p := range level.pieces {
			if p.magnetic {
				movePiece(p, 0, 1)
				keypress = ebiten.KeyArrowDown
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		for _, p := range level.pieces {
			if p.magnetic {
				movePiece(p, 0, -1)
				keypress = ebiten.KeyArrowUp
				redraw = true
			}
		}
	}

	if redraw {
		pushState(keypress)
		level.moves -= 1
	}
	return redraw

}

func movePiece(p *Piece, x, y int) bool {
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
					if movePiece(q, x, y) {
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

func updatePieceLocations() {
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

func checkForLevelComplete(g *Game) {

	// TODO: remove when I end with a splash screen
	if level.final {
		return
	}

	for _, p := range level.pieces {
		if p.tile { // For now, this means it's a target
			for _, q := range level.pieces {
				if p != q && p.color == q.color { // it's the matching gem
					if p.currentX != q.currentX || p.currentY != q.currentY {
						return // At least one target doesn't have the matching gem on it
					}
				}
			}
		}
	}
	g.splash = level.flag
	g.timer = 300
	g.level += 1
	level = getRainbowLevel(g.level)
	initializeLevel()
}
