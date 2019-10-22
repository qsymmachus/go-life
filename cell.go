package main

import "fmt"

type Position struct {
	X int
	Y int
}

// Checks if a position is "off grid", that is, either X or Y is less than zero.
func (p Position) IsOffGrid() bool {
	return p.X < 0 || p.Y < 0
}

type Cell struct {
	Position Position
	Alive    bool
}

// Prints a cell, '*' if it alive, ' ' if it is dead.
func (c Cell) String() string {
	if c.Alive {
		return "*"
	} else {
		return " "
	}
}

// Prints a position in an easy-to-read format.
func (p Position) String() string {
	return fmt.Sprintf("[x: %d, y: %d]", p.X, p.Y)
}
