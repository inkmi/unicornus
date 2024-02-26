package main

// S:1
import (
	// E:1
	"fmt"
	"net/http"
	// S:1
	uni "github.com/inkmi/unicornus/pkg"
)

type subData3 struct {
	Sub string
}
type data3 struct {
	Name string
	Sub  subData3
}

// E:1

func example3(w http.ResponseWriter, req *http.Request) {
	// S:1
	// The data of the form
	d := data3{
		Name: "Unicornus",
		Sub: subData3{
			Sub: "Ha",
		},
	}

	// Create a FormLayout
	// describing the form
	ui := uni.NewFormLayout().
		Add("Name", "Name Label", uni.WithDescription("Name Description")).
		AddGroup("Sub", "Group", "Group Description", func(f *uni.FormLayout) {
			f.
				Add("Sub", "Sub Label")
		})

	// Render form layout with data
	// to html
	html := ui.RenderForm(d)
	// E:1
	fmt.Fprintf(w, html)
}
