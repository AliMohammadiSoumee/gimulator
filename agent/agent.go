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
		return -inf
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

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	targetPoint := gs.it.winStates[1]
	a := abs(gs.ball.X-targetPoint.X) + abs(gs.ball.Y-targetPoint.Y)
	return -a * -a * -a + gs.it.moveNum
}



func run(world types.World, name string) types.Move {
	depth := 8
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
