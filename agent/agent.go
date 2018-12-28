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

	a := gs.it.disWin[gs.ball.X][gs.ball.Y] + 2
	return -a*a + gs.it.layer
}

var (
	lastDuration = time.Second
	lastDepth    = 10
)

func run(world types.World, name string) types.Move {

	it := newIteration(world, name)
	root := &gameState{
		it: it,
		ball: world.BallPos,
	}

	var depth int
	if lastDuration > time.Millisecond * 1700 {
		depth = lastDepth - 2
	} else if lastDuration > time.Second * 1 {
		depth = lastDepth - 1
	} else if lastDuration < time.Millisecond * 400 {
		depth = lastDepth + 1
	} else {
		depth = lastDepth
	}

	//fmt.Println(depth, lastDuration.Seconds(), len(world.Moves))
	t := time.Now()
	root.alphabeta(depth)
	lastDuration = time.Now().Sub(t)
	lastDepth = depth

	//PrintMemory()

	ans := types.Move{
		A: root.ball,
		B: root.bestChild.ball,
	}
	return ans
}
