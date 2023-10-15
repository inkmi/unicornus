package pkg

import (
	"fmt"
	"strings"
)

func (f *FormLayout) RenderForm(data any) string {
	var sb strings.Builder
	f.renderFormToBuilder(&sb, data, "")
	return sb.String()
}

func (f *FormLayout) renderFormToBuilder(sb *strings.Builder, data any, prefix string) {
	m := FieldsToMap(FieldGenerator(data))
	for _, e := range f.elements {
		switch e.Kind {
		case "header":
			f.Theme.themeRenderHeader(sb, e)
		case "group":
			newPrefix := e.Name
			if len(prefix) > 0 {
				newPrefix = prefix + "." + newPrefix
			}
			f.Theme.themeRenderGroup(sb, data, newPrefix, e)
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
					f.Theme.themeRenderMulti(sb, field, e, prefix)
				} else {
					if field.Kind == "bool" {
						f.Theme.themeRenderCheckbox(sb, e, field, prefix)
					} else if !field.Multi && len(field.Choices) > 0 {
						f.Theme.themeRenderSelect(sb, e, field, prefix)
					} else {
						f.Theme.themeRenderInput(sb, e, field, prefix)
					}
				}
			}
		}
	}
}

func renderCheckbox(sb *strings.Builder, f DataField, config ElementConfig, prefix string, class string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
		name := f.Name
		sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"%s\" %s%s/>", name, class, checked, configToHtml(config)))
	}
}

func renderSelect(sb *strings.Builder, f DataField, config ElementConfig, prefix string, class string) {
	sb.WriteString(fmt.Sprintf("<select name=\"%s\" class=\"%s\"><option value=\"0\">-</option>", f.Name, class))

	for _, c := range f.Choices {
		if c.IsSelected(f.Value) {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", c.Val(), c.L()))
		} else {
			sb.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", c.Val(), c.L()))
		}
	}
	sb.WriteString("</select>")
}

func renderTextInput(sb *strings.Builder, f DataField, val any, config ElementConfig, prefix string, class string) {
	sb.WriteString(fmt.Sprintf("<input name=\"%s\" value=\"%v\"%s/ class=\"%s\">", f.Name, val, configToHtml(config), class))
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
