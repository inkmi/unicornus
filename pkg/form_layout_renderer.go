package pkg

import (
	"fmt"
	"strings"
)

type RenderContext struct {
	out          *strings.Builder
	AnchorGroups bool
	DisplayMode  bool
}

func (rctx *RenderContext) OnlyDisplay(name string) bool {
	return rctx.DisplayMode
}

type RenderContextFunc func(ctx *RenderContext)

func defaultContext() *RenderContext {
	var sb strings.Builder
	rctx := RenderContext{
		out:          &sb,
		AnchorGroups: true,
	}
	return &rctx
}

func WithAnchorGroups(anchorGroups bool) RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.AnchorGroups = anchorGroups
	}
}

func WithDisplayMode() RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.DisplayMode = true
	}
}

func NewRenderContext(config ...RenderContextFunc) *RenderContext {
	c := defaultContext()
	for _, con := range config {
		con(c)
	}
	return c
}

func (f *FormLayout) RenderView(data any) string {
	errors := make(map[string]string)
	m := FieldsToMap(FieldGenerator(data, errors))
	rctx := NewRenderContext(WithDisplayMode())
	f.renderFormToBuilder(rctx, "", m)
	return rctx.out.String()
}

func (f *FormLayout) RenderForm(data any) string {
	errors := make(map[string]string)
	return f.RenderFormWithErrors(data, errors)
}

func (f *FormLayout) RenderFormWithErrors(data any, errors map[string]string) string {
	m := FieldsToMap(FieldGenerator(data, errors))
	rctx := NewRenderContext()
	f.renderFormToBuilder(rctx, "", m)
	return rctx.out.String()
}

func (f *FormLayout) RenderElementWithErrors(
	name string,
	data any,
	errors map[string]string,
) string {
	e := f.findByName(name)
	m := FieldsToMap(FieldGenerator(data, errors))
	var sb strings.Builder
	rctx := RenderContext{
		out: &sb,
	}
	// probably prefix taken from name instead of ""
	f.renderElement(*e, &rctx, "", m)
	return sb.String()
}

func (f *FormLayout) renderFormToBuilder(rctx *RenderContext, prefix string, m map[string]DataField,
) {
	for _, e := range f.elements {
		f.renderElement(e, rctx, prefix, m)
	}
}

func (f *FormLayout) renderElement(
	e FormElement,
	rctx *RenderContext,
	prefix string,
	m map[string]DataField,
) {
	switch e.Kind {
	case "hidden":
		fieldName := e.Name
		if len(prefix) > 0 {
			fieldName = prefix + "." + fieldName
		}
		field, ok := m[fieldName]
		if ok {
			rctx.out.WriteString(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"%v\" />", fieldName, field.Val()))
		}
	case "header":
		f.Theme.themeRenderHeader(rctx, e)
	case "group":
		newPrefix := e.Name
		if len(prefix) > 0 {
			newPrefix = prefix + "." + newPrefix
		}
		if rctx.AnchorGroups {
			rctx.out.WriteString(fmt.Sprintf("<a name=\"formgroup-%s\"></a>", stringToAnchor(e.Config.Label)))
		}
		f.Theme.themeRenderGroup(rctx, m, newPrefix, e)
	case "input":
		// take value string from MAP of name -> DataField
		// take type if no type is given from DataField
		fieldName := e.Name
		if len(prefix) > 0 {
			fieldName = prefix + "." + fieldName
		}
		field, ok := m[fieldName]

		if ok {
			if len(e.Config.Choices) > 0 {
				field.Choices = e.Config.Choices
			}
			if field.Multi {
				values := field.Value.([]string)
				for i := 0; i < len(field.Choices); i++ {
					choice := &field.Choices[i]
					if containsString(values, choice.Value) {
						choice.Checked = true
					}
				}
				f.Theme.themeRenderMulti(rctx, field, e, prefix)
			} else {
				description := e.Config.Description
				if len(e.Config.Description) > 0 {
					description = e.Config.Description
				}
				if field.Kind == "bool" {
					f.Theme.themeRenderCheckbox(rctx, e, field, description, prefix)
				} else if !field.Multi && len(field.Choices) > 0 {
					f.Theme.themeRenderSelect(rctx, e, field, description, prefix)
				} else {
					f.Theme.themeRenderInput(rctx, e, field, prefix)
				}
			}
		}
	case "yesno":
		// take value string from MAP of name -> DataField
		// take type if no type is given from DataField
		fieldName := e.Name
		if len(prefix) > 0 {
			fieldName = prefix + "." + fieldName
		}
		field, ok := m[fieldName]

		if ok {
			if len(e.Config.Choices) > 0 {
				field.Choices = e.Config.Choices
			}
			description := e.Config.Description
			if len(e.Config.Description) > 0 {
				description = e.Config.Description
			}
			f.Theme.themeRenderYesNo(rctx, e, field, description, prefix)
		}
	case "dropdown":
		// take value string from MAP of name -> DataField
		// take type if no type is given from DataField
		fieldName := e.Name
		if len(prefix) > 0 {
			fieldName = prefix + "." + fieldName
		}
		field, ok := m[fieldName]
		if ok {
			if len(e.Config.Choices) > 0 {
				field.Choices = e.Config.Choices
			}
			//if field.Multi {
			//	values := field.Value.([]string)
			//	for i := 0; i < len(field.Choices); i++ {
			//		choice := &field.Choices[i]
			//		if containsString(values, choice.Value) {
			//			choice.Checked = true
			//		}
			//	}
			//	f.Theme.themeRenderMulti(sb, field, e, prefix)
			//} else {
			f.Theme.themeRenderSelect(rctx, e, field, e.Config.Description, prefix)
			//}
		}
	}
}

