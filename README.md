<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/UnicornusLogo.png" width="300">

# Unicornus: Easy Form Generation in Go

**ALPHA - play, don't use**

![Build state](https://github.com/inkmi/unicornus/actions/workflows/test.yml/badge.svg)  ![Go Version](https://img.shields.io/github/go-mod/go-version/inkmi/unicornus) ![Version](https://img.shields.io/github/v/tag/inkmi/unicornus?include_prereleases)  ![Issues](https://img.shields.io/github/issues/inkmi/unicornus) ![Report](https://goreportcard.com/badge/github.com/inkmi/unicornus)

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-87%25-brightgreen.svg?longCache=true&style=flat)</a>

Unicornus is a simple library for building HTML forms in Go.

## Get started

### Install

```shell
go get github.com/inkmi/unicornus
```

## Examples

You can run the examples with

```
make examples
```


## License information

MIT License

## Contributors

### Major

- [Stephan Schmidt](https://github.com/StephanSchmidt): Author and maintainer

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


Unicornus can render form with errors. Errors are `map[string]string` and contain the field which created the error and an error text. These errors together with the data are rendered with `RenderFormWithErrors`.

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

