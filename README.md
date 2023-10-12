<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/UnicornusLogo.png" width="300">

# Unicornus: Easy Form Generation in Go

This is v0.01 and not useful at all.

![Build state](https://github.com/inkmi/unicornus/actions/workflows/test.yml/badge.svg)  ![Go Version](https://img.shields.io/github/go-mod/go-version/inkmi/unicornus) ![Version](https://img.shields.io/github/v/tag/inkmi/unicornus?include_prereleases)  ![Issues](https://img.shields.io/github/issues/inkmi/unicornus) ![Report](https://goreportcard.com/badge/github.com/inkmi/unicornus)

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-83%25-brightgreen.svg?longCache=true&style=flat)</a>

Unicornus is a simple library for building HTML forms in Go.

## Get started

### Install

```shell
go get github.com/inkmi/unicornus
```

## Idea

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

## License information

MIT License

## Contributors

### Major

- [Stephan Schmidt](https://github.com/StephanSchmidt): Author and maintainer
