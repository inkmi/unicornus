package pkg

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
	"strings"
)

// Generate field descriptions for HTML fields
// from annotated structs

type DataField struct {
	Name    string
	Kind    string
	SubKind string
	// https://stackoverflow.com/questions/3518002/how-can-i-set-the-default-value-for-an-html-select-element
	Choices      []Choice
	Multi        bool
	Value        any
	Validation   string
	Optional     bool
	ErrorMessage string
	// Replace by Map, every validator might have it's own message
	ErrorMessages []string
}

func (f DataField) HasError() bool {
	return len(f.ErrorMessages) > 0
}

func (f DataField) Errors() string {
	return strings.Join(f.ErrorMessages, ".")
}

func (f DataField) ViewVal() string {
	if f.Value == nil {
		return "-"
	}
	v, ok := f.Value.(int64)
	if ok && v == 0 {
		return "-"
	}
	if f.Choices != nil {
		return fmt.Sprintf("%s", f.choiceLabel())
	} else {
		if f.Kind == "int" {
			return fmt.Sprintf("%d", f.Val())
		} else {
			return fmt.Sprintf("%s", f.Val())
		}
	}
}

func (f DataField) choiceLabel() any {
	if f.Choices == nil {
		return f.Val()
	} else {
		var returnValue any
		if f.Value != nil {
			for _, choice := range f.Choices {
				if f.Kind == "int" {
					if choice.Value == strconv.FormatInt(f.Value.(int64), 10) {
						returnValue = choice.Label
					}
				} else if f.Kind == "string" {
					if choice.Value == f.Value {
						returnValue = choice.Label
					}
				}
			}
		} else {
			returnValue = ""
		}
		return sanitize(returnValue)
	}
}

func (f DataField) Val() any {
	var returnValue any
	if f.Value != nil {
		if f.Choices != nil {
			for _, v := range f.Choices {
				if f.Kind == "int" {
					if v.Value == strconv.FormatInt(f.Value.(int64), 10) {
						returnValue = v.Label
						// BUG: ??
					}
				} else if f.Kind == "string" {
					if v.Value == f.Value {
						returnValue = v.Label
						// BUG: ?? BREAk?
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
