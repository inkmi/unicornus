package main

// S:1
import (
	// E:1
	"fmt"
	"net/http"
	// S:1
	"github.com/inkmi/unicornus/uni"
)

type simpledata struct {
	Name string
}

// E:1

func example1(w http.ResponseWriter, req *http.Request) {
	// S:1
	// The data of the form
	d := simpledata{
		Name: "Unicornus",
	}
	// Create a FormLayout
	// describing the form
	ui := uni.NewFormLayout().
		Add("Name", "Name Label")

	// Render form layout with data
	// to html
	html := ui.RenderForm(d)
	// E:1
	fmt.Fprintf(w, html)
}
