package game

import (
	"time"
	"math/rand"
)

type turnState string

const (
	changeTurn turnState = "otherPlayer"
	fixedTurn  turnState = "fixedTurn"
	noTurn     turnState = "noTurn"
)

type game struct {
	p1         Player
	p2         Player
	lastAction time.Time
}

func (g *game) otherPlayer(name string) Player {
	if name == g.p1.Name {
		return g.p2
	} else {
		return g.p1
	}
}

func (g *game) updateTimer(name string) {
	if name == g.p1.Name {
		g.p1.UpdateTimer(g.lastAction)
	} else {
		g.p2.UpdateTimer(g.lastAction)
	}
}

func (g *game) setLastAction() {
	g.lastAction = time.Now()
}

func (g *game) update(action Action, world World) World {
	if action.Player.Name != world.Turn.Name {
		return world
	}

	g.updateTimer(action.Player.Name)

	actionRes := judge(action, world)

	switch actionRes {
	case invalidAction:
		return world
	case validAction:
		world = g.updateWorld(action, world, changeTurn)
	case validActionWithPrice:
		world = g.updateWorld(action, world, fixedTurn)
	case winningAction:
		world = g.updateWorld(action, world, noTurn)
		world.SetWinner(action.Player)
	case losingAction:
		world = g.updateWorld(action, world, noTurn)
		world.SetWinner(g.otherPlayer(action.Player.Name))
	}

	g.setLastAction()

	return world
}

func (g *game) updateWorld(action Action, world World, turnState turnState) World {
	world.BallPos = action.To

	move := Move{
		A: action.From,
		B: action.To,
	}
	world.Moves = append(world.Moves, move)

	switch turnState {
	case changeTurn:
		world.Turn = g.otherPlayer(action.Player.Name)
	case noTurn:
		world.Turn = Player{}
	}

	return world
}

func createNewGame(pName1, pName2 string) (World, game) {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(2)
	var player1, player2 Player

	if rnd == 0 {
		player1 = CreateNewPlayer(pName1, LowerPos)
		player2 = CreateNewPlayer(pName2, UpperPos)
	} else {
		player1 = CreateNewPlayer(pName2, LowerPos)
		player2 = CreateNewPlayer(pName1, UpperPos)
	}

	world := World{
		Moves:   initMoves,
		Turn:    player1,
		BallPos: State{X: 6, Y: 7},
	}

	gm := game{
		p1:         player1,
		p2:         player2,
		lastAction: time.Now(),
	}

	return world, gm
}

