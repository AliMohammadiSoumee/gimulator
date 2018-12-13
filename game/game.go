package game

func Update(action Action, world World) World {
	const (
		changeTurn = "otherPlayer"
		fixedTurn  = "fixedTurn"
		noTurn     = "noTurn"
	)

	updateWorld := func(turnState string) {
		world.BallPos = action.To
		world.Moves = append(world.Moves,
			Move{
				A: action.From,
				B: action.To,
			})
		switch turnState {
		case changeTurn:
			world.Turn = world.otherPlayer(action.PlayerName)
		case noTurn:
			world.Turn = ""
		}
	}

	if action.PlayerName != world.Turn {
		return world
	}

	world.updateTimer(action.PlayerName)

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
		world.Winner = world.otherPlayer(action.PlayerName)
	}

	world.setLastAction()

	return world
}
