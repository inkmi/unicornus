package pkg

import "C"
import (
	"github.com/fatih/structtag"
	"strings"
)

func GetInValidation(v string) string {
	start := strings.Index(v, "in:")
	end := strings.Index(v[start:], "|")
	if end == -1 {
		end = len(v)
	}
	return v[start+3 : end]
}

func ParseTag(tag string) Tag {
	tg := Tag{}
	tags, err := structtag.Parse(tag)
	if err == nil {
		var choices []string
		for _, t := range tags.Tags() {
			if t.Key == "validate" {
				v := t.Value()
				tg.Validation = &v
			}
			if t.Key == "choices" {
				choices = strings.Split(t.Value(), "|")
			}
			if t.Key == "message" {
				v := t.Value()
				tg.ErrorMessage = &v
			}
		}
		if choices != nil && tg.Validation != nil {
			in := GetInValidation(*tg.Validation)
			inValues := strings.Split(in, ",")
			if len(inValues) == len(choices) {
				ch := make([]Choice, 0)
				for i, v := range inValues {
					value := v
					c := Choice{
						Label: choices[i],
						Value: value,
					}
					ch = append(ch, c)
				}
				tg.Choices = ch
			}
		} else if choices != nil {
			ch := make([]Choice, 0)
			for i, v := range choices {
				value := v
				c := Choice{
					Label: choices[i],
					Value: value,
				}
				ch = append(ch, c)
			}
			tg.Choices = ch
		}
	}
	return tg
}
