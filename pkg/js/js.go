package js

import (
	"github.com/inkmi/unicornus/pkg/ui"
	. "github.com/moznion/go-optional"
	"strconv"
	"strings"
)

// Initially this was iterating over all values not types
// might be useful
// from: https://gist.github.com/tamalsaha/5badc2d686995d787715278c968c0bf0

func FillDataFromFields(fields []ui.DataField, data map[string]any) {
	for _, v := range fields {
		key := v.Name
		data[key] = v
	}
}

func Validation(v ui.DataField) string {
	validation := ""
	if v.Optional && len(v.Validation) > 0 {
		validation = "v.trim().length === 0 || "
	}
	validations := strings.Split(v.Validation, "|")

	// Check for int and min/max
	min := None[int]()
	max := None[int]()
	for _, vint := range validations {
		if strings.HasPrefix(vint, "min:") {
			minInt, _ := strconv.Atoi(vint[len("min:"):])
			min = Some(minInt)
		}
		if strings.HasPrefix(vint, "max:") {
			maxInt, _ := strconv.Atoi(vint[len("max:"):])
			max = Some(maxInt)
		}
	}
	if min.IsSome() && max.IsSome() {
		minInt := strconv.Itoa(min.TakeOr(0))
		maxInt := strconv.Itoa(max.TakeOr(0))
		validation = validation + "validator.isInt(v, { min: " + minInt + ", max: " + maxInt + " })"
	} else if min.IsSome() {
		minInt := strconv.Itoa(min.TakeOr(0))
		validation = validation + "validator.isInt(v, { min: " + minInt + " })"
	} else if max.IsSome() {
		maxInt := strconv.Itoa(max.TakeOr(0))
		validation = validation + "validator.isInt(v, { max: " + maxInt + " })"
	}

	// Check for in:
	for _, va := range validations {
		if strings.HasPrefix(va, "in:") {
			validation = validation + "true"
		}
	}
	if validation == "" {
		validation = "true"
	}
	js := "\"" +
		v.Name +
		"\" : { \"typ\":\"" + v.Kind + "\"" +
		", \"validate\" : function(v) {" +
		" return " + validation + "; " +
		"}," +
		"\"error\" : \"" + v.ErrorMessage + "\", " +
		"}"

	return js
}

func Generate(vals []ui.DataField) string {
	valJson := "{"
	for i, v := range vals {
		valJson = valJson + Validation(v)
		if i < len(vals)-1 {
			valJson = valJson + ","
		}
	}
	valJson = valJson + "}"
	return valJson
}
