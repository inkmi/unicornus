package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldMap(t *testing.T) {
	fields := []DataField{{
		Name: "a.b.c",
	},
	}
	m := FieldsToMap(fields)
	_, ok := m["a.b.c"]
	assert.True(t, ok)
}
