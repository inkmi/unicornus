package pkg

import (
	"fmt"
	"strings"
)

type Theme interface {
	themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string)
	themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, prefix string)
	themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, prefix string)
	themeRenderMulti(sb *strings.Builder, field DataField, e FormElement, prefix string)
	themeRenderHeader(sb *strings.Builder, e FormElement)
	themeRenderGroup(sb *strings.Builder, data any, prefix string, e FormElement)
}

type TailwindTheme struct {
}

func (t TailwindTheme) themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label>%s</label>", e.Config.Label))
	}
	renderTextInput(sb, field, field.Val(), e.Config, prefix)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label>%s</label>", e.Config.Label))
	}
	renderSelect(sb, field, e.Config, prefix)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div>")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label>%s</label>", e.Config.Label))
	}
	renderCheckbox(sb, field, e.Config, prefix)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderMulti(sb *strings.Builder, field DataField, e FormElement, prefix string) {
	sb.WriteString("<div>")
	renderMulti(sb, field, e.Config, prefix)
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
