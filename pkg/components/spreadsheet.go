package components

import (
	"fmt"
	"log"

	"spreadsheet/pkg/actions"
	"spreadsheet/pkg/dispatcher"
	"spreadsheet/pkg/spreadsheet"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// Spreadsheet of the spreadsheet
type Spreadsheet struct {
	vecty.Core

	rerender bool

	spreadsheet.State
}

// NewSpreadsheet creates a new state for the spreadsheet
func NewSpreadsheet() *Spreadsheet {
	s := &Spreadsheet{
		State: spreadsheet.Init(),
	}

	dispatcher.Register(s.update)

	return s
}

func (s *Spreadsheet) update(action interface{}) {

	switch a := action.(type) {
	case *actions.StartEdit:
		s.Active = a.Position
	case *actions.UpdateValue:
		s.Cells[a.Position] = a.Value
	}
}

func (s *Spreadsheet) renderHeader() *vecty.HTML {
	columns := vecty.List{}

	for _, col := range s.Columns {
		columns = append(columns, elem.TableHeader(vecty.Text(col)))
	}

	return elem.TableRow(
		elem.TableHeader(),
		columns,
	)
}

func (s *Spreadsheet) renderEditor(position spreadsheet.Position, value string) *vecty.HTML {
	return elem.TableData(
		vecty.Markup(
			vecty.Style("padding", "0"),
			vecty.Style("height", "30px"),
		),
		elem.Input(
			vecty.Markup(
				vecty.Style("min-width", "0"),
				vecty.Style("width", "100%"),
				vecty.Style("height", "100%"),
				prop.Value(value),
				prop.Type(prop.TypeText),
				event.Input(func(e *vecty.Event) {
					value := e.Target.Get("value").String()

					dispatcher.Dispatch(&actions.UpdateValue{
						Position: position,
						Value:    value,
					})
				}),
			),
		),
	)
}

func (s *Spreadsheet) renderView(position spreadsheet.Position, value string) *vecty.HTML {
	return elem.TableData(
		vecty.Markup(
			event.Click(func(e *vecty.Event) {
				dispatcher.Dispatch(&actions.StartEdit{
					Position: position,
				})
			}),
		),
		vecty.Text(value),
	)
}

func (s *Spreadsheet) renderCell(position spreadsheet.Position) *vecty.HTML {

	value, ok := s.Cells[position]
	if !ok {
		value = ""
	}

	if s.Active.Column == position.Column && s.Active.Row == position.Row {
		log.Println("rendering editor")
		return s.renderEditor(position, value)
	} else {
		return s.renderView(position, value)
	}
}

func (s *Spreadsheet) renderRows() *vecty.HTML {

	rows := vecty.List{}

	for _, row := range s.Rows {
		columns := vecty.List{}

		for _, col := range s.Columns {
			columns = append(columns, s.renderCell(spreadsheet.Position{Column: col, Row: row}))
		}

		rows = append(rows, elem.TableRow(
			elem.TableHeader(vecty.Text(fmt.Sprintf("%d", row))),
			columns,
		))
	}

	return elem.TableBody(
		rows,
	)
}

// Render implements the component interface
func (s *Spreadsheet) Render() vecty.ComponentOrHTML {
	log.Println("render")
	return elem.Body(
		elem.Table(
			vecty.Markup(
				vecty.Class("table"),
				vecty.Class("is-bordered"),
				vecty.Class("is-fullwidth"),
			),
			s.renderHeader(),
			s.renderRows(),
		),
	)
}
