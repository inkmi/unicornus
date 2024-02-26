# Documentation

The idea of Unicornus is

* define a data model in Go with structs
* add constraints for validation tags
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
