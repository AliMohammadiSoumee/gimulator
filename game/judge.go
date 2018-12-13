package game

var (
	dirX = []int{1, 1, 0, -1, -1, -1, 0, 1}
	dirY = []int{0, 1, 1, 1, 0, -1, -1, -1}
)

type ActionResult string

const (
	ValidActionWithPrice ActionResult = "validActionWithPrice"
	ValidAction          ActionResult = "validAction"
	InvalidAction        ActionResult = "invalidAction"
	WinningAction        ActionResult = "winningAction"
	LosingAction         ActionResult = "losingAction"
)

func Judge(action Action, world World) ActionResult {
	if !world.BallPos.Equal(action.From) {
		return InvalidAction
	}

	validMoves := createValidMoves(world.BallPos, world.Moves)
	playgroundAngs := createPlaygroundAngles(world.Moves)
	playerMove := Move{
		A: action.From,
		B: action.To,
	}

	winningStates := world.Player1.Side.WinStates
	losingStates := world.Player1.Side.LoseStates
	if action.PlayerName == world.Player2.Name {
		winningStates = world.Player2.Side.WinStates
		losingStates = world.Player2.Side.LoseStates
	}

	if !isValidMove(playerMove, validMoves) {
		return InvalidAction
	}
	if inStates(action, winningStates) {
		return WinningAction
	}
	if inStates(action, losingStates) {
		return LosingAction
	}
	if isBlockedState(action.To, playgroundAngs) {
		return LosingAction
	}
	if isValidActionWithPrice(action.To, playgroundAngs) {
		return ValidActionWithPrice
	}
	return ValidAction
}

func createValidMoves(ball State, moves []Move) []Move {
	var validMoves []Move

	for ind := 0; ind < 8; ind++ {
		x := ball.X + dirX[ind]
		y := ball.Y + dirY[ind]
		validMove := Move{
			A: ball,
			B: State{
				X: x,
				Y: y,
			},
		}

		isValid := true
		for _, m := range moves {
			if validMove.Equal(m) {
				isValid = false
			}
		}
		if isValid {
			validMoves = append(validMoves, validMove)
		}
	}
	return validMoves
}

func createPlaygroundAngles(moves []Move) [][]int {
	var playground = make([][]int, HeightOfMap+1)
	for i := 0; i < HeightOfMap+1; i++ {
		playground[i] = make([]int, WidthOfMap+1)
	}

	for _, move := range moves {
		a := move.A
		b := move.B
		playground[a.X][a.Y]++
		playground[b.X][b.Y]++
	}

	return playground
}

func isValidMove(move Move, validMoves []Move) bool {
	for _, m := range validMoves {
		if move.Equal(m) {
			return true
		}
	}
	return false
}

func inStates(action Action, states []State) bool {
	p := action.To
	for _, s := range states {
		if p.Equal(s) {
			return true
		}
	}
	return false
}

func isBlockedState(state State, playground [][]int) bool {
	if playground[state.X][state.Y] >= 7 {
		return true
	}
	return false
}

func isValidActionWithPrice(state State, playground [][]int) bool {
	if playground[state.X][state.Y] > 0 {
		return true
	}
	return false
}
