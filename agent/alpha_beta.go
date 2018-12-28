package agent

import (
	"log"
)

const (
	inf = int(1e6)
)

func (gs *gameState) alphabeta(depth int) int {
	return gs.max(depth, -inf, inf)
}

func (gs *gameState) max(depth, alpha, beta int) int {
	heur := gs.heuristic(true)
	if depth == 0 || heur <= -inf || heur >= inf {
		log.Println(gs.ball, "nil", heur)
		gs.hit(heur, nil)
		return heur
	}

	value := -3 * inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.max(depth-1, alpha, beta)
		} else {
			mm = child.min(depth-1, alpha, beta)
		}
		if value < mm {
			value = mm
			bestChild = child
		}
		if alpha < value {
			alpha = value
		}
		if alpha >= beta {
			gs.hit(value, bestChild)
			gs.it.prev(mv)
			return inf / 2
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic(true)
	}
	if bestChild != nil {
		log.Println(gs.ball, bestChild.ball, value)
	} else {
		log.Println(gs.ball, "nil", value)
	}
	gs.hit(value, bestChild)
	return value
}

func (gs *gameState) min(depth, alpha, beta int) int {
	heur := gs.heuristic(false)
	if depth == 0 || heur <= -inf || heur >= inf {
		log.Println(gs.ball, "nil", heur)
		gs.hit(heur, nil)
		return heur
	}

	value := 3 * inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.min(depth-1, alpha, beta)
		} else {
			mm = child.max(depth-1, alpha, beta)
		}
		if value > mm {
			value = mm
			bestChild = child
		}
		if beta > value {
			beta = value
		}
		if alpha >= beta {
			gs.hit(value, bestChild)
			gs.it.prev(mv)
			return -inf / 2
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic(false)
	}
	if bestChild != nil {
		log.Println(gs.ball, bestChild.ball, value)
	} else {
		log.Println(gs.ball, "nil", value)
	}
	gs.hit(value, bestChild)
	return value
}
