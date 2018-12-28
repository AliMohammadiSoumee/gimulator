package agent

import (
	"github.com/alidadar7676/gimulator/types"
)


type gameState struct {
	bestChild *gameState
	benefit   int
	it        *iteration
	ball      types.State
}

func (gs *gameState) hit(ben int, child *gameState) {
	gs.benefit = ben
	gs.bestChild = child
}

func (gs *gameState) heuristic(isMax bool) int {
	if gs.it.isBlockingState() {
		if isMax {
			return -inf
		}
		return inf
	}

	if gs.it.isWinState() {
		return inf
	}

	if gs.it.isLoseState() {
		return -inf
	}

	a := gs.it.disWin[gs.ball.X][gs.ball.Y]
	//b := gs.it.disLose[gs.ball.X][gs.ball.Y]

	return -a * a + gs.it.moveNum

	dis := gs.it.distanceFromWinStates()
	return -dis
}



func run(world types.World, name string) types.Move {
	depth := 11
	it := newIteration(world, name)
	root := &gameState{
		it: it,
		ball: world.BallPos,
	}
	root.alphabeta(depth)

	PrintMemory()

	return types.Move{
		A: root.ball,
		B: root.bestChild.ball,
	}
}
