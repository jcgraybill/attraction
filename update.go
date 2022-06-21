package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {

	if g.splashScreenCountdown > 0 {
		g.splashScreenCountdown -= 1

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.splashScreenCountdown = 0
		}

	} else if g.menu {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && menuSelected < len(flags)-1 {
			menuSelected += 1
			menuImage = generateMenuImage()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && menuSelected > 0 {
			menuSelected -= 1
			menuImage = generateMenuImage()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && !flags[menuSelected].completed {
			g.menu = false
			g.level = 0
			level = flags[menuSelected].levelGenerator(g.level)
			initializeLevel()
		}
		return nil
	} else if handleKeyPress(g) {
		updatePieceLocations()
		checkForPiecesOnTarget()
		checkForLevelComplete(g)
	}
	return nil
}

func handleKeyPress(g *Game) bool {
	redraw := false
	var keypress ebiten.Key

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		popState(1)
		return true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		g.splashScreenImage = helpSplash
		g.splashScreenCountdown = 1200
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

	// TODO: if no pieces have actually moved as a result
	// of the keypress, don't count it as a move.

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
	g.splashScreenImage = level.flag
	g.splashScreenCountdown = 300
	g.level += 1
	level = flags[menuSelected].levelGenerator(g.level)
	if level.final {
		g.menu = true
		flags[menuSelected].completed = true
		menuImage = generateMenuImage()
		return
	}
	initializeLevel()
}
