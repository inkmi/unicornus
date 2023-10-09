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
