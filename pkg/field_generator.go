package pkg

import (
	"reflect"
)
import . "github.com/moznion/go-optional"

func FieldGenerator(obj interface{}) []DataField {
	vals := make([]DataField, 0, 20)
	original := reflect.ValueOf(obj)
	return translateRecursive(vals, "", original)
	// return translateRecursive(vals, original.Type().Name(), original)
}

func translateRecursive(vals []DataField, prefix string, original reflect.Value) []DataField {
	switch original.Kind() {
	case reflect.Struct:
		vals = translateStruct(prefix, vals, original)
		return vals
	case reflect.Ptr:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return vals
		}
		return translateRecursive(vals, prefix, originalValue)
	case reflect.Interface:
		originalValue := original.Elem()
		translateRecursive(vals, prefix, originalValue)
		return vals
	default:
		return vals
	}
}

func translateStruct(prefix string, vals []DataField, original reflect.Value) []DataField {
	for i := 0; i < original.NumField(); i += 1 {
		df := DataField{}
		if len(prefix) == 0 {
			df.Name = original.Type().Field(i).Name
		} else {
			df.Name = prefix + "." + original.Type().Field(i).Name
		}
		// parse Tag for validation, choices and error messages
		tagS := string(original.Type().Field(i).Tag)
		if tagS != "" {
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
			if t.Choices != nil && len(t.Choices) > 0 {
				df.Choices = t.Choices
			}
		}

		// Optional:
		// 1: With Type e.g. Option[int]
		// 2: By pointer, e.g. *bool
		// 3: By validation, e.g. "validate:"optional"

		optionalValue := df.Optional || hasOptional(original.Type().Field(i).Type.Name())
		// Handle setFromStructPtr of type
		// those are handled as optional
		// *bool
		// *int
		if original.Type().Field(i).Type.Kind() == reflect.Ptr {
			df.Kind = original.Type().Field(i).Type.Elem().Name()
			df.Optional = true
			if original.Type().Field(i).Type.Elem().ConvertibleTo(reflect.TypeOf(0)) {
				df.Kind = "int"
				setValue(&df, original, i)
			} else if original.Type().Field(i).Type.Elem().ConvertibleTo(reflect.TypeOf(true)) {
				df.Kind = "bool"
				setValue(&df, original, i)
			}
		} else {
			typ := ""
			// it seems that the Kind of the Option is Slice
			// so check with the type we got with Name
			if !optionalValue && original.Type().Field(i).Type.Kind() == reflect.Slice {
				df.Multi = true
				typ = original.Type().Field(i).Type.Elem().Name()
			} else {
				typ = original.Type().Field(i).Type.Name()
			}
			if df.Optional == false {
				df.Optional = hasOptional(typ)
			}
			df.Kind = removeOptional(typ)
			setValue(&df, original, i)
		}

		vals = append(vals, df)
		newPrefix := prefix + " ." + original.Type().Field(i).Name
		if len(prefix) == 0 {
			newPrefix = original.Type().Field(i).Name
		}
		vals = translateRecursive(vals, newPrefix, original.Field(i))
	}
	return vals
}

func setValue(df *DataField, original reflect.Value, i int) {
	if df.Kind == "string" {
		setString(original, i, df)
	} else if df.Kind == "bool" {
		setBool(original, i, df)
	} else if df.Kind == "int" {
		setInt(original, i, df)
	}
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
