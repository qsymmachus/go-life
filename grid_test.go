package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCell(t *testing.T) {
	p := Position{1, 0}
	grid := Grid{
		[]Cell{{Alive: true, Position: Position{0, 0}}, {Alive: false, Position: p}},
	}
	want := Cell{Alive: false, Position: p}
	got := grid.GetCell(p)

	if got != want {
		t.Errorf("Did not retrieve expected cell at position %s; got %s", p, got.Position)
	}
}

func TestGetNeighbors(t *testing.T) {
	grid := Grid{
		[]Cell{{Position{0, 0}, true}, {Position{1, 0}, false}, {Position{2, 0}, true}},
		[]Cell{{Position{0, 1}, true}, {Position{1, 1}, false}, {Position{2, 1}, true}},
		[]Cell{{Position{0, 2}, true}, {Position{1, 2}, false}, {Position{2, 2}, true}},
	}

	cell := grid.GetCell(Position{1, 1})
	got := grid.GetNeighbors(cell)
	fmt.Println(got)
	want := []Cell{
		{Position{0, 0}, true},
		{Position{1, 0}, false},
		{Position{2, 0}, true},
		{Position{0, 1}, true},
		{Position{2, 1}, true},
		{Position{0, 2}, true},
		{Position{1, 2}, false},
		{Position{2, 2}, true},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Did not get expected neighbors")
	}
}

func TestStringGrid(t *testing.T) {
	grid := Grid{
		[]Cell{{Alive: true}, {Alive: false}, {Alive: true}},
		[]Cell{{Alive: true}, {Alive: true}, {Alive: true}},
		[]Cell{{Alive: true}, {Alive: false}, {Alive: true}},
	}

	want := "* *\n***\n* *\n"
	got := grid.String()

	if got != want {
		t.Errorf("Test grid should print %s; got %s", want, got)
	}
}

func TestRandomGrid(t *testing.T) {
	grid := RandomGrid(5, 5, 30)

	if len(grid) != 5 {
		t.Errorf("Random grid should have height of 5; got %d", len(grid))
	}

	for _, row := range grid {
		if len(row) != 5 {
			t.Errorf("Random grid should have width of 5; got %d", len(row))
		}
	}
}
