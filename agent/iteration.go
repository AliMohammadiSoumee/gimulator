package agent

import (
	"github.com/alidadar7676/gimulator/types"
)

var (
	dirx = []int{1, 1, 0, -1, -1, -1, 0, 1}
	diry = []int{0, 1, 1, 1, 0, -1, -1, -1}
)

type iteration struct {
	ball       types.State
	winStates  []types.State
	loseStates []types.State
	playground [][]map[int]struct{}
	hashTable map[string]*gameState
	width     int
	height    int
	layer     int
	disWin    [][]int
	disLose   [][]int
}

func (it *iteration) validMoves() []types.Move {
	return it.validMovesWithState(it.ball)
}

func (it *iteration) validMovesWithState(state types.State) []types.Move {
	validMoves := make([]types.Move, 0)
	x := state.X
	y := state.Y

	exist := make([]bool, 10)
	for i := range it.playground[x][y] {
		exist[i] = true
	}

	for i := 0; i < 8; i++ {
		if exist[i] {
			continue
		}
		xx := x + dirx[i]
		yy := y + diry[i]
		if xx < 1 || xx > it.width || yy < 1 || yy > it.height {
			continue
		}
		validMoves = append(validMoves, types.Move{
			A: state,
			B: types.State{X: xx, Y: yy},
		})
	}
	return validMoves
}

func (it *iteration) next(move types.Move) {
	it.layer++

	da, db := getDirection(move)
	it.playground[move.A.X][move.A.Y][da] = struct{}{}
	it.playground[move.B.X][move.B.Y][db] = struct{}{}

	it.ball = move.B
}

func (it *iteration) prev(move types.Move) {
	it.layer--

	da, db := getDirection(move)
	delete(it.playground[move.A.X][move.A.Y], da)
	delete(it.playground[move.B.X][move.B.Y], db)

	it.ball = move.A
}

func (it *iteration) hasPrice() bool {
	x := it.ball.X
	y := it.ball.Y
	if len(it.playground[x][y]) > 1 {
		return true
	}
	return false
}

func (it *iteration) isBlockingState() bool {
	x := it.ball.X
	y := it.ball.Y
	if len(it.playground[x][y]) >= 8 {
		return true
	}
	return false
}

func (it *iteration) isWinState() bool {
	for _, ws := range it.winStates {
		if ws.Equal(it.ball) {
			return true
		}
	}
	return false
}

func (it *iteration) isLoseState() bool {
	for _, ws := range it.loseStates {
		if ws.Equal(it.ball) {
			return true
		}
	}
	return false
}

/*
func (it *iteration) distanceFromWinStates() int {
	var val = 0
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for _, win := range it.winStates {
		a := abs(it.ball.X-win.X) + abs(it.ball.Y-win.Y)
		if a > val {
			val = a
		}
	}
	return val
}
*/


type node struct {
	state  types.State
	height int
}

func (it *iteration) bfs(states []types.State, dp [][]int) {
	mark := make([][]bool, it.width+5)
	for i := 0; i < it.width+5; i++ {
		mark[i] = make([]bool, it.height+5)
	}

	for i := 0; i < it.width+2; i++ {
		for j := 0; j < it.height+2; j++ {
			dp[i][j] = 30
			mark[i][j] = false
		}
	}

	que := make(chan node, 10000)
	for _, state := range states {
		que <- node{state: state, height: 0}
	}

	for len(que) > 0 {
		head := <-que
		/*
		if mark[head.state.X][head.state.Y] {
			continue
		}
		mark[head.state.X][head.state.Y] = true
		*/
		if dp[head.state.X][head.state.Y] <= head.height {
			continue
		}
		dp[head.state.X][head.state.Y] = head.height

		validMoves := it.validMovesWithState(head.state)
		for _, mv := range validMoves {
			x := mv.B.X
			y := mv.B.Y
			if mark[x][y] {
				continue
			}
			que <- node{state: mv.B, height: head.height + 1}
		}
	}
}

/*
func (it *iteration) checkMap() {
	for i := 0; i < it.width+1; i++ {
		for j := 0; j < it.height+1; j++ {
			it.mark[i][j] = false
		}
	}
	it.dfs()
}

func (it *iteration) dfs() {
	if it.mark[it.ball.X][it.ball.Y] {
		return
	}
	it.mark[it.ball.X][it.ball.Y] = true

	valids := it.validMoves()
	for _, mv := range valids {
		it.next(mv)
		it.dfs()
		it.prev(mv)
	}
}
*/

func newIteration(world types.World, name string) *iteration {
	player := types.Player{}
	if world.Player1.Name == name {
		player = world.Player1
	} else if world.Player2.Name == name {
		player = world.Player2
	} else {
		panic("Daste Khar")
	}

	pg := newPlayground(types.InitMoves, world.Width, world.Height)

	disWin := make([][]int, world.Width+5)
	for i := 0; i < world.Width+5; i++ {
		disWin[i] = make([]int, world.Height+5)
	}
	disLose := make([][]int, world.Width+5)
	for i := 0; i < world.Width+5; i++ {
		disLose[i] = make([]int, world.Height+5)
	}
	mark := make([][]bool, world.Width+5)
	for i := 0; i < world.Width+5; i++ {
		mark[i] = make([]bool, world.Height+5)
	}

	it := &iteration{
		ball:       world.BallPos,
		winStates:  player.Side.WinStates,
		loseStates: player.Side.LoseStates,
		playground: pg,
		width:      world.Width,
		height:     world.Height,
		hashTable:  make(map[string]*gameState),
		disWin:     disWin,
		disLose:    disLose,
		layer:      0,
	}
	it.bfs(it.winStates, it.disWin)
	it.bfs(it.loseStates, it.disLose)

	pg = newPlayground(world.Moves, world.Width, world.Height)
	it.playground = pg

	/*
	for j := 1; j <= it.height; j++ {
		for i := 1; i <= it.width; i++ {
			fmt.Print(it.disWin[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("----------------")
	for j := 1; j <= it.height; j++ {
		for i := 1; i <= it.width; i++ {
			fmt.Print(it.disLose[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("--------=================--------")
	*/

	return it
}

func newPlayground(moves []types.Move, width, height int) [][]map[int]struct{} {
	pg := make([][]map[int]struct{}, width+5)
	for i := 0; i < width+5; i++ {
		pg[i] = make([]map[int]struct{}, height+5)
		for j := 0; j < height+5; j++ {
			pg[i][j] = make(map[int]struct{})
		}
	}

	for _, mv := range moves {
		da, db := getDirection(mv)

		pg[mv.A.X][mv.A.Y][da] = struct{}{}
		pg[mv.B.X][mv.B.Y][db] = struct{}{}
	}

	return pg
}

func getDirection(mv types.Move) (int, int) {

	flag := false
	var dira int
	for i := 0; i < 8; i++ {
		x := mv.A.X + dirx[i]
		y := mv.A.Y + diry[i]

		if x == mv.B.X && y == mv.B.Y {
			dira = i
			flag = true
		}
	}

	if !flag {
		panic("POOOF1")
	}


	flag = false
	var dirb int
	for i := 0; i < 8; i++ {
		x := mv.B.X + dirx[i]
		y := mv.B.Y + diry[i]

		if x == mv.A.X && y == mv.A.Y {
			dirb = i
			flag = true
		}
	}
	if !flag {
		panic("POOOF2")
	}


	return dira, dirb
}

/*
func newMoveByte(moves []types.Move) []byte {
	moveByte := make([]byte, 200)
	for _, mv := range moves {
		toggleMoveByte(moveByte, mv)
	}
	return moveByte
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
*/
