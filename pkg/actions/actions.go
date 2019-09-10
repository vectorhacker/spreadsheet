package actions

import (
	"spreadsheet/pkg/spreadsheet"
)

type StartEdit struct {
	Position spreadsheet.Position
}

type UpdateValue struct {
	Position spreadsheet.Position
	Value    string
}
