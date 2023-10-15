package pkg

import (
	"fmt"
	"strings"
)

type TailwindTheme struct {
}

func (t TailwindTheme) themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
	renderTextInput(sb, field, field.Val(), e.Config, prefix, class)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
	renderSelect(sb, field, e.Config, prefix, class)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
	renderCheckbox(sb, field, e.Config, prefix, class)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderMulti(sb *strings.Builder, f DataField, e FormElement, prefix string) {
	sb.WriteString("<div>")
	// Should this move to Field generation?
	values := f.Value.([]string)
	for i := 0; i < len(f.Choices); i++ {
		choice := &f.Choices[i]
		if containsString(values, choice.Value) {
			choice.Checked = true
		}
	}
	if len(e.Config.Groups) > 0 {
		for _, group := range e.Config.Groups {
			t.renderGroup(sb, f, group, "", "")
		}
	} else {
		t.renderGroup(sb, f, "", "", "")
	}
	sb.WriteString("</div>")
}

func (t TailwindTheme) renderGroup(sb *strings.Builder, f DataField, group string, class1 string, class2 string) {
	sb.WriteString(fmt.Sprintf("<div class=\"%s\">", class1))
	sb.WriteString("<fieldset>")
	// range copies slice
	for _, c := range f.Choices {
		if len(group) == 0 || c.Group == group {
			name := f.Name + "#" + c.Val()
			sb.WriteString(fmt.Sprintf("<div class=\"%s\">", class2))
			if c.Checked {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked>", name))
			} else {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\">", name))
			}
			sb.WriteString(fmt.Sprintf(`<label>%s</label>`, c.L()))
			sb.WriteString("</div>")
		}
	}
	sb.WriteString("</fieldset>")
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderHeader(sb *strings.Builder, e FormElement) {
	sb.WriteString(fmt.Sprintf("<h2>%s</h2>", e.Name))
}

func (t TailwindTheme) themeRenderGroup(sb *strings.Builder, data any, prefix string, e FormElement) {
	sb.WriteString("<div>")
	sb.WriteString(e.Name)
	e.Config.SubLayout.renderFormToBuilder(sb, data, prefix)
	sb.WriteString("</div>")
}
