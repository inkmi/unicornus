package pkg

import (
	"github.com/microcosm-cc/bluemonday"
	"strconv"
)

// Generate field descriptions for HTML fields
// from annotated structs

type DataField struct {
	Name    string
	Kind    string
	SubKind string
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
	var returnValue any
	if f.Value != nil {
		if f.Choices != nil {
			for _, v := range f.Choices {
				if f.Kind == "int" {
					if v.Value == strconv.FormatInt(f.Value.(int64), 10) {
						returnValue = v.Label
					}
				} else if f.Kind == "string" {
					if v.Value == f.Value {
						returnValue = v.Label
					}
				}
			}
			returnValue = f.Value
		} else {
			if f.Kind == "int" && f.Value.(int64) == 0 {
				return ""
			}
			if f.Optional {
				returnValue = f.Value
			} else {
				returnValue = f.Value
			}
		}
	} else {
		returnValue = ""
	}
	return sanitize(returnValue)
}

func sanitize(value any) any {
	p := bluemonday.StrictPolicy()
	switch v := value.(type) {
	case string:
		return p.Sanitize(v)
	default:
		return v
	}
}
