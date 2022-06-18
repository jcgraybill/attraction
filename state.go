package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	keypress ebiten.Key
	pieces   []*Piece
}

func pushState(k ebiten.Key) {
	var s State
	s.keypress = k
	s.pieces = make([]*Piece, 0)
	for _, p := range level.pieces {
		s.pieces = append(s.pieces, &Piece{
			x:        p.x,
			y:        p.y,
			moveable: p.moveable,
		})
	}
	state = append(state, &s)
}

func popState(steps int) {
	if steps >= len(state) {
		return
	}

	for i, p := range state[len(state)-1-steps].pieces {
		if level.pieces[i] != nil {
			level.pieces[i].x = p.x
			level.pieces[i].y = p.y
			level.pieces[i].moveable = p.moveable
		}
	}

	if len(state) > 1 { // state[0] is the initial board position
		state = state[:len(state)-steps]
		level.moves += steps
	}

}
