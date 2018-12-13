package main

import (
	"fmt"
)

type worldDrawer struct {
	World
	width, height int
}

func (w worldDrawer) DrawField() string {
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
			html += newCircle(xx, yy, 5, "yellow")
			grid[x+1][y+1] = State{xx, yy}
		}
	}

	for _, move := range initMoves {
		html += newLine(grid[move.A.X][move.A.Y].X, grid[move.A.X][move.A.Y].Y, grid[move.B.X][move.B.Y].X, grid[move.B.X][move.B.Y].Y, "red")
	}

	return html
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func newCircle(cx, cy, r int, col string) string {
	return fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d" stroke="black" stroke-width="1" fill="%s" />`, cx, cy, r, col)
}

func newLine(x1, y1, x2, y2 int, col string) string {
	return fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="2" />`, x1, y1, x2, y2, col)
}

var initMoves = []Move{
	{A: State{X: 1, Y: 9}, B: State{X: 2, Y: 10}},
	{A: State{X: 2, Y: 9}, B: State{X: 2, Y: 10}},
	{A: State{X: 9, Y: 13}, B: State{X: 10, Y: 13}},
	{A: State{X: 9, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 11, Y: 9}, B: State{X: 11, Y: 10}},
	{A: State{X: 1, Y: 8}, B: State{X: 1, Y: 9}},
	{A: State{X: 2, Y: 10}, B: State{X: 1, Y: 11}},
	{A: State{X: 5, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 7, Y: 12}, B: State{X: 6, Y: 13}},
	{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 4}},
	{A: State{X: 10, Y: 12}, B: State{X: 10, Y: 13}},
	{A: State{X: 7, Y: 12}, B: State{X: 8, Y: 12}},
	{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 1}},
	{A: State{X: 10, Y: 2}, B: State{X: 11, Y: 3}},
	{A: State{X: 1, Y: 12}, B: State{X: 1, Y: 13}},
	{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 2, Y: 2}, B: State{X: 3, Y: 2}},
	{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 9, Y: 2}, B: State{X: 10, Y: 2}},
	{A: State{X: 10, Y: 8}, B: State{X: 10, Y: 9}},
	{A: State{X: 2, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 11, Y: 10}, B: State{X: 10, Y: 11}},
	{A: State{X: 2, Y: 2}, B: State{X: 1, Y: 3}},
	{A: State{X: 1, Y: 5}, B: State{X: 2, Y: 6}},
	{A: State{X: 10, Y: 4}, B: State{X: 10, Y: 5}},
	{A: State{X: 2, Y: 9}, B: State{X: 1, Y: 10}},
	{A: State{X: 10, Y: 12}, B: State{X: 11, Y: 13}},
	{A: State{X: 10, Y: 10}, B: State{X: 10, Y: 11}},
	{A: State{X: 2, Y: 12}, B: State{X: 3, Y: 12}},
	{A: State{X: 3, Y: 12}, B: State{X: 4, Y: 13}},
	{A: State{X: 11, Y: 8}, B: State{X: 10, Y: 9}},
	{A: State{X: 8, Y: 2}, B: State{X: 9, Y: 2}},
	{A: State{X: 1, Y: 6}, B: State{X: 2, Y: 7}},
	{A: State{X: 8, Y: 12}, B: State{X: 7, Y: 13}},
	{A: State{X: 3, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 4, Y: 2}, B: State{X: 5, Y: 2}},
	{A: State{X: 11, Y: 2}, B: State{X: 11, Y: 3}},
	{A: State{X: 2, Y: 12}, B: State{X: 1, Y: 13}},
	{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 3}},
	{A: State{X: 7, Y: 2}, B: State{X: 8, Y: 2}},
	{A: State{X: 10, Y: 3}, B: State{X: 10, Y: 4}},
	{A: State{X: 3, Y: 13}, B: State{X: 4, Y: 13}},
	{A: State{X: 4, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 11, Y: 6}, B: State{X: 11, Y: 7}},
	{A: State{X: 10, Y: 9}, B: State{X: 11, Y: 10}},
	{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 7}},
	{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 9}},
	{A: State{X: 8, Y: 13}, B: State{X: 9, Y: 13}},
	{A: State{X: 10, Y: 5}, B: State{X: 11, Y: 6}},
	{A: State{X: 2, Y: 6}, B: State{X: 2, Y: 7}},
	{A: State{X: 11, Y: 10}, B: State{X: 11, Y: 11}},
	{A: State{X: 2, Y: 4}, B: State{X: 1, Y: 5}},
	{A: State{X: 10, Y: 8}, B: State{X: 11, Y: 9}},
	{A: State{X: 9, Y: 12}, B: State{X: 8, Y: 13}},
	{A: State{X: 1, Y: 10}, B: State{X: 1, Y: 11}},
	{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 1}},
	{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 2}},
	{A: State{X: 1, Y: 5}, B: State{X: 1, Y: 6}},
	{A: State{X: 2, Y: 2}, B: State{X: 2, Y: 3}},
	{A: State{X: 8, Y: 1}, B: State{X: 7, Y: 2}},
	{A: State{X: 11, Y: 12}, B: State{X: 10, Y: 13}},
	{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 1}},
	{A: State{X: 11, Y: 9}, B: State{X: 10, Y: 10}},
	{A: State{X: 11, Y: 7}, B: State{X: 11, Y: 8}},
	{A: State{X: 10, Y: 10}, B: State{X: 11, Y: 11}},
	{A: State{X: 8, Y: 12}, B: State{X: 9, Y: 13}},
	{A: State{X: 1, Y: 2}, B: State{X: 1, Y: 3}},
	{A: State{X: 11, Y: 11}, B: State{X: 11, Y: 12}},
	{A: State{X: 2, Y: 7}, B: State{X: 1, Y: 8}},
	{A: State{X: 5, Y: 12}, B: State{X: 4, Y: 13}},
	{A: State{X: 11, Y: 2}, B: State{X: 10, Y: 3}},
	{A: State{X: 1, Y: 12}, B: State{X: 2, Y: 13}},
	{A: State{X: 10, Y: 7}, B: State{X: 11, Y: 7}},
	{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 1, Y: 6}, B: State{X: 1, Y: 7}},
	{A: State{X: 10, Y: 3}, B: State{X: 11, Y: 4}},
	{A: State{X: 1, Y: 9}, B: State{X: 2, Y: 9}},
	{A: State{X: 7, Y: 1}, B: State{X: 6, Y: 2}},
	{A: State{X: 10, Y: 6}, B: State{X: 11, Y: 6}},
	{A: State{X: 10, Y: 3}, B: State{X: 11, Y: 3}},
	{A: State{X: 10, Y: 5}, B: State{X: 10, Y: 6}},
	{A: State{X: 1, Y: 10}, B: State{X: 2, Y: 10}},
	{A: State{X: 11, Y: 5}, B: State{X: 11, Y: 6}},
	{A: State{X: 1, Y: 13}, B: State{X: 2, Y: 13}},
	{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 5}},
	{A: State{X: 1, Y: 11}, B: State{X: 2, Y: 11}},
	{A: State{X: 7, Y: 12}, B: State{X: 8, Y: 13}},
	{A: State{X: 2, Y: 12}, B: State{X: 2, Y: 13}},
	{A: State{X: 11, Y: 3}, B: State{X: 10, Y: 4}},
	{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 1}},
	{A: State{X: 8, Y: 1}, B: State{X: 8, Y: 2}},
	{A: State{X: 10, Y: 11}, B: State{X: 10, Y: 12}},
	{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 1}},
	{A: State{X: 4, Y: 12}, B: State{X: 4, Y: 13}},
	{A: State{X: 4, Y: 1}, B: State{X: 4, Y: 2}},
	{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 2}},
	{A: State{X: 2, Y: 8}, B: State{X: 2, Y: 9}},
	{A: State{X: 1, Y: 5}, B: State{X: 2, Y: 5}},
	{A: State{X: 3, Y: 2}, B: State{X: 4, Y: 2}},
	{A: State{X: 10, Y: 7}, B: State{X: 10, Y: 8}},
	{A: State{X: 3, Y: 12}, B: State{X: 2, Y: 13}},
	{A: State{X: 7, Y: 13}, B: State{X: 8, Y: 13}},
	{A: State{X: 2, Y: 4}, B: State{X: 2, Y: 5}},
	{A: State{X: 11, Y: 11}, B: State{X: 10, Y: 12}},
	{A: State{X: 2, Y: 12}, B: State{X: 3, Y: 13}},
	{A: State{X: 11, Y: 4}, B: State{X: 10, Y: 5}},
	{A: State{X: 2, Y: 10}, B: State{X: 2, Y: 11}},
	{A: State{X: 11, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 10, Y: 13}, B: State{X: 11, Y: 13}},
	{A: State{X: 11, Y: 5}, B: State{X: 10, Y: 6}},
	{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 4}},
	{A: State{X: 6, Y: 1}, B: State{X: 7, Y: 1}},
	{A: State{X: 9, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 10, Y: 9}, B: State{X: 11, Y: 9}},
	{A: State{X: 11, Y: 3}, B: State{X: 11, Y: 4}},
	{A: State{X: 2, Y: 3}, B: State{X: 2, Y: 4}},
	{A: State{X: 5, Y: 1}, B: State{X: 5, Y: 2}},
	{A: State{X: 7, Y: 12}, B: State{X: 7, Y: 13}},
	{A: State{X: 10, Y: 5}, B: State{X: 11, Y: 5}},
	{A: State{X: 10, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 10, Y: 11}, B: State{X: 11, Y: 12}},
	{A: State{X: 10, Y: 11}, B: State{X: 11, Y: 11}},
	{A: State{X: 10, Y: 4}, B: State{X: 11, Y: 4}},
	{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 1}},
	{A: State{X: 4, Y: 12}, B: State{X: 3, Y: 13}},
	{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 2}},
	{A: State{X: 4, Y: 12}, B: State{X: 5, Y: 12}},
	{A: State{X: 2, Y: 11}, B: State{X: 1, Y: 12}},
	{A: State{X: 9, Y: 12}, B: State{X: 10, Y: 12}},
	{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 1}},
	{A: State{X: 10, Y: 7}, B: State{X: 11, Y: 8}},
	{A: State{X: 1, Y: 3}, B: State{X: 1, Y: 4}},
	{A: State{X: 3, Y: 1}, B: State{X: 2, Y: 2}},
	{A: State{X: 10, Y: 6}, B: State{X: 11, Y: 7}},
	{A: State{X: 10, Y: 9}, B: State{X: 10, Y: 10}},
	{A: State{X: 1, Y: 10}, B: State{X: 2, Y: 11}},
	{A: State{X: 2, Y: 5}, B: State{X: 2, Y: 6}},
	{A: State{X: 1, Y: 7}, B: State{X: 1, Y: 8}},
	{A: State{X: 2, Y: 3}, B: State{X: 1, Y: 4}},
	{A: State{X: 11, Y: 6}, B: State{X: 10, Y: 7}},
	{A: State{X: 11, Y: 8}, B: State{X: 11, Y: 9}},
	{A: State{X: 2, Y: 6}, B: State{X: 1, Y: 7}},
	{A: State{X: 9, Y: 12}, B: State{X: 9, Y: 13}},
	{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 10, Y: 2}, B: State{X: 11, Y: 2}},
	{A: State{X: 1, Y: 11}, B: State{X: 2, Y: 12}},
	{A: State{X: 10, Y: 1}, B: State{X: 10, Y: 2}},
	{A: State{X: 11, Y: 12}, B: State{X: 11, Y: 13}},
	{A: State{X: 2, Y: 11}, B: State{X: 2, Y: 12}},
	{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 3}},
	{A: State{X: 8, Y: 12}, B: State{X: 8, Y: 13}},
	{A: State{X: 2, Y: 7}, B: State{X: 2, Y: 8}},
	{A: State{X: 11, Y: 7}, B: State{X: 10, Y: 8}},
	{A: State{X: 10, Y: 12}, B: State{X: 11, Y: 12}},
	{A: State{X: 1, Y: 9}, B: State{X: 1, Y: 10}},
	{A: State{X: 5, Y: 12}, B: State{X: 5, Y: 13}},
	{A: State{X: 3, Y: 12}, B: State{X: 4, Y: 12}},
	{A: State{X: 1, Y: 6}, B: State{X: 2, Y: 6}},
	{A: State{X: 10, Y: 12}, B: State{X: 9, Y: 13}},
	{A: State{X: 2, Y: 13}, B: State{X: 3, Y: 13}},
	{A: State{X: 7, Y: 1}, B: State{X: 7, Y: 2}},
	{A: State{X: 2, Y: 8}, B: State{X: 1, Y: 9}},
	{A: State{X: 1, Y: 4}, B: State{X: 1, Y: 5}},
	{A: State{X: 4, Y: 13}, B: State{X: 5, Y: 13}},
	{A: State{X: 1, Y: 1}, B: State{X: 1, Y: 2}},
	{A: State{X: 2, Y: 5}, B: State{X: 1, Y: 6}},
	{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 8}},
	{A: State{X: 10, Y: 6}, B: State{X: 10, Y: 7}},
	{A: State{X: 1, Y: 11}, B: State{X: 1, Y: 12}},
	{A: State{X: 2, Y: 1}, B: State{X: 1, Y: 2}},
	{A: State{X: 10, Y: 8}, B: State{X: 11, Y: 8}},
	{A: State{X: 10, Y: 4}, B: State{X: 11, Y: 5}},
	{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 8}},
	{A: State{X: 10, Y: 2}, B: State{X: 10, Y: 3}},
	{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 2}},
	{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 1}},
	{A: State{X: 4, Y: 12}, B: State{X: 5, Y: 13}},
	{A: State{X: 3, Y: 12}, B: State{X: 3, Y: 13}},
	{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 2}},
	{A: State{X: 6, Y: 13}, B: State{X: 7, Y: 13}},
	{A: State{X: 10, Y: 10}, B: State{X: 11, Y: 10}},
	{A: State{X: 8, Y: 12}, B: State{X: 9, Y: 12}},
	{A: State{X: 11, Y: 4}, B: State{X: 11, Y: 5}},
	{A: State{X: 11, Y: 1}, B: State{X: 11, Y: 2}},
	{A: State{X: 9, Y: 12}, B: State{X: 10, Y: 13}},
	{A: State{X: 1, Y: 12}, B: State{X: 2, Y: 12}},
}
