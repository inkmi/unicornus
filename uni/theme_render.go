package uni

import (
	"strings"
)

type Theme interface {
	themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string)
	themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, prefix string)
	themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, description string, prefix string)
	themeRenderMulti(sb *strings.Builder, field DataField, e FormElement, prefix string)
	themeRenderHeader(sb *strings.Builder, e FormElement)
	themeRenderGroup(sb *strings.Builder, data any, prefix string, e FormElement)
}
