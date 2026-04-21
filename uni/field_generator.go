package uni

import (
	"reflect"
	"time"

	. "github.com/moznion/go-optional"
)

// Converts a struct and errors into a slice of DataFields

func FieldGenerator(obj interface{}, errors map[string]string) []DataField {
	vals := make([]DataField, 0, 20)
	original := reflect.ValueOf(obj)
	return translateRecursive(vals, "", original, errors)
}

func translateRecursive(vals []DataField, prefix string, original reflect.Value, errors map[string]string) []DataField {
	switch original.Kind() {
	case reflect.Struct:
		vals = translateStruct(prefix, vals, original, errors)
		return vals
	case reflect.Ptr:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return vals
		}
		return translateRecursive(vals, prefix, originalValue, errors)
	case reflect.Interface:
		originalValue := original.Elem()
		translateRecursive(vals, prefix, originalValue, errors)
		return vals
	default:
		return vals
	}
}

func translateStruct(prefix string, vals []DataField, original reflect.Value, errors map[string]string) []DataField {
	for i := 0; i < original.NumField(); i += 1 {
		df, fieldType := buildDataField(prefix, original, i, errors)
		vals = append(vals, df)

		if df.Kind == "Time" || df.Kind == "Date" {
			continue
		}
		newPrefix := df.Name
		_ = fieldType
		vals = translateRecursive(vals, newPrefix, original.Field(i), errors)
	}
	return vals
}

func buildDataField(prefix string, original reflect.Value, i int, errors map[string]string) (DataField, reflect.Type) {
	df := DataField{}
	fieldType := original.Type().Field(i)
	name := fieldType.Name
	if len(prefix) == 0 {
		df.Name = name
	} else {
		df.Name = prefix + "." + name
	}
	if errorMsg, hasError := errors[name]; hasError {
		df.ErrorMessages = []string{errorMsg}
	}

	parsedTag := applyTag(&df, string(fieldType.Tag))

	if fieldType.Type.Kind() == reflect.Ptr {
		applyPtrKind(&df, original, i)
	} else {
		applyValueKind(&df, original, i, parsedTag)
	}
	return df, fieldType.Type
}

func applyTag(df *DataField, tagS string) *Tag {
	if tagS == "" {
		return nil
	}
	t := ParseTag(tagS)
	if t.Optional {
		df.Optional = t.Optional
	}
	if t.Validation != nil {
		df.Validation = *t.Validation
	}
	if t.ErrorMessage != nil {
		df.ErrorMessage = *t.ErrorMessage
	}
	if len(t.Choices) > 0 {
		df.Choices = t.Choices
	}
	return &t
}

func applyPtrKind(df *DataField, original reflect.Value, i int) {
	elem := original.Type().Field(i).Type.Elem()
	df.Kind = elem.Name()
	df.Optional = true
	if elem.ConvertibleTo(reflect.TypeOf(0)) {
		df.Kind = "int"
		setValue(df, original, i)
	} else if elem.ConvertibleTo(reflect.TypeOf(true)) {
		df.Kind = "bool"
		setValue(df, original, i)
	}
}

func applyValueKind(df *DataField, original reflect.Value, i int, parsedTag *Tag) {
	fieldType := original.Type().Field(i).Type
	optionalValue := df.Optional || hasOptional(fieldType.Name())
	typ := ""
	if !optionalValue && fieldType.Kind() == reflect.Slice {
		df.Multi = true
		typ = fieldType.Elem().Name()
	} else {
		typ = fieldType.Name()
	}
	if !df.Optional {
		df.Optional = hasOptional(typ)
	}
	df.Kind = removeOptional(typ)

	if df.Kind == "Time" && parsedTag != nil && parsedTag.InputType != nil && *parsedTag.InputType == "date" {
		df.Kind = "Date"
	}

	setValue(df, original, i)
}

func setValue(df *DataField, original reflect.Value, i int) {
	if df.Kind == "string" {
		setString(original, i, df)
	} else if df.Kind == "bool" {
		setBool(original, i, df)
	} else if df.Kind == "int" {
		setInt(original, i, df)
	} else if df.Kind == "Time" {
		setTime(original, i, df)
	} else if df.Kind == "Date" {
		setDate(original, i, df)
	}
}

func setTime(original reflect.Value, i int, f *DataField) {
	f.Value = original.Field(i).Interface().(time.Time).Format("2006-01-02T15:04")
}

func setDate(original reflect.Value, i int, f *DataField) {
	f.Value = original.Field(i).Interface().(time.Time).Format("2006-01-02")
}

func setString(original reflect.Value, i int, f *DataField) {
	if f.Multi {
		// https://stackoverflow.com/questions/32890137/how-to-get-slice-underlying-value-via-reflect-value
		f.Value = original.Field(i).Interface().([]string)
	} else {
		f.Value = original.Field(i).String()
	}
}

func setBool(original reflect.Value, i int, f *DataField) {
	if !f.Optional {
		f.Value = original.Field(i).Bool()
	} else {
		if hasOptional(original.Type().Field(i).Type.Name()) {
			o := original.Field(i).Interface().(Option[bool])
			if o.IsSome() {
				val, _ := o.Take()
				f.Value = val
			} else {
				f.Value = false
			}
		} else {
			// this is *bool
			valid := original.Field(i).Elem().IsValid()
			if valid {
				f.Value = original.Field(i).Elem().Bool()
			}
		}
	}
}

func setInt(original reflect.Value, i int, f *DataField) {
	if !f.Optional {
		f.Value = original.Field(i).Int()
	} else {
		if hasOptional(original.Type().Field(i).Type.Name()) {
			o := original.Field(i).Interface().(Option[int])
			if o.IsSome() {
				val, _ := o.Take()
				f.Value = int64(val)
			} else {
				f.Value = nil
			}
		} else if original.Field(i).Kind() == reflect.Ptr {
			// this is *int
			valid := original.Field(i).Elem().IsValid()
			if valid {
				f.Value = original.Field(i).Elem().Int()
			}
		} else {
			// this is int `validate:"optional"`
			valid := original.Field(i).IsValid()
			if valid {
				f.Value = original.Field(i).Int()
			}
		}
	}
}
