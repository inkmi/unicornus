package main

import (
	"fmt"
	uni "github.com/inkmi/unicornus/pkg"
	"net/http"
)

type data2 struct {
	Name string
}

func example2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><body><div>")

	// The data of the form
	d := data2{
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
	fmt.Fprintf(w, html)

	fmt.Fprintf(w, "</div></body></html>")

}
