package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestNestG struct {
	A    int
	Nest TestG
}

type TestOptionalCheckbox struct {
	A *bool
}

type TestG struct {
	B int  `validate:"int|min:10" message:"The age is 10" LABEL:"Minimum Age"`
	C *int `validate:"int|min:20" message:"The age is 20" LABEL:"Minimum Age"`
	D *int `validate:"int|min:30" message:"The age is 30" LABEL:"Minimum Age"`
	E []string
}

type TestIntOptional struct {
	A int `validate:"optional|int|min:30" message:"The age is 30" LABEL:"Minimum Age"`
}

type TestG2 struct {
	A []string `choices:"A1|A2|A3"`
}

func TestFieldOptionalInt(t *testing.T) {
	d := TestIntOptional{
		A: 3,
	}
	fields := FieldGenerator(d, nil)
	assert.Equal(t, 1, len(fields))
	assert.Equal(t, "A", fields[0].Name)
	assert.Equal(t, "int", fields[0].Kind)
	assert.False(t, fields[0].Multi)
	assert.True(t, fields[0].Optional)
}

func TestFieldOptionalCheckbox(t *testing.T) {
	d := TestOptionalCheckbox{
		A: nil,
	}
	fields := FieldGenerator(d, nil)
	assert.Equal(t, 1, len(fields))
	assert.Equal(t, "A", fields[0].Name)
	assert.Equal(t, "bool", fields[0].Kind)
	assert.False(t, fields[0].Multi)
}

func TestFieldsMulti(t *testing.T) {
	d := TestG2{
		A: []string{"A1", "A2"},
	}
	fields := FieldGenerator(d, nil)
	assert.Equal(t, 1, len(fields))

	assert.Equal(t, "A", fields[0].Name)
	assert.Equal(t, "string", fields[0].Kind)
	assert.True(t, fields[0].Multi)
	assert.Equal(t, 3, len(fields[0].Choices))
}

func TestFields(t *testing.T) {
	x := 3
	d := TestNestG{
		A: 8,
		Nest: TestG{
			B: 10,
			C: &x,
			D: nil,
			E: []string{"a"},
		},
	}
	fields := FieldGenerator(d, nil)
	assert.Equal(t, 6, len(fields))

	assert.Equal(t, "A", fields[0].Name)
	assert.Equal(t, "int", fields[0].Kind)
	assert.Equal(t, int64(8), fields[0].Value)

	assert.Equal(t, "Nest", fields[1].Name)
	assert.Equal(t, "TestG", fields[1].Kind)

	assert.Equal(t, "Nest.B", fields[2].Name)
	assert.Equal(t, "int", fields[2].Kind)
	assert.Equal(t, int64(10), fields[2].Value)

	assert.Equal(t, "Nest.C", fields[3].Name)
	assert.Equal(t, "int", fields[3].Kind)
	assert.Equal(t, int64(3), fields[3].Value)

	assert.Equal(t, "Nest.E", fields[5].Name)
	assert.Equal(t, "string", fields[5].Kind)
	assert.Equal(t, []string{"a"}, fields[5].Value)
}
