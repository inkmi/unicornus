package main

import (
	"fmt"
	uni "github.com/inkmi/unicornus/pkg"
	"net/http"
)

type subData3 struct {
	Sub string
}
type data3 struct {
	Name   string
	Check  bool
	Select int      `validate:"int|in:1,2,3" choices:"A|B|C"`
	Multi  []string `choices:"A|B|C"`
	Sub    subData3
}

func example3(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><body>")

	// The data of the form
	d := data3{
		Name:   "Unicornus",
		Check:  true,
		Select: 2,
		Multi:  []string{"C"},
		Sub: subData3{
			Sub: "Ha",
		},
	}

	// Create a FormLayout
	// describing the form
	ui := uni.NewFormLayout().
		AddHeader("Form").
		Add("Name", "Name Label", uni.WithDescription("Name Description")).
		Add("Check", "Check Label").
		Add("Select", "Select Label").
		Add("Multi", "Multi Label").
		AddGroup("Sub", "Group", "Group Description", func(f *uni.FormLayout) {
			f.
				Add("Sub", "Sub Label")
		})

	// Render form layout with data
	// to html
	html := ui.RenderForm(d)
	fmt.Fprintf(w, html)

	fmt.Fprintf(w, "</body></html>")

}
