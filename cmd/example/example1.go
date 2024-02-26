package main

// S:1
import (
	"fmt"
	uni "github.com/inkmi/unicornus/pkg"
	// E:1
	"net/http"
	// S:1
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
	fmt.Fprintf(w, html)
	// E:1
}
