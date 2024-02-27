package uni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClean(t *testing.T) {
	cleaned := RemoveSpacesInHtml(`
<html>   <body  class="body">
</body>
</html>
`)
	assert.Equal(t, "<html><body  class=\"body\"></body></html>", cleaned)
}

func TestNormalize(t *testing.T) {
	norm := RemoveSpacesInHtml(RemoveClassAndStyle(`
<div class="a">
<div class="b">
</div>
</div>
`))
	assert.Equal(t, "<div><div></div></div>", norm)
}
