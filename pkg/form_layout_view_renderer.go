package pkg

import (
	"fmt"
	"strings"
)

func (f *FormLayout) RenderView(data any) string {
	fields := FieldGenerator(data)
	m := FieldsToMap(fields)
	var sb strings.Builder
	for _, e := range f.elements {
		field, ok := m[e.Name]
		if ok {
			sb.WriteString(fmt.Sprintf("%s", field.Val()))
		}
	}
	return sb.String()
}
