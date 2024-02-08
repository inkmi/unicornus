package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderView(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")

	data := TestA{
		A: "b",
	}
	html := f.RenderView(data)
	assert.Equal(t, "<div style=\"margin-top: 1.5rem;\"><div class=\"block text-sm font-medium text-gray-500\">A</div><div class=\"text-sm font-medium text-gray-900\">b</div></div>", html)
}
