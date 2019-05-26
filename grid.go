package life

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid [][]Cell

func (g Grid) String() string {
	str := ""

	for _, row := range g {
		for _, cell := range row {
			str = fmt.Sprintf("%s%s", str, cell)
		}
		str = fmt.Sprintf("%s\n", str)
	}

	return str
}

// Generates a grid with the specified width and height, and Cells with random initial state.
// The 'liveness' parameter determines what % of cells will be alive on the grid.
func RandomGrid(width, height, liveness int) Grid {
	var grid Grid

	for y := 0; y < height; y++ {
		var row []Cell
		for x := 0; x < width; x++ {
			row = append(row, Cell{Position{x, y}, randomBool(liveness)})
		}
		grid = append(grid, row)
	}

	return grid
}

// Returns a boolean that is randomly either true or false.
// The higher the likelihood percentage, the more likely that the boolean will be true.
func randomBool(likelihood int) bool {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(100) <= likelihood
}
