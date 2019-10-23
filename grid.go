package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A Grid is a two-dimensional array of Cells.
type Grid [][]Cell

// Prints a Grid.
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

// Retrieves a Cell at a given position in the grid.
func (g Grid) GetCell(p Position) Cell {
	return g[p.Y][p.X]
}

// Returns the neighbors of a given cell in the grid.
//
// For the game of life, we use the Moore neighborhood.
// See: https://en.wikipedia.org/wiki/Moore_neighborhood
func (g Grid) GetNeighbors(c Cell) []Cell {
	neighboringPositions := []Position{
		Position{c.Position.X - 1, c.Position.Y - 1},
		Position{c.Position.X, c.Position.Y - 1},
		Position{c.Position.X + 1, c.Position.Y - 1},
		Position{c.Position.X - 1, c.Position.Y},
		Position{c.Position.X + 1, c.Position.Y},
		Position{c.Position.X - 1, c.Position.Y + 1},
		Position{c.Position.X, c.Position.Y + 1},
		Position{c.Position.X + 1, c.Position.Y + 1},
	}

	var neighbors []Cell
	for _, position := range neighboringPositions {
		if g.IsOffGrid(position) {
			continue
		}

		neighbors = append(neighbors, g.GetCell(position))
	}

	return neighbors
}

// Checks if a position is "off grid", beyond either the X or Y dimensions of the grid.
func (g Grid) IsOffGrid(pos Position) bool {
	return (pos.X < 0 || pos.Y < 0) || (pos.X > len(g[0])-1 || pos.Y > len(g)-1)
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
