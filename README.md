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

### Examples

You can run the examples with

```
make examples
```

## License information

MIT License

## Contributors

- [Stephan Schmidt](https://github.com/StephanSchmidt): Author and maintainer

# Documentation

<p><b>Table Of Contents</b></p><ul><li><a href="#code-examples">Code Examples</a></li><ul><li><a href="#simple-example">Simple Example</a></li><li><a href="#displaying-errors">Displaying Errors</a></li><li><a href="#nested-data">Nested Data</a></li></ul></ul>

The idea of Unicornus is

* define a data model in Go with structs
* add constraints with validation tags
* define the layout in Go
* render data to HTML with the defined layout


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
An embedded struct is best layouted with `AddGroup`. The name of the group is the name of the
embedded struct, in this case `Sub`. The label of the group is displayed as a header, the
description of the group is displayed for explanation.

The `AddGroup` is given a function `func(f *uni.FormLayout)`. Inside this function (kind of a callback)
the layout of the group is defined. The root is the sub struct of the group.


```go
import (
  uni "github.com/inkmi/unicornus/pkg"
)

type subData3 struct {
  SubName string
}
type data3 struct {
  Name string
  Sub  subData3
}

// The data of the form
d := data3{
  Name: "Unicornus",
  Sub: subData3{
    SubName: "Ha my name!",
  },
}

// Create a FormLayout
// describing the form
ui := uni.NewFormLayout().
Add("Name", "Name Label").
AddGroup("Sub", "Group", "Group Description", func(f *uni.FormLayout) {
  f.
  Add("SubName", "Sub Label")
})

// Render form layout with data
// to html
html := ui.RenderForm(d)
```
From [cmd/example/example3.go](cmd/example/example3.go)

The names of the fields in the HTML forms are dot se