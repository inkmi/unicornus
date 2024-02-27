package uni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderView(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")

	data := TestA{
		A: "b",
	}
	html := f.RenderView(data)
	assert.Equal(t, "<div style=\"margin-top: 1.5rem;\"><div style=\"display: block; font-size: 0.875rem; font-weight: 500; color: #6B7280\">A</div><div class=\"font-size: 0.875rem; font-weight: 500; color: #1F2937;\">b</div></div>", html)
}
