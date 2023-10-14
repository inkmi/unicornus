package pkg

import (
	"strconv"
)

// Generate field descriptions for HTML fields
// from annotated structs

type DataField struct {
	Name string
	Kind string
	// https://stackoverflow.com/questions/3518002/how-can-i-set-the-default-value-for-an-html-select-element
	Choices    []Choice
	Multi      bool
	Value      any
	Validation string
	Optional   bool
	// Replace by Map, every validator might have it's own message
	ErrorMessage string
}

func (f DataField) Val() any {
	if f.Value != nil {
		if f.Choices != nil {
			for _, v := range f.Choices {
				if f.Kind == "int" {
					if v.Value == strconv.FormatInt(f.Value.(int64), 10) {
						return v.Label
					}
				} else if f.Kind == "string" {
					if v.Value == f.Value {
						return v.Label
					}
				}
			}
			return f.Value
		} else {
			if f.Optional {
				return f.Value
			} else {
				return f.Value
			}
		}
	} else {
		return ""
	}
}
