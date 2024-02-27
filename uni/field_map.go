package uni

func FieldsToMap(fields []DataField) map[string]DataField {
	m := make(map[string]DataField)
	for _, f := range fields {
		m[f.Name] = f
	}
	return m
}
