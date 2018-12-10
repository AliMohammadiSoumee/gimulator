package game

import (
	"github.com/alidadar7676/gimulator/types"
)

var (
	dirX = []int{1, 1, 0, -1, -1, -1, 0, 1}
	dirY = []int{0, 1, 1, 1, 0, -1, -1, -1}
)

type actionResult string

const (
	validActionWithPrice actionResult = "validActionWithPrice"
	validAction          actionResult = "validAction"
	invalidAction        actionResult = "invalidAction"
	winningAction        actionResult = "winningAction"
	losingAction         actionResult = "losingAction"
)

func judge(action types.Action, world types.World) actionResult {
	if world.BallPos.NotEqual(action.From) {
		return invalidAction
	}

	playgroundAngs := createPlaygroundAngles(world.Moves)
	validMoves := createValidMoves(world.BallPos, world.Moves)
	playerMove := types.Move{
		A: action.From,
		B: action.To,
	}

	if !isValidMove(playerMove, validMoves) {
		return invalidAction
	}
	if winningWithGoal(action) {
		return winningAction
	}
	if losingWithGoal(action) {
		return losingAction
	}
	if isBlockedState(action.To, playgroundAngs) {
		return losingAction
	}
	if isValidActionWithPrice(action.To, playgroundAngs) {
		return validActionWithPrice
	}
	return validAction
}

func createValidMoves(ball types.State, moves[]types.Move) []types.Move {
	var validMoves []types.Move

	for ind := 0; ind < 8; ind++ {
		x := ball.X + dirX[ind]
		y := ball.Y + dirY[ind]
		validMove := types.Move{
			A: ball,
			B: types.State{
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

func createPlaygroundAngles(moves []types.Move) [][]int {
	var playground = make([][]int, types.HeightOfMap + 1)
	for i := 0; i < types.HeightOfMap + 1; i++ {
		playground[i] = make([]int, types.WidthOfMap + 1, 0)
	}

	for _, move := range moves {
		a := move.A
		b := move.B
		playground[a.X][a.Y]++
		playground[b.X][b.Y]++
	}

	return playground
}

func isValidMove(move types.Move, validMoves []types.Move) bool {
	for _, m := range validMoves {
		if move.Equal(m) {
			return true
		}
	}
	return false
}

func winningWithGoal(action types.Action) bool {
	p := action.To
	for _, winState := range action.Player.Side.WinStates {
		if p.Equal(winState) {
			return true
		}
	}
	return false
}

func losingWithGoal(action types.Action) bool {
	p := action.To
	for _, loseState := range action.Player.Side.LoseStates {
		if p.Equal(loseState) {
			return true
		}
	}
	return false
}

func isBlockedState(state types.State, playground [][]int) bool {
	if playground[state.X][state.Y] >= 7 {
		return true
	}
	return false
}

func isValidActionWithPrice(state types.State, playground [][]int) bool {
	if playground[state.X][state.Y] > 0 {
		return true
	}
	return false
}
