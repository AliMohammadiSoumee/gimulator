package game

import (
	"fmt"
	"time"
	"math/rand"
)

type turnState string

const (
	changeTurn turnState = "changeTurn"
	fixedTurn  turnState = "fixedTurn"
	noTurn     turnState = "noTurn"
)

type players struct {
	p1, p2 Player
}

func (p *players) changeTurn(player Player) Player {
	if player.Name == p.p1.Name {
		return p.p2
	} else {
		return p.p1
	}
}

func (p *players) updateTime(name string) {
	if name == p.p1.Name {
		p.p1.UpdateTimer()
	} else {
		p.p2.UpdateTimer()
	}
}

func (p *players) setLastAction(name string) {
	if name == p.p1.Name {
		p.p1.SetLastAction()
	} else {
		p.p2.SetLastAction()
	}
}

type game struct {
	players players
}

func createNewGame(pName1, pName2 string) (*World, *game) {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(2)
	var player1, player2 Player

	if rnd == 0 {
		player1 = CreateNewPlayer(pName1, LowerPos)
		player2 = CreateNewPlayer(pName2, UpperPos)
	} else {
		player1 = CreateNewPlayer(pName1, LowerPos)
		player2 = CreateNewPlayer(pName2, UpperPos)
	}
	gm := game{players: players{p1: player1, p2: player2}}

	world := World{
		Moves: initMoves,
		Turn: player1,
		BallPos: State{X: 6, Y: 7},
	}

	gm.players.setLastAction(player1.Name)

	return &world, &gm
}

func CreateNewPlayer(name string, pos Position) Player {
	var side Side
	if pos == UpperPos {
		side = UpperSide
	} else {
		side = LowerSide
	}

	return Player{
		Name: name,
		Side: side,
		Duration: time.Second * 5,
		LastAction: time.Now(),
	}
}

func (g *game) update(action Action, world World) (World, error) {
	if action.Player.Name != world.Turn.Name {
		return world, fmt.Errorf("outOfTurn")
	}

	g.players.updateTime(action.Player.Name)

	actionRes := judge(action, world)

	switch actionRes {
	case invalidAction:
		return world, nil
	case validAction:
		world = g.updateWorld(action, world, changeTurn)
	case validActionWithPrice:
		world = g.updateWorld(action, world, fixedTurn)
	case winningAction:
		world = g.updateWorld(action, world, noTurn)
	case losingAction:
		world = g.updateWorld(action, world, noTurn)
	}

	g.players.setLastAction(action.Player.Name)

	return world, nil
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
		world.Turn = g.players.changeTurn(world.Turn)
	case noTurn:
		world.Turn = Player{}
	}

	return world
}
