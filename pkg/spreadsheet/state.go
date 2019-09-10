package spreadsheet

import (
	"strings"
)

// Position is a position value for the spreadsheet
type Position struct {
	Column string
	Row    int
}

type State struct {
	Columns []string
	Rows    []int
	Active  Position
	Cells   map[Position]string
}


func Init() State {
	return State{
		Rows:    makeRange(1, 15),
		Columns: strings.Split("ABCDEFGHIJK", ""),
		Active:  Position{},
		Cells:   make(map[Position]string),
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}

	return a
}