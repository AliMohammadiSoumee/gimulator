package agent

const (
	inf = int(1e6)
)

func (gs *gameState) alphabeta(depth int) int {
	return gs.max(depth, -inf, inf)
}

func (gs *gameState) max(depth, alpha, beta int) int {
	heur := gs.heuristic()
	if depth == 0 || heur <= -inf || heur >= inf {
		gs.hit(heur, nil)
		return heur
	}

	value := -2 * inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.max(depth-1, -inf, inf)
		} else {
			mm = child.min(depth-1, alpha, beta)
		}
		if value < mm {
			value = mm
		}
		if alpha < value {
			alpha = value
			bestChild = child
		}
		if alpha > beta {
			gs.hit(value, bestChild)
			gs.it.prev(mv)
			return 2 * inf
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic()
	}
	gs.hit(value, bestChild)
	return value
}

func (gs *gameState) min(depth, alpha, beta int) int {
	if depth == 0 {
		h := gs.heuristic()
		gs.hit(h, nil)
		return h
	}

	value := 2 * inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.min(depth-1, -inf, inf)
		} else {
			mm = child.max(depth-1, alpha, beta)
		}
		if value > mm {
			value = mm
		}
		if beta > value {
			beta = value
			bestChild = child
		}
		if alpha > beta {
			gs.it.prev(mv)
			gs.hit(value, bestChild)
			return -2 * inf
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic()
	}
	gs.hit(value, bestChild)
	return value
}

/*
func (gs *gameState) minimax(depth int) int {
	return gs.max(depth)
}

func (gs *gameState) max(depth int) int {
	heur := gs.heuristic()
	if depth == 0 || heur <= -inf || heur >= inf {
		gs.hit(heur, nil)
		return heur
	}

	value := -inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.max(depth-1)
		} else {
			mm = child.min(depth-1)
		}
		if value < mm {
			value = mm
			bestChild = child
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic()
	}
	gs.hit(value, bestChild)
	return value
}

func (gs *gameState) min(depth int) int {
	if depth == 0 {
		h := gs.heuristic()
		gs.hit(h, nil)
		return h
	}

	value := inf
	var bestChild *gameState

	validMoves := gs.it.validMoves()

	for _, mv := range validMoves {
		gs.it.next(mv)

		var mm int
		child := &gameState{it: gs.it, ball: gs.it.ball}
		hasPrice := gs.it.hasPrice()
		if hasPrice {
			mm = child.min(depth-1)
		} else {
			mm = child.max(depth-1)
		}
		if value > mm {
			value = mm
			bestChild = child
		}
		gs.it.prev(mv)
	}

	if bestChild == nil {
		//TODO
		value = gs.heuristic()
	}
	gs.hit(value, bestChild)
	return value
}
*/
