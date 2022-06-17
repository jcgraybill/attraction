package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if handleKeyPress() {
		updateLocations()
	}
	return nil
}

func handleKeyPress() bool {
	redraw := false
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		for _, p := range pieces {
			if p.magnetic {
				move(p, 1, 0)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		for _, p := range pieces {
			if p.magnetic {
				move(p, -1, 0)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		for _, p := range pieces {
			if p.magnetic {
				move(p, 0, 1)
				redraw = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		for _, p := range pieces {
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

	if p.destX = p.cellX + x; p.destX < 0 || p.destX >= cells {
		p.destX = p.cellX
		return false
	}

	if p.destY = p.cellY + y; p.destY < 0 || p.destY >= cells {
		p.destY = p.cellY
		return false
	}

	for _, q := range pieces {
		if p != q {
			if p.destX == q.destX && p.destY == q.destY {
				if move(q, x, y) {
					return true
				} else {
					p.destX = p.cellX
					p.destY = p.cellY
					return false
				}
			}
		}
	}
	return true
}

func updateLocations() {
	for _, p := range pieces {
		p.cellX = p.destX
		p.cellY = p.destY
		p.opts.GeoM.Reset()
		p.opts.GeoM.Translate(float64((w-bs)/2+p.cellX*bs/cells), float64((w-bs)/2+p.cellY*bs/cells))
	}
}
