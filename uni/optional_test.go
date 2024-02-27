package uni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOptional(t *testing.T) {
	typ := "int"
	assert.Equal(t, "int", removeOptional(typ))
	typ = ""
	assert.Equal(t, "", removeOptional(typ))
	typ = "Option[int]"
	assert.Equal(t, "int", removeOptional(typ))
}
