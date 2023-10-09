package builder

import (
	"fmt"
	"strings"
	"unicornus/internal/ui"
)

func (f *FormLayout) RenderForm(data any) string {
	m := ui.FieldsToMap(ui.FieldGenerator(data))
	var sb strings.Builder
	for _, e := range f.elements {
		switch e.Kind {
		case "header":
			sb.WriteString(fmt.Sprintf("<h2>%s</h2>", e.Name))
		case "input":
			// take value string from MAP of name -> DataField
			// take type if no type is given from DataField
			field, ok := m[e.Name]
			if ok {
				if len(e.Config.Choices) > 0 {
					field.Choices = e.Config.Choices
				}
				//sb.WriteString("<p>")
				if field.Multi {
					renderMulti(&sb, field, e.Config)
				} else {
					if len(e.Config.Label) > 0 {
						sb.WriteString(fmt.Sprintf("<label>%s</label>", e.Config.Label))
					}
					if field.Kind == "bool" {
						renderCheckbox(&sb, field, e.Config)
					} else if field.Multi == false && len(field.Choices) > 0 {
						renderSelect(&sb, field, e.Config)
					} else {
						renderTextInput(&sb, field, field.Val(), e.Config)
					}
				}
				//sb.WriteString("</p>")
			}
		}
	}
	return sb.String()
}

func renderCheckbox(sb *strings.Builder, f ui.DataField, config ElementConfig) {
	checked := ""
	sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" %s%s/>", f.Name, checked, configToHtml(config)))
	return
}

func renderMulti(sb *strings.Builder, f ui.DataField, config ElementConfig) {
	if len(config.Groups) > 0 {
		for _, group := range config.Groups {
			sb.WriteString(fmt.Sprintf("<fieldset>"))
			for _, c := range f.Choices {
				if c.Group == group {
					if c.IsSelected(f.Value) {
						sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked>", c.Label+"#"+c.Val()))
					} else {
						sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\">", f.Name+"#"+c.Val()))
					}
					sb.WriteString(fmt.Sprintf(`<label>%s</label>`, c.L()))
				}
			}
		}
		sb.WriteString("</fieldset>")
	} else {
		sb.WriteString(fmt.Sprintf("<fieldset>"))
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
	return
}

func renderSelect(sb *strings.Builder, f ui.DataField, config ElementConfig) {
	sb.WriteString(fmt.Sprintf("<select name=\"%s\"><option value=\"0\">-</option>", f.Name))
	for _, c := range f.Choices {
		if c.IsSelected(f.Value) {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))

		} else {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
		}
	}
	sb.WriteString("</select>")
	return
}

func renderTextInput(sb *strings.Builder, f ui.DataField, val any, config ElementConfig) {
	sb.WriteString(fmt.Sprintf("<input name=\"%s\" value=\"%s\"%s/>", f.Name, val, configToHtml(config)))
	return
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
