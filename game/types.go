package game

import "time"

type Position string

const (
	UpperPos Position = "upperPlace"
	LowerPos Position = "lowerPlace"

	HeightOfMap = 13
	WidthOfMap  = 11
)

type Side struct {
	Pos        Position `json:"pos"`
	WinStates  []State  `json:"win_states"`
	LoseStates []State  `json:"lose_states"`
}

type State struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *State) Equal(point State) bool {
	if p.X == point.X && p.Y == point.Y {
		return true
	}
	return false
}

func (p *State) NotEqual(point State) bool {
	return !p.Equal(point)
}

type Move struct {
	A State `json:"a"`
	B State `json:"b"`
}

func (e *Move) Equal(edge Move) bool {
	if e.A.Equal(edge.A) && e.B.Equal(edge.B) {
		return true
	}
	if e.B.Equal(edge.A) && e.A.Equal(edge.B) {
		return true
	}
	return false
}

type World struct {
	Moves   []Move `json:"moves"`
	Turn    Player `json:"turn"`
	BallPos State  `json:"ball_pos"`
	Winner  Player `json:"winner"`
}

func (w *World) SetWinner(p Player) {
	w.Winner = p
}

type Action struct {
	Player Player `json:"player"`
	From   State  `json:"from"`
	To     State  `json:"to"`
}

type Player struct {
	Duration time.Duration `json:"duration"`
	Name     string        `json:"name"`
	Side     Side          `json:"side"`
}

func (p *Player) UpdateTimer(t time.Time) {
	now := time.Now()
	p.Duration -= now.Sub(t)
}

func CreateNewPlayer(name string, pos Position) Player {
	var side Side
	if pos == UpperPos {
		side = UpperSide
	} else {
		side = LowerSide
	}

	return Player{
		Name:     name,
		Side:     side,
		Duration: time.Minute * 5,
	}
}
