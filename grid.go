package life

import "fmt"

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
