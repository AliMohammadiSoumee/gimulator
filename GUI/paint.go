package GUI

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel"
)

func CreateGame(world World) *imdraw.IMDraw {
	field := createGrid()
	imd := createIMDraw(field, world.Moves)

	return imd
}

func createGrid() [][]State {
	field := make([][]State, WidthOfMap+1)
	for i := 0; i < WidthOfMap+1; i++ {
		field[i] = make([]State, HeightOfMap+1)
	}

	for row := 1; row < WidthOfMap+1; row++ {
		for col := 1; col < HeightOfMap+1; col++ {
			field[row][col] = State{X: row*45 + 100, Y: col*45 + 100}
		}
	}

	return field
}

func createIMDraw(field [][]State, moves []Move) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	for row := 1; row < WidthOfMap+1; row++ {
		for col := 1; col < HeightOfMap+1; col++ {
			imd.Push(pixel.V(float64(field[row][col].X), float64(field[row][col].Y)))
			imd.Circle(3.0, 1.0)
		}
	}

	for _, move := range moves {
		imd.Push(pixel.V(float64(field[move.A.Y][move.A.X].X), float64(field[move.A.Y][move.A.X].Y)))
		imd.Push(pixel.V(float64(field[move.B.Y][move.B.X].X), float64(field[move.B.Y][move.B.X].Y)))
		imd.Line(1)
	}

	return imd
}
