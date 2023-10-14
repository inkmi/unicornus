package js

import (
	"github.com/inkmi/unicornus/pkg/ui"
	. "github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test2 struct {
	MaxAge int `validate:"int|max:20|min:10" message:"The age is 10" label:"Maximum and Minimum Age"`
}

type TestG struct {
	B int  `validate:"int|min:10" message:"The age is 10" label:"Minimum Age"`
	C *int `validate:"int|min:20" message:"The age is 20" label:"Minimum Age"`
	D *int `validate:"int|min:30" message:"The age is 30" label:"Minimum Age"`
	E []string
}

type TestOG struct {
	C Option[int]
	D int
	E Option[int]
}

func TestFillDataFromFields(t *testing.T) {
	d := TestOG{
		C: Some(4),
		D: 3,
		E: None[int](),
	}

	data := map[string]any{}
	fields := ui.FieldGenerator(d)
	FillDataFromFields(fields, data)

	value, isPresent := data["D"]
	assert.True(t, isPresent)
	assert.Equal(t, int64(3), value.(ui.DataField).Value)

	value, isPresent = data["C"]
	assert.True(t, isPresent)
	assert.Equal(t, int64(4), value.(ui.DataField).Value)

	value, isPresent = data["E"]
	assert.True(t, isPresent)
	assert.Equal(t, nil, value.(ui.DataField).Value)

}

func TestValidation(t *testing.T) {
	d := Test2{
		MaxAge: 3,
	}
	fields := ui.FieldGenerator(d)
	validation := Validation(fields[0])
	expected := "\"MaxAge\" : { \"typ\":\"int\", \"validate\" : function(v) { return validator.isInt(v, { min: 10, max: 20 }); },\"error\" : \"The age is 10\", }"
	assert.Equal(t, expected, validation)
}

func TestGenerate(t *testing.T) {
	expected := "{\"B\" : { \"typ\":\"int\", \"validate\" : function(v) { return validator.isInt(v, { min: 10 }); },\"error\" : \"The age is 10\", },\"C\" : { \"typ\":\"int\", \"validate\" : function(v) { return v.trim().length === 0 || validator.isInt(v, { min: 20 }); },\"error\" : \"The age is 20\", },\"D\" : { \"typ\":\"int\", \"validate\" : function(v) { return v.trim().length === 0 || validator.isInt(v, { min: 30 }); },\"error\" : \"The age is 30\", },\"E\" : { \"typ\":\"string\", \"validate\" : function(v) { return true; },\"error\" : \"\", }}"

	x := 3
	d := TestG{
		B: 8,
		C: &x,
	}
	fields := ui.FieldGenerator(d)
	g := Generate(fields)
	assert.Equal(t, expected, g)
}
