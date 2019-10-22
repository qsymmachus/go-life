package main

import "testing"

func TestPositionIsOffGrid(t *testing.T) {
	positionOnGrid := Position{5, 0}
	positionXOffGrid := Position{-4, 3}
	positionYOffGrid := Position{9, -2}

	if positionOnGrid.IsOffGrid() != false {
		t.Errorf("A position on grid was reported to be off grid")
	}

	if positionXOffGrid.IsOffGrid() != true {
		t.Errorf("A position off grid was reported to be on grid")
	}

	if positionYOffGrid.IsOffGrid() != true {
		t.Errorf("A position off grid was reported to be on grid")
	}
}

func TestCellStringAlive(t *testing.T) {
	cell := Cell{Alive: true}
	got := cell.String()

	if got != "*" {
		t.Errorf("A live cell should print *; got %s", got)
	}
}

func TestCellStringDead(t *testing.T) {
	cell := Cell{Alive: false}
	got := cell.String()

	if got != " " {
		t.Errorf("A live cell should print *; got %s", got)
	}
}
