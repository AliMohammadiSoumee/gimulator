package game

import "github.com/alidadar7676/gimulator/types"

func Update(action types.Action, world types.World) types.World {
	const (
		changeTurn = "otherPlayer"
		fixedTurn  = "fixedTurn"
		noTurn     = "noTurn"
	)

	updateWorld := func(turnState string) {
		world.BallPos = action.To
		world.Moves = append(world.Moves,
			types.Move{
				A: action.From,
				B: action.To,
			})
		switch turnState {
		case changeTurn:
			world.Turn = world.OtherPlayer(action.PlayerName)
		case noTurn:
			world.Turn = ""
		}
	}

	if action.PlayerName != world.Turn {
		return world
	}

	world.UpdateTimer(action.PlayerName)

	actionRes := Judge(action, world)
	switch actionRes {
	case InvalidAction:
		return world
	case ValidAction:
		updateWorld(changeTurn)
	case ValidActionWithPrice:
		updateWorld(fixedTurn)
	case WinningAction:
		updateWorld(noTurn)
		world.Winner = action.PlayerName
	case LosingAction:
		updateWorld(noTurn)
		world.Winner = world.OtherPlayer(action.PlayerName)
	}

	world.SetLastAction()

	return world
}
