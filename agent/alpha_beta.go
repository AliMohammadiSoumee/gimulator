package agent


const (
	inf = int(1e6)
)


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
		hash := gs.it.hash()

		var mm int
		child, ok := gs.it.hashTable[hash]
		if ok {
			panic("RIDI")
			mm = child.benefit
		} else {
			child = &gameState{it: gs.it, ball: gs.it.ball}
			hasPrice := gs.it.hasPrice()
			if hasPrice {
				mm = child.max(depth-1)
			} else {
				mm = child.min(depth-1)
			}
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
		hash := gs.it.hash()

		var mm int
		child, ok := gs.it.hashTable[hash]
		if ok {
			mm = child.benefit
		} else {
			child = &gameState{it: gs.it, ball: gs.it.ball}
			hasPrice := gs.it.hasPrice()
			if hasPrice {
				mm = child.min(depth-1)
			} else {
				mm = child.max(depth-1)
			}
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
