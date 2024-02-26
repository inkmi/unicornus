# Documentation

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


Creating a form in Unicornus is very simple. You define the data structure and
then the form layout. Then you can simply call `RenderForm` with the data on the form layout to create HTML.

```go
import (
  uni "github.com/inkmi/unicornus/pkg"
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
```
From [cmd/example/example1.go](cmd/example/example1.go)


## Displaying Errors


Unicornus can render form with errors. Errors are `map[string]string` and contain the field which created the error and an error text. These errors together with the data are rendered with `RenderFormWithErrors`.

```go
import (
  uni "github.com/inkmi/unicornus/pkg"
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

## Nested Data


Data in Unicornus can be nested. A struct can have sub structs and those are rendered into HTML.

```go
import (
  uni "github.com/inkmi/unicornus/pkg"
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
```
From [cmd/example/example3.go](cmd/example/example3.go)


