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
