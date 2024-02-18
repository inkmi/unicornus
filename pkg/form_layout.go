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
	ViewOnly    bool
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
	Theme    BaseTheme
	elements []FormElement
}

type FormElement struct {
	Kind   string
	Name   string
	Config ElementOpts
}

func NewFormLayout() *FormLayout {
	fl := &FormLayout{
		Theme: BaseTheme{
			// TopSeparator
			NewStyles(
				InputStyle("display: block; width: 100%%; margin-top: 0.25rem; border: 1px solid #D1D5DB; border-radius: 0.375rem; box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05); padding-top: 0.5rem; padding-bottom: 0.5rem; padding-left: 0.75rem; padding-right: 0.75rem; outline: none;"),
				LabelStyle("display: block; font-size: 0.875rem; font-weight: 500; color: #6B7280"),
				ErrorStyle("margin-top: 0.5rem; font-size: 0.875rem; color: #e3342f;"),
				TopSeparator("1.5rem"),
			),
		},
	}
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

func (f *FormLayout) AddViewOnlyGroup(name string,
	label string,
	description string,
	layout func(f *FormLayout),
) *FormLayout {
	l := NewFormLayout()
	e := FormElement{
		Kind: "group",
		Name: name,
		Config: ElementOpts{
			SubLayout:   l,
			Label:       label,
			Description: description,
			ViewOnly:    true,
		},
	}
	layout(l)
	f.elements = append(f.elements, e)
	return f
}

func (f *FormLayout) AddGroup(name string,
	label string,
	description string,
	layout func(f *FormLayout),
) *FormLayout {
	l := NewFormLayout()
	e := FormElement{
		Kind: "group",
		Name: name,
		Config: ElementOpts{
			SubLayout:   l,
			Label:       label,
			Description: description,
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
