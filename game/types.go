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
	Pos        Position
	WinStates  []State
	LoseStates []State
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
}

type Action struct {
	Player Player `json:"player"`
	From   State  `json:"from"`
	To     State  `json:"to"`
}

type Player struct {
	Duration   time.Duration
	Name       string
	Side       Side
	LastAction time.Time
}

func (p *Player) UpdateTimer() {
	now := time.Now()
	p.Duration += now.Sub(p.LastAction)
}

func (p *Player) SetLastAction() {
	p.LastAction = time.Now()
}
