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
