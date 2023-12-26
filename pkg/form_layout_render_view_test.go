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
	assert.Equal(t, "<div class=\"mt-6\"><div class=\"block text-sm font-medium text-gray-700\">A</div><div>b</div></div>", html)
}
