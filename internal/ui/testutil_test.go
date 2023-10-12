package ui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClean(t *testing.T) {
	cleaned := Clean(`
<html>   <body  class="body">
</body>
</html>
`)
	assert.Equal(t, "<html><body  class=\"body\"></body></html>", cleaned)
}

func TestNormalize(t *testing.T) {
	norm := Clean(Normalize(`
<div class="a">
<div class="b">
</div>
</div>
`))
	assert.Equal(t, "<div><div></div></div>", norm)
}
