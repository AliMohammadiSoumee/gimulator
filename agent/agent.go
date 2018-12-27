package agent

import (
	"github.com/alidadar7676/gimulator/types"
)


type gameState struct {
	bestNode *gameState
	benefit  int
	it       *iteration
}

func (gs *gameState) Hit(ben int, child *gameState) {
	gs.benefit = ben
	gs.bestNode = child
}

func (gs *gameState) heuristic() int {
	for _, ws := range gs.it.winStates {
		if ws.Equal(gs.it.ball) {
			return inf
		}
	}

	for _, ls := range gs.it.loseStates {
		if ls.Equal(gs.it.ball) {
			return -inf
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	targetPoint := gs.it.winStates[1]
	a := abs(gs.it.ball.X-targetPoint.X) + abs(gs.it.ball.Y-targetPoint.Y)
	return -a
}

func run(world types.World, name string) types.Move {
	depth := 6
	it := newIteration(world, name)
	root := &gameState{
		it: it,
	}
	root.minimax(depth)

	return types.Move{
		A: root.it.ball,
		B: root.bestNode.it.ball,
	}
}
