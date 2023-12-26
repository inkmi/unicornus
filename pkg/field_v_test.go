package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVal(t *testing.T) {
	v := 10
	d := DataField{
		Optional: true,
		Value:    v,
	}
	assert.Equal(t, 10, d.Val())
}

func TestChoice(t *testing.T) {
	v := int64(10)
	choices := []Choice{
		{
			Label: "Zehn",
			Value: "10",
		},
		{
			Label: "Elf",
			Value: "11",
		},
	}
	d := DataField{
		Kind:    "int",
		Choices: choices,
		Value:   v,
	}
	assert.Equal(t, "Zehn", d.ViewVal())
}

func TestSanitizeVal(t *testing.T) {
	v := "<b>x</b><script>alert(\"hello\");</script>"
	d := DataField{
		Optional: true,
		Value:    v,
	}
	assert.Equal(t, "x", d.Val())
}
