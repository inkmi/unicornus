package main

// S:1
import (
	"fmt"
	uni "github.com/inkmi/unicornus/pkg"
	"net/http"
)

type errorexample struct {
	Name string
}

// E:1

func example2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><body><div>")

	// S:1
	// The data of the form
	d := errorexample{
		Name: "Unicornus",
	}
	// Create a FormLayout
	// describing the form
	ui := uni.NewFormLayout().
		Add("Name", "Name")

	// Errors are a map of string -> string
	// with field names and error texts
	errors := map[string]string{"Name": "Name can't be Unicornus"}

	// Render form layout with data
	// to html
	html := ui.RenderFormWithErrors(d, errors)
	// E:1

	fmt.Fprintf(w, html)

	fmt.Fprintf(w, "</div></body></html>")

}
