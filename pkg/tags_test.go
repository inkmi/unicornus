package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTagValidation(t *testing.T) {
	tagString := `validate:"int|min:1000"`
	tag := ParseTag(tagString)
	expected := "int|min:1000"
	assert.Equal(t, &expected, tag.Validation)
}

func TestParseTagErrorMessage(t *testing.T) {
	tagString := `message:"error"`
	tag := ParseTag(tagString)
	expected := "error"
	assert.Equal(t, &expected, tag.ErrorMessage)
}

func TestValidateChoices(t *testing.T) {
	tagString := `validate:"in:1,2,3" choices:"a|b|c"`
	tag := ParseTag(tagString)
	expected := []Choice{
		{Label: "a", Value: "1"},
		{Label: "b", Value: "2"},
		{Label: "c", Value: "3"},
	}
	validation := "in:1,2,3"
	assert.Equal(t, &validation, tag.Validation)
	assert.Equal(t, expected, tag.Choices)
}

func TestChoices(t *testing.T) {
	tagString := `choices:"a|b|c"`
	tag := ParseTag(tagString)
	expected := []Choice{
		{Label: "a", Value: "a"},
		{Label: "b", Value: "b"},
		{Label: "c", Value: "c"},
	}
	assert.Equal(t, expected, tag.Choices)
}

func TestOptional(t *testing.T) {
	tagString := `validate:"optional"`
	tag := ParseTag(tagString)
	assert.True(t, tag.Optional)
}

func TestGetInValidation(t *testing.T) {
	val1 := `in:0,2,3|int`
	assert.Equal(t, "0,2,3", GetInValidation(val1))
	val2 := "in:0,2,4"
	assert.Equal(t, "0,2,4", GetInValidation(val2))
}
