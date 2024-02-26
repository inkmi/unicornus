# Documentation

## Idea

The idea of Unicornus is to combine a data model in Go described as structs with validation tags and a description of the form layout in Go to render an HTML form.


```
           ┌─────────────┐             ┌─────────────┐
           │             │             │             │
Validation │ Data Model  ├──────┬──────┤ Form Layout │
           │             │      │      │             │
           │             │      │      │             │
           └─────────────┘      │      └─────────────┘
                  ▲             │
                  │             │
                  │             │
                  │             │
                  │             │
                  │             │
                  │             ▼
                  │       ┌───────────┐
                  │       │           │
           Submit │       │           │
                  │       │   HTML    │
                  └───────┤   Form    │
                          │           │
                          │           │
                          └───────────┘
```

# Code Examples



## Simple Example


Creating a form in Unicornus is very simple. You defined the data structure and
then the form layout. Then you can simply call `RenderForm` with the data on the form layout to create html.

```go
import (
  "fmt"
  uni "github.com/inkmi/unicornus/pkg"
  "net/http"
)

type simpledata struct {
  Name string
}

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
```
From [cmd/example/example1.go](cmd/example/example1.go)


## Displaying Errors


```go
import (
  "fmt"
  uni "github.com/inkmi/unicornus/pkg"
  "net/http"
)

type errorexample struct {
  Name string
}

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
```
From [cmd/example/example2.go](cmd/example/example2.go)

Results in

<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/formexample.png" width="600">

