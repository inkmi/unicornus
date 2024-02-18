package pkg

import (
	"fmt"
	"html"
	"strings"
)

type RenderContext struct {
	out          *strings.Builder
	AnchorGroups bool
	ViewMode     bool
}

func Safe(str string) string {
	return html.EscapeString(str)
}

func WriteString(s *strings.Builder, str string) {
	s.WriteString(Safe(str))
}

func WriteValue(s *strings.Builder, value any) {
	WriteString(s, fmt.Sprintf("%v", value))
}

func (r *RenderContext) INPUT(name string, typ string, value any, style string, config ElementOpts) {
	r.out.WriteString("<input name=\"")
	r.out.WriteString(name)
	r.out.WriteString("\" type=\"")
	r.out.WriteString(typ)
	r.out.WriteString("\" value=\"")
	WriteValue(r.out, value)
	r.out.WriteString("\" ")
	r.out.WriteString(configToHtml(config))
	r.out.WriteString(" style=\"")
	r.out.WriteString(style)
	r.out.WriteString("\">")
}

func (r *RenderContext) DIVv(content string, class ...string) {
	r.out.WriteString("<div class=\"")
	for i := 0; i < len(class); i++ {
		r.out.WriteString(class[i])
		r.out.WriteString(" ")
	}
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</div>")
}

func (r *RenderContext) H2no(content string) {
	r.out.WriteString("<h2>")
	r.out.WriteString(content)
	r.out.WriteString("</h2>")
}

func (r *RenderContext) DIVclose() {
	r.out.WriteString("</div>")
}

func (r *RenderContext) H2(content string, class string) {
	r.out.WriteString("<h2 class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</h2>")
}

func (r *RenderContext) H3(content string, class string) {
	r.out.WriteString("<h3 class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</h3>")
}

func (r *RenderContext) DIV(content string, class string) {
	r.out.WriteString("<div class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</div>")
}

func (r *RenderContext) DIVS(content string, style string) {
	r.out.WriteString("<div style=\"")
	r.out.WriteString(style)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</div>")
}

func (r *RenderContext) DIVopen(class string) {
	r.out.WriteString("<div class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
}
func (r *RenderContext) DIVopenS(style string) {
	r.out.WriteString("<div style=\"")
	r.out.WriteString(style)
	r.out.WriteString("\">")
}

func (r *RenderContext) LABEL(content string, class string) {
	r.out.WriteString("<label class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</label>")
}

func (r *RenderContext) LABELS(content string, style string) {
	r.out.WriteString("<label style=\"")
	r.out.WriteString(style)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</label>")
}

func (r *RenderContext) PS(content string, style string) {
	r.out.WriteString("<p style=\"")
	r.out.WriteString(style)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</p>")
}

func (r *RenderContext) p(content string, class string) {
	r.out.WriteString("<p class=\"")
	r.out.WriteString(class)
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</p>")
}

func (r *RenderContext) OnlyDisplay(name string) bool {
	return r.ViewMode
}

type RenderContextFunc func(ctx *RenderContext)

func defaultContext() *RenderContext {
	var sb strings.Builder
	r := RenderContext{
		out:          &sb,
		AnchorGroups: true,
	}
	return &r
}

func WithAnchorGroups(anchorGroups bool) RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.AnchorGroups = anchorGroups
	}
}

func WithDisplayMode() RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.ViewMode = true
	}
}

func NewRenderContext(config ...RenderContextFunc) *RenderContext {
	c := defaultContext()
	for _, con := range config {
		con(c)
	}
	return c
}

func (f *FormLayout) RenderView(data ...any) string {
	errors := make(map[string]string)
	datas := make([]DataField, 0)
	for _, d := range data {
		// [:]... MAGIC from chatgpt
		datas = append(datas, FieldGenerator(d, errors)[:]...)
	}
	m := FieldsToMap(datas)
	r := NewRenderContext(WithDisplayMode())
	f.renderFormToBuilder(r, "", m)
	return r.out.String()
}

func (f *FormLayout) RenderForm(data any) string {
	errors := make(map[string]string)
	return f.RenderFormWithErrors(data, errors)
}

