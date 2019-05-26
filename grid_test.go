package life

import "testing"

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
