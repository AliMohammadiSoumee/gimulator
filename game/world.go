package game

import (
	"github.com/alidadar7676/gimulator/types"
	"fmt"
	"time"
)

type turnState string

const (
	changeTurn turnState = "changeTurn"
	fixedTurn  turnState = "fixedTurn"
	noTurn     turnState = "noTurn"
)

type players struct {
	p1, p2 types.Player
}

func (p *players) changeTurn(player types.Player) types.Player {
	if player.Name == p.p1.Name {
		return p.p2
	} else {
		return p.p1
	}
}

func (p *players) updateTime(name string) {
	now := time.Now()
	if name == p.p1.Name {
		p.p1.Duration += now.Sub(p.p1.LastAction)
	} else {
		p.p2.Duration += now.Sub(p.p2.LastAction)
	}
	p.p1.LastAction = now
	p.p2.LastAction = now
}

type game struct {
	players players
}


func (g *game) update(action types.Action, world types.World) (types.World, error) {
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

	return world, nil
}

func (g *game)updateWorld(action types.Action, world types.World, turnState turnState) types.World {
	world.BallPos = action.To

	move := types.Move{
		A: action.From,
		B: action.To,
	}
	world.Moves = append(world.Moves, move)

	switch turnState {
	case changeTurn:
		world.Turn = g.players.changeTurn(world.Turn)
	case noTurn:
		world.Turn = types.Player{}
	}

	return world
}
