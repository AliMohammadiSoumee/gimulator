package agent

import (
	"github.com/alidadar7676/gimulator/types"
	"time"
	"fmt"
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

var (
	lastDuration = time.Second * 2
	lastDepth = 10
)

func run(world types.World, name string) types.Move {

	it := newIteration(world, name)
	root := &gameState{
		it: it,
		ball: world.BallPos,
	}

	var depth int
	if lastDuration > time.Second * 1 {
		depth = lastDepth - 1
	} else if lastDuration < time.Millisecond * 200 {
		depth = lastDepth + 1
	} else {
		depth = lastDepth
	}

	fmt.Println("Depth and time:", depth, lastDuration.Seconds(), lastDepth)
	t := time.Now()
	root.alphabeta(depth)
	lastDuration = time.Now().Sub(t)
	lastDepth = depth

	PrintMemory()

	return types.Move{
		A: root.ball,
		B: root.bestChild.ball,
	}
}
