package ui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveOptional(t *testing.T) {
	typ := "int"
	assert.Equal(t, "int", removeOptional(typ))
	typ = ""
	assert.Equal(t, "", removeOptional(typ))
	typ = "Option[int]"
	assert.Equal(t, "int", removeOptional(typ))
}
