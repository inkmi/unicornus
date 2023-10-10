package builder

import (
	"github.com/Inkmi/unicornus/internal/ui"
)

type ElementConfig struct {
	Placeholder string
	Id          string
	Label       string
	Choices     []ui.Choice
	Groups      []string
	SubLayout   *FormLayout
}

type FormLayout struct {
	elements []FormElement
}

type FormElement struct {
	Kind   string
	Name   string
	Config ElementConfig
}

func NewFormLayout() *FormLayout {
	return new(FormLayout)
}

func (f *FormLayout) AddHeader(name string) *FormLayout {
	e := FormElement{
		Kind: "header",
		Name: name,
	}
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) AddGroup(name string, layout func(f *FormLayout)) *FormLayout {
	l := NewFormLayout()
	e := FormElement{
		Kind: "group",
		Name: name,
		Config: ElementConfig{
			SubLayout: l,
		},
	}
	layout(l)
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) Add(name string, label string, config ...ElementConfig) *FormLayout {
	var c ElementConfig
	if len(config) > 0 {
		c = config[0]
	} else {
		c = ElementConfig{}
	}
	if len(c.Label) == 0 {
		c.Label = label
	}
	e := FormElement{
		Kind:   "input",
		Name:   name,
		Config: c,
	}
	f.elements = append(f.elements, e)
	return f
}

//func containsString(slice []string, target string) bool {
//	for _, s := range slice {
//		if s == target {
//			return true
//		}
//	}
//	return false
//}
