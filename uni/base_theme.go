package uni

import (
	"crypto/rand"
	"fmt"
	"log"
	"regexp"
)

type StyleFunc func(style *ThemeStyles)

func defaultStyle() *ThemeStyles {
	return &ThemeStyles{}
}

func NewStyles(ops ...StyleFunc) *ThemeStyles {
	style := defaultStyle()
	for _, sf := range ops {
		sf(style)
	}
	return style
}

func ErrorStyle(errorStyle string) StyleFunc {
	return func(t *ThemeStyles) {
		t.errorStyle = errorStyle
	}
}

func LabelStyle(labelStyle string) StyleFunc {
	return func(t *ThemeStyles) {
		t.labelStyle = labelStyle
	}
}

func InputStyle(inputStyle string) StyleFunc {
	return func(t *ThemeStyles) {
		t.inputStyle = inputStyle
	}
}

func TopSeparator(separator string) StyleFunc {
	return func(t *ThemeStyles) {
		t.topSeparator = "margin-top: " + separator + ";"
	}
}

type ThemeStyles struct {
	topSeparator string
	errorStyle   string
	labelStyle   string
	inputStyle   string
}

type BaseTheme struct {
	styles *ThemeStyles
}

func (t BaseTheme) themeRenderInput(r *RenderContext, e FormElement, field DataField, prefix string) {
	r.DIVopenS(t.styles.topSeparator)
	if r.OnlyDisplay(field.Name) {
		if len(e.Config.Label) > 0 {
			r.DIVS(e.Config.Label, t.styles.labelStyle)
		}
		r.DIV(Safe(field.ViewVal()), "font-size: 0.875rem; font-weight: 500; color: #1F2937;")
	} else {
		// Label for input
		if len(e.Config.Label) > 0 {
			r.LABELS(e.Config.Label, t.styles.labelStyle)
		}
		// Render input element
		renderTextInputS(r, field, field.Val(), e.Config, t.styles.inputStyle, t.styles.errorStyle)
		// Render description
		if len(e.Config.Description) > 0 {
			r.PS(e.Config.Description, "margin-top: 0.5rem; color: #6B7280; ")
		}
	}
	r.DIVclose()
}

func (t BaseTheme) themeRenderSelect(r *RenderContext, e FormElement, field DataField, description string, prefix string) {
	r.DIVopenS(t.styles.topSeparator)
	if r.OnlyDisplay(field.Name) {
		if len(e.Config.Label) > 0 {
			r.DIVS(e.Config.Label, t.styles.labelStyle)
		}
		r.DIV(Safe(field.ViewVal()), "font-weight: 500; color: #1F2937;")
	} else {
		if len(e.Config.Label) > 0 {
			r.LABELS(e.Config.Label, t.styles.labelStyle)
		}
		style := "margin-top: 4px; display: block; width: 100%; border-radius: 6px; border: 1px solid #D1D5DB; padding: 8px 40px 8px 12px; font-size: 16px;"
		renderSelect(r, field, e.Config, prefix, style, e)
		if field.HasError() {
			r.PS(field.Errors(), t.styles.errorStyle)
		}
		if len(description) > 0 {
			r.PS(description, "margin-top: 0.5rem; color: #6B7280; ")
		}
	}
	r.DIVclose()
}

func (t BaseTheme) themeRenderYesNo(r *RenderContext, e FormElement, field DataField, description string, prefix string) {
	id := generateRandomID(10)
	checked := ""
	v, ok := field.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := field.Name
	r.out.WriteString(fmt.Sprintf(`
<div class="mt-6">
  <div class="block pb-3 font-medium text-gray-900">%s</div>
  <label for="%s" class="inline-flex cursor-pointer items-center space-x-4 text-gray-900">
    <span class="text-sm font-medium text-gray-700">%s</span>
    <span class="relative">`, e.Config.Label, id, "No"))
	r.out.WriteString(fmt.Sprintf("<input id=\"%s\" type=\"checkbox\" name=\"%s\" class=\"hidden peer\" %s%s/>", id, name, checked,
		configToHtml(e.Config)))
	r.out.WriteString(fmt.Sprintf(`
<div class="h-7 w-11 rounded-full shadow-inner bg-gray-200 peer-checked:bg-indigo-500"></div>
      <div class="absolute inset-y-0 left-0 m-1 h-5 w-5 rounded-full 
      ring-1 ring-gray-800
      bg-gray-100 shadow peer-checked:left-auto peer-checked:right-0 peer-checked:bg-gray-100 peer-checked:ring-1 peer-checked:ring-indigo-800"></div>
    </span>
    <span class="text-sm font-medium text-gray-700">%s</span>
  </label>
</div>`, "Yes"))
}

