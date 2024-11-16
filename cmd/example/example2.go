package main

// S:1
import (
	// E:1
	"net/http"
	// S:1
	// "github.com/inkmi/unicornus/uni"
)

type errorexample struct {
	Name string
}

// E:1

func example2(w http.ResponseWriter, req *http.Request) {
	// // S:1
	// // The data of the form
	// d := errorexample{
	// 	Name: "Unicornus",
	// }
	// // Create a FormLayout
	// // describing the form
	// ui := uni.NewFormLayout().
	// 	Add("Name", "Name")

	// // Errors are a map of string -> string
	// // with field names and error texts
	// errors := map[string]string{"Name": "Name can't be Unicornus"}

	// // Render form layout with data
	// // to html
	// html := ui.RenderFormWithErrors(d, errors)
	// // E:1

	// fmt.Fprintf(w, html)
}
