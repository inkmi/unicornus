package uni

import (
	"strconv"
	"strings"

	. "github.com/moznion/go-optional"
)

type ValidationData struct {
	Optional bool
	Min      Option[int]
	Max      Option[int]
	Subtype  Option[string]
}

func GetValidations(v []DataField) []ValidationData {
	vd := make([]ValidationData, 0)
	for _, d := range v {
		vd = append(vd, GetValidation(d))
	}
	return vd
}

func GetValidation(v DataField) ValidationData {
	vd := ValidationData{}
	vd.Optional = v.Optional

	validations := strings.Split(v.Validation, "|")

	if strings.HasPrefix(strings.ToLower(v.Validation), "email") {
		vd.Subtype = Some("email")
	} else if strings.HasPrefix(strings.ToLower(v.Validation), "url") {
		vd.Subtype = Some("url")
	}

	// Check for int and minV/max
	minV := None[int]()
	maxV := None[int]()
	for _, vint := range validations {
		if strings.HasPrefix(vint, "min:") {
			minInt, _ := strconv.Atoi(vint[len("min:"):])
			minV = Some(minInt)
		}
		if strings.HasPrefix(vint, "max:") {
			maxInt, _ := strconv.Atoi(vint[len("max:"):])
			maxV = Some(maxInt)
		}
	}
	vd.Min = minV
	vd.Max = maxV
	return vd
}
