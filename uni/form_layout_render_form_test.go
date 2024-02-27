package uni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestA struct {
	A string
}

type TestOptionCheckbox struct {
	A *bool
}

type TestMulti struct {
	A []string `choices:"A1|A2|A3"`
}

type TestBool struct {
	A bool
}

type TestGroup struct {
	B bool
	C *int
	D *int
}

type TestSubGroup struct {
	A TestGroup
}

type TestAB struct {
	A string
	B string
}

type TestB struct {
	B int `validate:"int|in:1,2,3" choices:"B1|B2|B3"`
}

type TestMinMax struct {
	S string `validate:"min:10|max:20"`
}

func TestRenderForm(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestA{
		A: "b",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<input name="A" type="text" value="b"/>
`), html)
}

func TestRenderValidationForm(t *testing.T) {
	f := NewFormLayout().Add("S", "S")
	data := TestMinMax{
		S: "b",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>S</label>
<input name="S" type="text" value="b"/>
`), html)
}

func TestRenderOptionalCheckbox(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestOptionCheckbox{A: nil}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<input type="checkbox" name="A"/>
<label>A</label>
<p></p>
`), html)
}

func TestRenderCheckbox(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestBool{
		A: true,
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<input type="checkbox" name="A" checked=""/>
<label>A</label>
<p></p>
`), html)
}

func TestRenderCheckboxUnchecked(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestBool{
		A: false,
	}
	html := Normalize(Normalize(f.RenderForm(data)))
	assert.Equal(t, RemoveSpacesInHtml(`
<input type="checkbox" name="A"/>
<label>A</label>
<p></p>
`), html)
}

func TestRenderGroup(t *testing.T) {
	f := NewFormLayout().
		AddGroup("A", "X", "Y", func(fl *FormLayout) {
			fl.Add("B", "B", WithDescription("What a description")).
				Add("C", "C").
				Add("D", "D")
		})
	c := 10
	data := TestSubGroup{
		A: TestGroup{B: true, C: &c},
	}
	html := RemoveClassAndStyle(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<a name="formgroup-X"></a>
<div>
<h2>X</h2><p>Y</p>
<div>
<div>
<input type="checkbox" name="A.B" checked=""/>
</div><div>
<label>B</label>
<p>What a description</p>
</div>
</div>
<div>
<label>C</label>
<input name="A.C:int" type="text" value="10"/>
</div>
<div>
<label>D</label>
<input name="A.D:int" type="text" value=""/>
</div>
</div>
`), html)
}

func TestRenderMultiWithDiv(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestMulti{
		A: []string{"A1", "A2"},
	}
	html := RemoveClassAndStyle(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<div>
<div>
<fieldset>
<div>
<div>
<input type="checkbox" name="A#A1" checked=""/>
</div>
<div>
<label>A1</label>
</div>
</div>
<div>
<div>
<input type="checkbox" name="A#A2" checked=""/>
</div>
<div>
<label>A2</label>
</div>
</div>
<div>
<div>
<input type="checkbox" name="A#A3"/>
</div>
<div>
<label>A3</label>
</div>
</div>
</fieldset>
</div>
</div>
`), html)
}

func TestRenderMulti(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A")
	data := TestMulti{
		A: []string{"A1", "A2"},
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`

<fieldset>
<input type="checkbox" name="A#A1" checked=""/>
<label>A1</label>
<input type="checkbox" name="A#A2" checked=""/>
<label>A2</label>
<input type="checkbox" name="A#A3"/>
<label>A3</label>
</fieldset>
`), html)
}

func TestRenderMultiGroup(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A", WithChoices(
			[]Choice{
				{
					Label: "A",
					Value: "A",
					Group: "G1",
				},
				{
					Label: "B",
					Value: "B",
					Group: "G1",
				},
				{
					Label: "C",
					Value: "C",
					Group: "G2",
				},
			}),
			WithGroups(map[string]string{"G1": "Group 1", "G2": "Group 2"}),
		)
	data := TestMulti{
		A: []string{"A", "B"},
	}
	html := RemoveClassAndStyle(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<div>
<div>
<h3>Group 1</h3>
  <fieldset>
<div>
<div>
  <input type="checkbox" name="A#A" checked=""/>
  </div>
  <div>
  <label>A</label>
  </div>
</div>
<div>
<div>
  <input type="checkbox" name="A#B" checked=""/>
  </div>
  <div>
  <label>B</label>
  </div>
</div>
</fieldset>
</div>
<div>
<h3>Group 2</h3>
<fieldset>
<div>
<div>
  <input type="checkbox" name="A#C"/>
  </div>
  <div>
  <label>C</label>
  </div>
</div>
</fieldset>
</div>
</div>
`), html)
}

func TestTwoElementRenderForm(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A").
		Add("B", "B")
	data := TestAB{
		A: "a",
		B: "b",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<input name="A" type="text" value="a"/>
<label>B</label>
<input name="B" type="text" value="b"/>
`), html)
}

func TestTwoElementRenderFormWithError(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A").
		Add("B", "B")
	data := TestAB{
		A: "a",
		B: "b",
	}

	errors := map[string]string{"B": "B not long enough"}
	html := Normalize(f.RenderFormWithErrors(data, errors))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<input name="A" type="text" value="a"/>
<label>B</label>
<input name="B" type="text" value="b"/>
<p>B not long enough</p>
`), html)
}

func TestHeaderRenderForm(t *testing.T) {
	f := NewFormLayout().
		AddHeader("A")
	data := TestA{
		A: "a",
	}
	html := f.RenderForm(data)
	assert.Equal(t, RemoveSpacesInHtml(`
<h2>A</h2>
`), html)
}

func TestRenderSelectForm(t *testing.T) {
	f := NewFormLayout().
		Add("B", "B")
	data := TestB{
		B: 3,
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>B</label>
<select name="B:int">
   <option value="0">-</option>
   <option value="1">B1</option>
   <option value="2">B2</option>
   <option value="3" selected="selected">B3</option>
</select>
`), html)
}

func TestRenderSelectWithChoicesForm(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A", WithChoices([]Choice{
			{
				Label:   "A",
				Value:   "A",
				Checked: false,
			},
			{
				Label:   "B",
				Value:   "B",
				Checked: true,
			},
			{
				Label:   "C",
				Value:   "C",
				Checked: false,
			},
		}))
	data := TestA{
		A: "B",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<select name="A">
   <option value="0">-</option>
   <option value="A">A</option>
   <option value="B" selected="selected">B</option>
   <option value="C">C</option>
</select>
`), html)
}

func TestConfigToHtml(t *testing.T) {
	configHtml := configToHtml(ElementOpts{
		Id:          "id",
		Placeholder: "p",
		Label:       "l",
	})
	assert.Equal(t, " id=\"id\" placeholder=\"p\"", configHtml)
}

func TestRenderFormPlaceHolder(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A", WithPlaceholder("c"))
	data := TestA{
		A: "b",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<input name="A" type="text" value="b" placeholder="c"/>
`), html)
}

func TestRenderFormPlaceId(t *testing.T) {
	f := NewFormLayout().
		Add("A", "A", WithId("c"))
	data := TestA{
		A: "b",
	}
	html := Normalize(f.RenderForm(data))
	assert.Equal(t, RemoveSpacesInHtml(`
<label>A</label>
<input name="A" type="text" value="b" id="c"/>
`), html)
}
