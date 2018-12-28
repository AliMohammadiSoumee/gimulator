package agent

import (
	"github.com/alidadar7676/gimulator/types"
	"time"
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

	dis := gs.it.distanceFromWinStates()
	return -dis * dis + gs.it.moveNum
}



func run(world types.World, name string) types.Move {
	time.Sleep(time.Millisecond * 100)
	depth := 4
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
