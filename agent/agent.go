package agent

import (
	"github.com/alidadar7676/gimulator/types"
	"github.com/alidadar7676/gimulator/game"
)

type gameState struct {
	player   types.Player
	ball     types.State
	child    []*gameState
	moves    []types.Move
	bestNode *gameState
	benefit  int
}

func (gs *gameState) GetPos() types.State {
	return gs.ball
}

func (gs *gameState) HasPrice() bool {
	pg := game.CreatePlaygroundAngles(gs.moves)
	/*
	for i := 1; i < len(pg[0]); i++ {
		for j := 1; j < len(pg); j++ {
			fmt.Printf("%d ", pg[j][i])
		}
		fmt.Println()
	}
	fmt.Println(gs.ball)
	*/
	pg[gs.ball.X][gs.ball.Y]--
	if game.IsValidActionWithPrice(gs.ball, pg) {
		return true
	}
	return false
}

func (gs *gameState) Hit(ben int, child Node) {
	//log.Println("----> ", gs.ball, ben)
	gs.benefit = ben
	gs.bestNode, _ = child.(*gameState)
}

func (gs *gameState) Neighbor() []Node {
	if gs.child == nil {
		gs.child = make([]*gameState, 0, 8)
		validMoves := game.CreateValidMoves(gs.ball, gs.moves)

		//log.Println(gs.ball, validMoves)

		for _, mv := range validMoves {
			cpy := copyMoves(gs.moves)
			gs.child = append(gs.child, &gameState{
				player: gs.player,
				ball:   mv.B,
				moves:  append(cpy, mv),
			})
		}
	}
	nodes := make([]Node, len(gs.child))
	for i := range gs.child {
		nodes[i] = gs.child[i]
	}
	return nodes
}

func (gs *gameState) Equal(node Node) bool {
	return false
}

func (gs *gameState) Heuristic() int {
	for _, ws := range gs.player.Side.WinStates {
		if ws.Equal(gs.ball) {
			return inf
		}
	}

	for _, ls := range gs.player.Side.LoseStates {
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

	targetPoint := gs.player.Side.WinStates[1]
	//log.Println(targetPoint, gs.ball)
	a := abs(gs.ball.X-targetPoint.X) + abs(gs.ball.Y-targetPoint.Y)
	//log.Println(abs(gs.ball.X - targetPoint.X), "--------", abs(gs.ball.Y - targetPoint.Y), "=======", a)
	return -a
	//return 0
}

type agent struct {
	world types.World
	root  *gameState
}

func (a *agent) run(world types.World, player types.Player) types.Move {
	a.world = world
	a.root = &gameState{
		player: player,
		ball:   world.BallPos,
		moves:  world.Moves,
	}

	Minimax(a.root, 6, true)

	return types.Move{
		A: a.root.ball,
		B: a.root.bestNode.ball,
	}
}

func copyMoves(moves []types.Move) []types.Move {
	cpy := make([]types.Move, len(moves))
	for i := range moves {
		cpy[i] = moves[i]
	}
	return cpy
}
