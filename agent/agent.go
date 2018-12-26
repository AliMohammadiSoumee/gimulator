package agent

import (
	"github.com/alidadar7676/gimulator/types"
	"github.com/alidadar7676/gimulator/game"
	"encoding/base64"
	"crypto/md5"
)

type gameState struct {
	ball     types.State
	child    []*gameState
	bestNode *gameState
	benefit  int
	price    bool
	heur     int
	hash     string
}

func (gs *gameState) GetPos() types.State {
	return gs.ball
}

func (gs *gameState) hasPrice(moves []types.Move) bool {
	pg := game.CreatePlaygroundAngles(moves)
	pg[gs.ball.X][gs.ball.Y]--
	if game.IsValidActionWithPrice(gs.ball, pg) {
		return true
	}
	return false
}

func (gs *gameState) HasPrice() bool {
	return gs.price
}

func (gs *gameState) Hit(ben int, child Node) {
	gs.benefit = ben
	gs.bestNode, _ = child.(*gameState)
}

func (gs *gameState) Neighbor() []Node {
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
	return gs.heur
}

func (gs *gameState) heuristic(winStates, loseStates []types.State) int {
	for _, ws := range winStates {
		if ws.Equal(gs.ball) {
			return inf
		}
	}

	for _, ls := range loseStates {
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

	targetPoint := winStates[1]
	a := abs(gs.ball.X-targetPoint.X) + abs(gs.ball.Y-targetPoint.Y)
	return -a
}


type agent struct {
	root       *gameState
	nodeTable  map[string]*gameState
	winStates  []types.State
	loseStates []types.State
	moves      []types.Move
}

func newAgent(winState, loseState []types.State, moves []types.Move) *agent {
	return &agent{
		winStates:  winState,
		loseStates: loseState,
		nodeTable:  make(map[string]*gameState),
		moves:      moves,
	}
}

func (a *agent) run(ball types.State) types.Move {

	depth := 6
	a.root = &gameState{
		ball:  ball,
		child: make([]*gameState, 0),
	}

	moveByte := make([]byte, 200)
	for _, mv := range a.moves {
		toggleMoveByte(moveByte, mv)
	}

	a.Graph(moveByte, depth, a.root)

	Minimax(a.root, depth, true)

	return types.Move{
		A: a.root.ball,
		B: a.root.bestNode.ball,
	}
}

func (a *agent) Graph(moveByte []byte, depth int, gs *gameState) *gameState {
	hash := hash(moveByte)
	if node, exist := a.nodeTable[hash]; exist {
		return node
	}
	a.nodeTable[hash] = gs
	gs.hash = hash

	gs.heur = gs.heuristic(a.winStates, a.loseStates)
	gs.price = gs.hasPrice(a.moves)

	if depth == 0 {
		return gs
	}

	validMoves := game.CreateValidMoves(gs.ball, a.moves)
	for _, mv := range validMoves {
		toggleMoveByte(moveByte, mv)
		a.moves = append(a.moves, mv)
		child := &gameState{
			ball:  mv.B,
			child: []*gameState{},
		}
		gs.child = append(gs.child, a.Graph(moveByte, depth-1, child))

		a.moves = a.moves[:len(a.moves)-1]
		toggleMoveByte(moveByte, mv)
	}

	return gs
}

func hash(moveByte []byte) string {
	a := md5.Sum(moveByte)
	b := a[0:]
	return base64.StdEncoding.EncodeToString(b)
}

func toggleMoveByte(moveByte []byte, move types.Move) {
	ind := moveToInt(move)
	div := ind / 8
	mod := ind % 8

	var sum byte = 1 << mod
	moveByte[div] ^= sum
}

func moveToInt(move types.Move) uint {
	a := move.A
	b := move.B

	if a.Y > b.Y {
		a, b = b, a
	} else if a.Y == b.Y && a.X > b.X {
		a, b = b, a
	}

	num := uint(0)
	num += uint((a.Y - 1) * 4 * 11)
	num += uint((a.X - 1) * 4)

	switch {
	case a.X+1 == b.X && a.Y == b.Y:
		num += 1
	case a.X+1 == b.X && a.Y+1 == b.Y:
		num += 2
	case a.X == b.X && a.Y+1 == b.Y:
		num += 3
	case a.X-1 == b.X && a.Y+1 == b.Y:
		num += 4
	}
	return num
}
