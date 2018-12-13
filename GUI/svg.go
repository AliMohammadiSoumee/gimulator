package main

import (
	"fmt"
)

type worldDrawer struct {
	World
	width, height int
}

func (w worldDrawer) Draw() string {
	var (
		html    = ""
		delta   = min(w.width/(WidthOfMap+1), w.height/(HeightOfMap+1))
		marginx = (w.width - delta*(WidthOfMap-1)) / 2
		marginy = (w.height - delta*(HeightOfMap-1)) / 2
	)
	fmt.Println(w.width, w.height)
	fmt.Println(marginx, marginy)
	fmt.Println(delta)

	grid := make([][]State, WidthOfMap+1)
	for i := 0; i < WidthOfMap+1; i++ {
		grid[i] = make([]State, HeightOfMap+1)
	}

	for x := 0; x < WidthOfMap; x++ {
		for y := 0; y < HeightOfMap; y++ {
			xx := marginx + x*delta
			yy := marginy + y*delta
			html += newCircle(xx, yy, 5)
			grid[x+1][y+1] = State{xx, yy}
		}
	}

	for _, move := range initMoves {
		html += newLine(grid[move.A.Y][move.A.X].X, grid[move.A.Y][move.A.X].Y, grid[move.B.Y][move.B.X].X, grid[move.B.Y][move.B.X].Y)
	}

	return html
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func newSVG(world World) string {
	svg := ""

	grid := newGrid(Height)
	svg = drawGrid(grid, svg)
	svg = drawMoves(world.Moves, grid, svg)

	return svg
}

func newGrid(height int) [][]State {

	dis := height / (HeightOfMap + 1)

	field := make([][]State, WidthOfMap+1)
	for i := 0; i < WidthOfMap+1; i++ {
		field[i] = make([]State, HeightOfMap+1)
	}

	for row := 1; row < WidthOfMap+1; row++ {
		for col := 1; col < HeightOfMap+1; col++ {
			field[row][col] = State{X: row*dis + dis, Y: col*dis + dis}
		}
	}

	return field
}

func drawGrid(grid [][]State, svg string) string {
	for row := 1; row < WidthOfMap+1; row++ {
		for col := 1; col < HeightOfMap+1; col++ {
			svg += newCircle(grid[row][col].X, grid[row][col].Y, 3)
		}
	}
	return svg
}

func drawMoves(moves []Move, grid [][]State, svg string) string {
	for _, move := range moves {
		svg += newLine(grid[move.A.X][move.A.Y].X, grid[move.A.X][move.A.Y].Y, grid[move.B.X][move.B.Y].X, grid[move.B.X][move.B.Y].Y)
	}
	return svg
}

func newCircle(cx, cy, r int) string {
	return fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d" stroke="yellow" stroke-width="1" />`, cx, cy, r)
}

func newLine(x1, y1, x2, y2 int) string {
	return fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="2" />`, x1, y1, x2, y2)
}

var initMoves = []Move{
	{A: State{X: 13, Y: 3}, B: State{X: 12, Y: 4}},
	{A: State{X: 9, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 1, Y: 8}, B: State{X: 1, Y: 9}},
	{A: State{X: 5, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 4}},
	{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 1}},
	{A: State{X: 12, Y: 3}, B: State{X: 12, Y: 4}},
	{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 2, Y: 2}, B: State{X: 3, Y: 2}},
	{A: State{X: 5, Y: 1}, B: State{X: 6, Y: 1}},
	{A: State{X: 12, Y: 1}, B: State{X: 13, Y: 1}},
	{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 9, Y: 2}, B: State{X: 10, Y: 2}},
	{A: State{X: 2, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 13, Y: 8}, B: State{X: 13, Y: 9}},
	{A: State{X: 5, Y: 2}, B: State{X: 6, Y: 2}},
	{A: State{X: 2, Y: 2}, B: State{X: 1, Y: 3}},
	{A: State{X: 12, Y: 7}, B: State{X: 13, Y: 8}},
	{A: State{X: 12, Y: 7}, B: State{X: 13, Y: 7}},
	{A: State{X: 12, Y: 9}, B: State{X: 13, Y: 9}},
	{A: State{X: 11, Y: 1}, B: State{X: 12, Y: 2}},
	{A: State{X: 8, Y: 2}, B: State{X: 9, Y: 2}},
	{A: State{X: 12, Y: 2}, B: State{X: 12, Y: 3}},
	{A: State{X: 13, Y: 3}, B: State{X: 13, Y: 4}},
	{A: State{X: 3, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 4, Y: 2}, B: State{X: 5, Y: 2}},
	{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 3}},
	{A: State{X: 12, Y: 4}, B: State{X: 12, Y: 5}},
	{A: State{X: 7, Y: 2}, B: State{X: 8, Y: 2}},
	{A: State{X: 6, Y: 1}, B: State{X: 7, Y: 2}},
	{A: State{X: 4, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 7}},
	{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 9}},
	{A: State{X: 2, Y: 4}, B: State{X: 1, Y: 5}},
	{A: State{X: 12, Y: 1}, B: State{X: 12, Y: 2}},
	{A: State{X: 12, Y: 8}, B: State{X: 13, Y: 8}},
	{A: State{X: 11, Y: 2}, B: State{X: 12, Y: 2}},
	{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 1}},
	{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 2}},
	{A: State{X: 2, Y: 2}, B: State{X: 2, Y: 3}},
	{A: State{X: 8, Y: 1}, B: State{X: 7, Y: 2}},
	{A: State{X: 1, Y: 5}, B: State{X: 1, Y: 6}},
	{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 1}},
	{A: State{X: 12, Y: 2}, B: State{X: 13, Y: 3}},
	{A: State{X: 1, Y: 2}, B: State{X: 1, Y: 3}},
	{A: State{X: 2, Y: 7}, B: State{X: 1, Y: 8}},
	{A: State{X: 13, Y: 6}, B: State{X: 13, Y: 7}},
	{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 1, Y: 6}, B: State{X: 1, Y: 7}},
	{A: State{X: 1, Y: 9}, B: State{X: 2, Y: 9}},
	{A: State{X: 7, Y: 1}, B: State{X: 6, Y: 2}},
	{A: State{X: 13, Y: 4}, B: State{X: 12, Y: 5}},
	{A: State{X: 13, Y: 1}, B: State{X: 13, Y: 2}},
	{A: State{X: 12, Y: 4}, B: State{X: 13, Y: 4}},
	{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 5}},
	{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 1}},
	{A: State{X: 13, Y: 7}, B: State{X: 13, Y: 8}},
	{A: State{X: 8, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 1}},
	{A: State{X: 4, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 2}},
	{A: State{X: 2, Y: 8}, B: State{X: 2, Y: 9}},
	{A: State{X: 1, Y: 5}, B: State{X: 2, Y: 5}},
	{A: State{X: 3, Y: 2}, B: State{X: 4, Y: 2}},
	{A: State{X: 2, Y: 4}, B: State{X: 2, Y: 5}},
	{A: State{X: 11, Y: 1}, B: State{X: 12, Y: 1}},
	{A: State{X: 11, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 6, Y: 1}, B: State{X: 5, Y: 2}},
	{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 4}},
	{A: State{X: 6, Y: 1}, B: State{X: 7, Y: 1}},
	{A: State{X: 9, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 13, Y: 8}, B: State{X: 12, Y: 9}},
	{A: State{X: 2, Y: 3}, B: State{X: 2, Y: 4}},
	{A: State{X: 5, Y: 1}, B: State{X: 5, Y: 2}},
	{A: State{X: 10, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 12, Y: 8}, B: State{X: 13, Y: 9}},
	{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 1}},
	{A: State{X: 12, Y: 3}, B: State{X: 13, Y: 3}},
	{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 12, Y: 7}, B: State{X: 12, Y: 8}},
	{A: State{X: 12, Y: 5}, B: State{X: 13, Y: 5}},
	{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 1}},
	{A: State{X: 6, Y: 2}, B: State{X: 7, Y: 2}},
	{A: State{X: 1, Y: 3}, B: State{X: 1, Y: 4}},
	{A: State{X: 3, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 12, Y: 4}, B: State{X: 13, Y: 5}},
	{A: State{X: 6, Y: 1}, B: State{X: 6, Y: 2}},
	{A: State{X: 1, Y: 7}, B: State{X: 1, Y: 8}},
	{A: State{X: 2, Y: 3}, B: State{X: 1, Y: 4}},
	{A: State{X: 12, Y: 8}, B: State{X: 12, Y: 9}},
	{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 10, Y: 2}, B: State{X: 11, Y: 2}},
	{A: State{X: 10, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 12, Y: 1}, B: State{X: 13, Y: 2}},
	{A: State{X: 5, Y: 1}, B: State{X: 6, Y: 2}},
	{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 3}},
	{A: State{X: 2, Y: 7}, B: State{X: 2, Y: 8}},
	{A: State{X: 13, Y: 2}, B: State{X: 13, Y: 3}},
	{A: State{X: 7, Y: 1}, B: State{X: 7, Y: 2}},
	{A: State{X: 13, Y: 4}, B: State{X: 13, Y: 5}},
	{A: State{X: 2, Y: 8}, B: State{X: 1, Y: 9}},
	{A: State{X: 13, Y: 5}, B: State{X: 1, Y: 6}},
	{A: State{X: 1, Y: 4}, B: State{X: 1, Y: 5}},
	{A: State{X: 13, Y: 1}, B: State{X: 12, Y: 2}},
	{A: State{X: 1, Y: 1}, B: State{X: 1, Y: 2}},
	{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 8}},
	{A: State{X: 2, Y: 1}, B: State{X: 1, Y: 2}},
	{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 8}},
	{A: State{X: 13, Y: 7}, B: State{X: 12, Y: 8}},
	{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 1}},
	{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 2}},
	{A: State{X: 13, Y: 2}, B: State{X: 12, Y: 3}},
	{A: State{X: 12, Y: 3}, B: State{X: 13, Y: 4}},
	{A: State{X: 11, Y: 1}, B: State{X: 11, Y: 2}},
	{A: State{X: 12, Y: 2}, B: State{X: 13, Y: 2}},
	{A: State{X: 12, Y: 1}, B: State{X: 11, Y: 2}},
}
