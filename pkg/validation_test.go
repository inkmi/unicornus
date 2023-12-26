package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test2 struct {
	MaxAge int    `validate:"int|max:20|min:10" message:"The age is 10" label:"Maximum and Minimum Age"`
	Url    string `validate:"url"`
	Email  string `validate:"email"`
}

func TestValidation(t *testing.T) {
	d := Test2{
		MaxAge: 3,
		Url:    "https://www.inkmi.com",
		Email:  "test@example.com",
	}
	fields := FieldGenerator(d, nil)
	assert.Equal(t, 3, len(fields))

	vds := GetValidations(fields)
	assert.Equal(t, 3, len(vds))

	assert.Equal(t, 20, vds[0].Max.Unwrap())
	assert.Equal(t, 10, vds[0].Min.Unwrap())

	assert.Equal(t, "url", vds[1].Subtype.Unwrap())
	assert.Equal(t, "email", vds[2].Subtype.Unwrap())

}
