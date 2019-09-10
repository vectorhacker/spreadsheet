package main

import (
	"spreadsheet/pkg/components"

	"spreadsheet/pkg/dispatcher"

	"github.com/gopherjs/vecty"
)

func main() {

	vecty.AddStylesheet("https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.5/css/bulma.css")

	s := components.NewSpreadsheet()

	dispatcher.Register(func(_ interface{}) {
		vecty.Rerender(s)
	})
	vecty.RenderBody(s)
}