func renderCheckbox(rctx *RenderContext, f DataField, config ElementOpts, prefix string, class string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := f.Name
	rctx.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"%s\" %s%s/>", name, class, checked, configToHtml(config)))
}

func renderSelect(rctx *RenderContext, f DataField, config ElementOpts, prefix string, class string, e FormElement) {
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	// optgroup https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
	if len(e.Config.Groups) > 0 {
		rctx.out.WriteString(fmt.Sprintf("<select name=\"%s\" class=\"%s\"><option value=\"0\">-</option>", name, class))
		for group, name := range e.Config.Groups {
			rctx.out.WriteString(fmt.Sprintf("<optgroup label=\"%s\">", name))
			for _, c := range f.Choices {
				if len(group) == 0 || c.Group == group {
					if c.IsSelected(f.Value) {
						rctx.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))
					} else {
						rctx.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
					}
				}
			}
			rctx.out.WriteString("</optgroup>")
		}
		rctx.out.WriteString("</select>")
	} else {
		rctx.out.WriteString(fmt.Sprintf("<select name=\"%s\" class=\"%s\"><option value=\"0\">-</option>", name, class))
		for _, c := range f.Choices {
			if c.IsSelected(f.Value) {
				rctx.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))
			} else {
				rctx.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
			}
		}
		rctx.out.WriteString("</select>")
	}
}

func renderTextInput(rctx *RenderContext, f DataField, val any, config ElementOpts, prefix string, class string) {
	inputConstraints := ""

	inputType := "text"
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	if f.SubKind == "email" {
		inputType = "email"
	} else if f.SubKind == "url" {
		inputType = "url"
	}

	//if f.Name == "Name" {
	//	sb.WriteString(fmt.Sprintf("<div class=\"field_%s\" hx-target-422=\"this\" hx-select=\".field_%s\" hx-target=\"this\" hx-swap=\"outerHTML\" hx-trigger=\"change\" hx-post=\"/ob/onboarding-dreamer?validate=yes\">\n", f.Name, f.Name))
	//	sb.WriteString(fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s value=\"%v\"%s class=\"%s\"/>",
	//		name, inputType,
	//		strings.TrimSpace(inputConstraints), val, configToHtml(config), class))
	//	if f.HasError() {
	//		sb.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", f.Errors()))
	//	}
	//	sb.WriteString("</div>")
	//
	//} else {
	rctx.out.WriteString(fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s value=\"%v\"%s class=\"%s\"/>",
		name, inputType,
		strings.TrimSpace(inputConstraints), val, configToHtml(config), class))
	if f.HasError() {
		rctx.out.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", f.Errors()))
		//}
	}

}

func configToHtml(config ElementOpts) string {
	id := ""
	if len(config.Id) > 0 {
		id = fmt.Sprintf(" id=\"%s\"", config.Id)
	}
	placeholder := ""
	if len(config.Placeholder) > 0 {
		placeholder = fmt.Sprintf(" placeholder=\"%s\"", config.Placeholder)
	}
	configStr := fmt.Sprintf("%s%s", id, placeholder)
	return configStr
}

func containsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

/*

func SetChoices(setKey string, fields []FieldV, allValues []string) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, Choice{
					Label:    p,
					Value:    p,
					Selected: lo.Contains(values, p),
				})
			}

			fields[i].Choices = choices
			fields[i].Kind = "string"
		}
	}
}

func SetKey(
	setKey string,
	fields []FieldV,
	allValues []string,
	group func(k string) string,
	label func(l string) string,
) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, Choice{
					Group:    group(p),
					Label:    label(p),
					Value:    p,
					Selected: lo.Contains(values, p),
				})
			}

			fields[i].Choices = choices
			fields[i].Kind = "string"
		}
	}
}

*/
