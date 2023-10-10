package builder

import (
	"fmt"
	"github.com/Inkmi/unicornus/internal/ui"
	"strings"
)

func (f *FormLayout) RenderForm(data any) string {
	var sb strings.Builder
	f.renderFormToBuilder(&sb, data, "")
	return sb.String()
}

func (f *FormLayout) renderFormToBuilder(sb *strings.Builder, data any, prefix string) {
	m := ui.FieldsToMap(ui.FieldGenerator(data))
	for _, e := range f.elements {
		switch e.Kind {
		case "header":
			sb.WriteString(fmt.Sprintf("<h2>%s</h2>", e.Name))
		case "group":
			sb.WriteString("<div>")
			sb.WriteString(e.Name)
			newPrefix := e.Name
			if len(prefix) > 0 {
				newPrefix = prefix + "." + newPrefix
			}
			e.Config.SubLayout.renderFormToBuilder(sb, data, newPrefix)
			sb.WriteString("</div>")
		case "input":
			// take value string from MAP of name -> DataField
			// take type if no type is given from DataField
			field, ok := m[e.Name]
			if ok {
				if len(e.Config.Choices) > 0 {
					field.Choices = e.Config.Choices
				}
				if field.Multi {
					renderMulti(sb, field, e.Config, prefix)
				} else {
					if len(e.Config.Label) > 0 {
						sb.WriteString(fmt.Sprintf("<label>%s</label>", e.Config.Label))
					}
					if field.Kind == "bool" {
						renderCheckbox(sb, field, e.Config, prefix)
					} else if !field.Multi && len(field.Choices) > 0 {
						renderSelect(sb, field, e.Config, prefix)
					} else {
						renderTextInput(sb, field, field.Val(), e.Config, prefix)
					}
				}
			}
		}
	}
}

func renderCheckbox(sb *strings.Builder, f ui.DataField, config ElementConfig, prefix string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v == true {
			checked = "checked"
		}
		name := f.Name
		if len(prefix) > 0 {
			name = prefix + "." + name
		}
		sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" %s%s/>", name, checked, configToHtml(config)))
	}
}

func renderMulti(sb *strings.Builder, f ui.DataField, config ElementConfig, prefix string) {
	if len(config.Groups) > 0 {
		for _, group := range config.Groups {
			sb.WriteString("<fieldset>")
			values := f.Value.([]string)
			// range copies slice
			for i := 0; i < len(f.Choices); i++ {
				choice := &f.Choices[i]
				if containsString(values, choice.Value) {
					choice.Selected = true
				}
			}
			for _, c := range f.Choices {
				if c.Group == group {
					name := f.Name + "#" + c.Val()
					if len(prefix) > 0 {
						name = prefix + "." + name
					}
					if c.Selected {
						sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked>", name))
					} else {
						sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\">", name))
					}
					sb.WriteString(fmt.Sprintf(`<label>%s</label>`, c.L()))
				}
			}
			sb.WriteString("</fieldset>")
		}
	} else {
		sb.WriteString("<fieldset>")
		for _, c := range f.Choices {
			if c.IsSelected(f.Value) {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked>", c.Label+"#"+c.Val()))
			} else {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\">", f.Name+"#"+c.Val()))
			}
			sb.WriteString(fmt.Sprintf(`<label>%s</label>`, c.L()))

		}
		sb.WriteString("</fieldset>")
	}
}

func renderSelect(sb *strings.Builder, f ui.DataField, config ElementConfig, prefix string) {
	sb.WriteString(fmt.Sprintf("<select name=\"%s\"><option value=\"0\">-</option>", f.Name))
	for _, c := range f.Choices {
		if c.IsSelected(f.Value) {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))

		} else {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
		}
	}
	sb.WriteString("</select>")
}

func renderTextInput(sb *strings.Builder, f ui.DataField, val any, config ElementConfig, prefix string) {
	sb.WriteString(fmt.Sprintf("<input name=\"%s\" value=\"%s\"%s/>", f.Name, val, configToHtml(config)))
}

func configToHtml(config ElementConfig) string {
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

func SetChoices(setKey string, fields []ui.FieldV, allValues []string) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []ui.Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, ui.Choice{
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
	fields []ui.FieldV,
	allValues []string,
	group func(k string) string,
	label func(l string) string,
) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []ui.Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, ui.Choice{
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
