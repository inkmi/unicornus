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


## Idea

The idea of Unicornus is to combine a data model in Go described as structs with validation tags
with a description of the form layout in Go to render an HTML form.


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


## Code Example

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
See [cmd/example/example2.go](cmd/example/example2.go)

Results in

<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/formexample.png" width="600">