func (f *FormLayout) RenderFormWithErrors(data any, errors map[string]string) string {
	m := FieldsToMap(FieldGenerator(data, errors))
	r := NewRenderContext()
	f.renderFormToBuilder(r, "", m)
	return r.out.String()
}

func (f *FormLayout) RenderElementWithErrors(
	name string,
	data any,
	errors map[string]string,
) string {
	e := f.findByName(name)
	m := FieldsToMap(FieldGenerator(data, errors))
	var sb strings.Builder
	r := RenderContext{
		out: &sb,
	}
	// probably prefix taken from name instead of ""
	f.renderElement(*e, &r, "", m)
	return sb.String()
}

func (f *FormLayout) renderFormToBuilder(r *RenderContext, prefix string, m map[string]DataField,
) {
	for _, e := range f.elements {
		f.renderElement(e, r, prefix, m)
	}
}

func (f *FormLayout) renderElement(
	e FormElement,
	r *RenderContext,
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
			r.out.WriteString(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"%v\" />", fieldName, field.Val()))
		}
	case "header":
		f.Theme.themeRenderHeader(r, e)
	case "group":
		if (!e.Config.ViewOnly) || r.ViewMode {
			newPrefix := e.Name
			if len(prefix) > 0 {
				newPrefix = prefix + "." + newPrefix
			}
			if r.AnchorGroups {
				r.out.WriteString(fmt.Sprintf("<a name=\"formgroup-%s\"></a>", stringToAnchor(e.Config.Label)))
			}
			f.Theme.themeRenderGroup(r, m, newPrefix, e)
		}
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
				f.Theme.themeRenderMulti(r, field, e, prefix)
			} else {
				description := e.Config.Description
				if len(e.Config.Description) > 0 {
					description = e.Config.Description
				}
				if field.Kind == "bool" {
					f.Theme.themeRenderCheckbox(r, e, field, description, prefix)
				} else if !field.Multi && len(field.Choices) > 0 {
					f.Theme.themeRenderSelect(r, e, field, description, prefix)
				} else {
					f.Theme.themeRenderInput(r, e, field, prefix)
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
			f.Theme.themeRenderYesNo(r, e, field, description, prefix)
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
			f.Theme.themeRenderSelect(r, e, field, e.Config.Description, prefix)
			//}
		}
	}
}

func renderCheckbox(r *RenderContext, f DataField, config ElementOpts, prefix string, class string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := f.Name
	r.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"%s\" %s%s/>", name, class, checked, configToHtml(config)))
}

func renderSelect(r *RenderContext, f DataField, config ElementOpts, prefix string, class string, e FormElement) {
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	// optgroup https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
	if len(e.Config.Groups) > 0 {
		r.out.WriteString(fmt.Sprintf("<select name=\"%s\" class=\"%s\"><option value=\"0\">-</option>", name, class))
		for group, name := range e.Config.Groups {
			r.out.WriteString(fmt.Sprintf("<optgroup LABEL=\"%s\">", name))
			for _, c := range f.Choices {
				if len(group) == 0 || c.Group == group {
					if c.IsSelected(f.Value) {
						r.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))
					} else {
						r.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
					}
				}
			}
			r.out.WriteString("</optgroup>")
		}
		r.out.WriteString("</select>")
	} else {
		r.out.WriteString(fmt.Sprintf("<select name=\"%s\" class=\"%s\"><option value=\"0\">-</option>", name, class))
		for _, c := range f.Choices {
			if c.IsSelected(f.Value) {
				r.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))
			} else {
				r.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
			}
		}
		r.out.WriteString("</select>")
	}
}

func renderTextInputS(r *RenderContext, f DataField, val any, config ElementOpts, style string, errorStyle string) {
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

	r.INPUT(name, inputType, val, style, config)
	if f.HasError() {
		r.PS(f.Errors(), errorStyle)
	}
}

func renderTextInput(r *RenderContext, f DataField, val any, config ElementOpts, prefix string, class string) {
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
	r.out.WriteString(fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s value=\"%v\"%s class=\"%s\"/>",
		name, inputType,
		strings.TrimSpace(inputConstraints), val, configToHtml(config), class))
	if f.HasError() {
		r.out.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", f.Errors()))
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
	LABEL func(l string) string,
) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, Choice{
					Group:    group(p),
					Label:    LABEL(p),
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
