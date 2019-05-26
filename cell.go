package life

type Position struct {
	X int
	Y int
}

type Cell struct {
	Position Position
	Alive    bool
}

func (c Cell) String() string {
	if c.Alive {
		return "*"
	} else {
		return " "
	}
}
