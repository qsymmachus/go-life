package main

import "testing"

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
