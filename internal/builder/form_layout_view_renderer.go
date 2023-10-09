package builder

import (
	"fmt"
	"strings"
	"unicornus/internal/ui"
)

func (f *FormLayout) RenderView(data any) string {
	fields := ui.FieldGenerator(data)
	m := ui.FieldsToMap(fields)
	var sb strings.Builder
	for _, e := range f.elements {
		field, ok := m[e.Name]
		if ok {
			sb.WriteString(fmt.Sprintf("%s", field.Val()))
		}
	}
	return sb.String()
}
