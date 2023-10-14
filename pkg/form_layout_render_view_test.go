package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderView(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	tdata := TestA{
		A: "b",
	}
	html := f.RenderView(tdata)
	assert.Equal(t, "b", html)
}
