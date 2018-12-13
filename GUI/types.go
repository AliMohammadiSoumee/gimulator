package main

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
	WinStates  []State  `json:"winStates"`
	LoseStates []State  `json:"loseStates"`
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
	Turn    string `json:"turn"`
	BallPos State  `json:"ball_pos"`
	Winner  string `json:"winner"`
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
}

func (w *World) SetWinner(p Player) {
	w.Winner = p.Name
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
