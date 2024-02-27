package main

// S:1
import (
	// E:1
	"fmt"
	"net/http"

	// S:1
	"github.com/inkmi/unicornus/uni"
)

type subData4 struct {
	Sub string
}
type data4 struct {
	Name   string
	Check  bool
	Select int      `validate:"int|in:1,2,3" choices:"A|B|C"`
	Multi  []string `choices:"A|B|C"`
	Sub    subData4
}

// E:1

func example4(w http.ResponseWriter, req *http.Request) {
	// S:1
	// The data of the form
	d := data4{
		Name:   "Unicornus",
		Check:  true,
		Select: 2,
		Multi:  []string{"C"},
		Sub: subData4{
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
	// E:1
	fmt.Fprintf(w, html)
}
