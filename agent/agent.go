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

func (gs *gameState) heuristic() int {
	for _, ws := range gs.it.winStates {
		if ws.Equal(gs.ball) {
			return inf
		}
	}

	for _, ls := range gs.it.loseStates {
		if ls.Equal(gs.ball) {
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
	a := abs(gs.ball.X-targetPoint.X) + abs(gs.ball.Y-targetPoint.Y)
	return -a
}

func run(world types.World, name string) types.Move {
	depth := 10
	it := newIteration(world, name)
	root := &gameState{
		it: it,
		ball: world.BallPos,
	}
	//root.minimax(depth)
	root.alphabeta(depth)

	PrintMemory()

	return types.Move{
		A: root.ball,
		B: root.bestChild.ball,
	}
}
