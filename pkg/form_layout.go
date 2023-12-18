package pkg

import "strings"

type ElementOpts struct {
	Placeholder string
	Id          string
	Label       string
	Description string
	Choices     []Choice
	Groups      map[string]string
	SubLayout   *FormLayout
}

type OptFunc func(config *ElementOpts)

func defaultOpts() *ElementOpts {
	return &ElementOpts{}
}

func WithDescription(description string) OptFunc {
	return func(config *ElementOpts) {
		config.Description = description
	}
}

func WithChoices(choices []Choice) OptFunc {
	return func(config *ElementOpts) {
		config.Choices = choices
	}
}

func WithPlaceholder(placeholder string) OptFunc {
	return func(config *ElementOpts) {
		config.Placeholder = placeholder
	}
}

func WithId(id string) OptFunc {
	return func(config *ElementOpts) {
		config.Id = id
	}
}

func WithGroups(groups map[string]string) OptFunc {
	return func(config *ElementOpts) {
		config.Groups = groups
	}
}

type FormLayout struct {
	Theme    TailwindTheme
	elements []FormElement
}

type FormElement struct {
	Kind        string
	Name        string
	Label       string
	Description string
	Config      ElementOpts
}

func NewFormLayout() *FormLayout {
	fl := &FormLayout{}
	return fl
}

func (f *FormLayout) AddHeader(name string) *FormLayout {
	e := FormElement{
		Kind: "header",
		Name: name,
	}
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) AddHidden(name string) *FormLayout {
	e := FormElement{
		Kind: "hidden",
		Name: name,
	}
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) findByName(name string) *FormElement {
	if f == nil {
		return nil
	}
	// Split the name by the first dot to handle nested group names
	parts := strings.SplitN(name, ".", 2)
	for _, element := range f.elements {
		// Check if the current element matches the first part of the name
		if element.Name == parts[0] {
			// If it's a direct match or not a group, return the element
			if len(parts) == 1 || element.Kind != "group" {
				return &element
			}
			// If it's a group and there are more parts, search in the SubLayout
			if element.Config.SubLayout != nil {
				return element.Config.SubLayout.findByName(parts[1])
			}
		}
	}
	return nil
}

func (f *FormLayout) AddGroup(name string,
	label string,
	description string,
	layout func(f *FormLayout),
) *FormLayout {
	l := NewFormLayout()
	e := FormElement{
		Kind:        "group",
		Name:        name,
		Label:       label,
		Description: description,
		Config: ElementOpts{
			SubLayout: l,
		},
	}
	layout(l)
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) Add(name string, label string, config ...OptFunc) *FormLayout {
	c := defaultOpts()
	for _, con := range config {
		con(c)
	}

	if len(c.Label) == 0 {
		c.Label = label
	}
	e := FormElement{
		Kind:   "input",
		Name:   name,
		Config: *c,
	}
	f.elements = append(f.elements, e)
	return f
}

// Support multi select dropdowns
// https://tw-elements.com/docs/standard/forms/select/''
func (f *FormLayout) AddDropdown(name string, label string, config ...OptFunc) *FormLayout {
	c := defaultOpts()
	for _, con := range config {
		con(c)
	}
	if len(c.Label) == 0 {
		c.Label = label
	}
	e := FormElement{
		Kind:   "dropdown",
		Name:   name,
		Config: *c,
	}
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) AddYesNo(name string, label string, config ...OptFunc) *FormLayout {
	c := defaultOpts()
	for _, con := range config {
		con(c)
	}
	if len(c.Label) == 0 {
		c.Label = label
	}
	e := FormElement{
		Kind:   "yesno",
		Name:   name,
		Config: *c,
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