func (t BaseTheme) themeRenderCheckbox(r *RenderContext, e FormElement, field DataField, description string, prefix string) {
	r.DIVopenS("display: flex; padding: 8px 16px; align-items: flex-start;")
	r.DIVopenS("display: flex; height: 20px; align-items: center;")
	if r.OnlyDisplay(field.Name) {
		v, ok := field.Val().(bool)
		if ok {
			if v {
				r.out.WriteString("[x]")
			} else {
				r.out.WriteString("[ ]")
			}
		}
	} else {
		if len(e.Config.Label) > 0 {
			r.LABELopenS(t.styles.labelStyle)
		}
		style := "height: 16px; vertical-align: -4px; width: 16px; border-radius: 4px; border: 1px solid #D1D5DB; color: #4F46E5;"
		renderCheckboxS(r, field, e.Config, prefix, style)
		if len(e.Config.Label) > 0 {
			r.DIVS(e.Config.Label, "display: inline; padding-left: 0.5rem;")
			r.LABELclose()
		}
	}
	r.DIVclose()
	r.DIVopenS("margin-left: 12px; font-size: 14px;")
	if len(description) > 0 {
		r.PS(description, "margin-left: 12px; font-size: 14px;")
	}
	if field.HasError() {
		r.PS(field.Errors(), t.styles.errorStyle)
	}
	r.DIVclose()
	r.DIVclose()
}

func (t BaseTheme) themeRenderMulti(r *RenderContext, f DataField, e FormElement, prefix string) {
	r.DIVopenS(t.styles.topSeparator)
	// Should this move to Field generation?
	if len(e.Config.Groups) > 0 {
		for group, name := range e.Config.Groups {
			t.renderMultiGroup(r, f, group, name)
		}
	} else {
		t.renderMultiGroup(r, f, "", "")
	}
	r.DIVclose()
}

func (t BaseTheme) renderMultiGroup(r *RenderContext, f DataField, group string, groupName string) {
	r.DIVopenS(t.styles.topSeparator)
	if len(groupName) > 0 {
		r.H3(groupName, "font-bold text-gray-900")
	}
	r.out.WriteString("<fieldset class=\"space-y-1\">")
	// range copies slice
	for _, c := range f.Choices {
		if len(group) == 0 || c.Group == group {
			name := f.Name + "#" + c.Val()
			r.DIVopen("relative flex items-start")
			r.DIVopen("flex h-5 items-center")
			r.LABELopenS(t.styles.labelStyle)
			if c.Checked {
				r.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500\">", name))
			} else {
				r.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500\">", name))
			}
			r.DIVS(c.L(), "display: inline; padding-left: 0.5rem;")
			r.LABELclose()
			r.DIVclose()
			r.DIVclose()
		}
	}
	r.out.WriteString("</fieldset>")
	r.DIVclose()
}

func (t BaseTheme) themeRenderHeader(r *RenderContext, e FormElement) {
	r.H2no(e.Name)
}

func (t BaseTheme) themeRenderGroup(r *RenderContext, m map[string]DataField, prefix string, e FormElement) {
	r.DIVopenS(t.styles.topSeparator)
	r.H2(e.Config.Label, "text-lg leading-6 font-bold text-gray-900")
	r.p(e.Config.Description, "mt-1 text-sm text-gray-500")
	e.Config.SubLayout.renderFormToBuilder(r, prefix, m)
	r.DIVclose()
}

func stringToAnchor(input string) string {
	// Replace multiple spaces with a single dash
	spaceRegex := regexp.MustCompile(`\s+`)
	result := spaceRegex.ReplaceAllString(input, "-")

	// Remove non-alphanumeric characters
	alphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9-]`)
	result = alphanumericRegex.ReplaceAllString(result, "")

	return result
}

func generateRandomID(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
