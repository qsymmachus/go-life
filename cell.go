package main

import "fmt"

type Position struct {
	X int
	Y int
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
