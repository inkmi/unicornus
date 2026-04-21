package uni

import (
	"fmt"
	"html"
	"strings"
)

// RenderContext holds rendering state shared across a single render pass
// (output buffer and global flags such as ViewMode). Construct one with
// NewRenderContext; it is not intended for concurrent use.
type RenderContext struct {
	out          *strings.Builder
	AnchorGroups bool
	ViewMode     bool
}

// Safe HTML-escapes str (&, <, >, ", ') for safe embedding in HTML
// attribute values and text nodes.
func Safe(str string) string {
	return html.EscapeString(str)
}

func WriteString(s *strings.Builder, str string) {
	s.WriteString(Safe(str))
}

func WriteValue(s *strings.Builder, value any) {
	WriteString(s, fmt.Sprintf("%v", value))
}

func WriteUnsafeValue(s *strings.Builder, value any) {
	s.WriteString(fmt.Sprintf("%v", value))
}

func (r *RenderContext) TEXTAREA(name string, value any, style string, config ElementOpts) {
	r.out.WriteString("<textarea name=\"")
	r.out.WriteString(Safe(name))
	r.out.WriteString("\" ")
	r.out.WriteString(configToHtml(config))
	r.out.WriteString(" style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
	if config.NoEscape {
		WriteUnsafeValue(r.out, value)
	} else {
		WriteValue(r.out, value)
	}
	r.out.WriteString("</textarea>")
}

func (r *RenderContext) INPUT(name string, typ string, value any, style string, config ElementOpts) {
	r.out.WriteString("<input name=\"")
	r.out.WriteString(Safe(name))
	r.out.WriteString("\" type=\"")
	r.out.WriteString(Safe(typ))
	r.out.WriteString("\" value=\"")
	WriteValue(r.out, value)
	r.out.WriteString("\" ")
	r.out.WriteString(configToHtml(config))
	r.out.WriteString(" style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
}

func (r *RenderContext) DIVv(content string, class ...string) {
	r.out.WriteString("<div class=\"")
	for i := 0; i < len(class); i++ {
		r.out.WriteString(Safe(class[i]))
		r.out.WriteString(" ")
	}
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</div>")
}
func (r *RenderContext) H1no(content string) {
	r.out.WriteString("<h1>")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</h1>")
}

func (r *RenderContext) H2no(content string) {
	r.out.WriteString("<h2>")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</h2>")
}

func (r *RenderContext) DIVclose() {
	r.out.WriteString("</div>")
}

func (r *RenderContext) H2(content string, class string) {
	r.out.WriteString("<h2 class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</h2>")
}

func (r *RenderContext) H3(content string, class string) {
	r.out.WriteString("<h3 class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</h3>")
}

func (r *RenderContext) DIV(content string, class string) {
	r.out.WriteString("<div class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</div>")
}

// DIVRaw writes `content` unescaped. Use ONLY when `content` is already
// a trusted HTML fragment produced by this package. `class` is still escaped.
func (r *RenderContext) DIVRaw(content string, class string) {
	r.out.WriteString("<div class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(content)
	r.out.WriteString("</div>")
}

func (r *RenderContext) DIVS(content string, style string) {
	r.out.WriteString("<div style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</div>")
}

func (r *RenderContext) DIVopen(class string) {
	r.out.WriteString("<div class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
}
func (r *RenderContext) DIVopenS(style string) {
	r.out.WriteString("<div style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
}

func (r *RenderContext) LABEL(content string, class string) {
	r.out.WriteString("<label class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</label>")
}

func (r *RenderContext) LABELS(content string, style string) {
	r.out.WriteString("<label style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</label>")
}

func (r *RenderContext) LABELopenS(style string) {
	r.out.WriteString("<label style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
}
func (r *RenderContext) LABELclose() {
	r.out.WriteString("</label>")

}

func (r *RenderContext) PS(content string, style string) {
	r.out.WriteString("<p style=\"")
	r.out.WriteString(Safe(style))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</p>")
}

func (r *RenderContext) p(content string, class string) {
	r.out.WriteString("<p class=\"")
	r.out.WriteString(Safe(class))
	r.out.WriteString("\">")
	r.out.WriteString(Safe(content))
	r.out.WriteString("</p>")
}

func (r *RenderContext) OnlyDisplay(name string) bool {
	return r.ViewMode
}

type RenderContextFunc func(ctx *RenderContext)

func defaultContext() *RenderContext {
	var sb strings.Builder
	r := RenderContext{
		out:          &sb,
		AnchorGroups: true,
	}
	return &r
}

func WithAnchorGroups(anchorGroups bool) RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.AnchorGroups = anchorGroups
	}
}

func WithDisplayMode() RenderContextFunc {
	return func(ctx *RenderContext) {
		ctx.ViewMode = true
	}
}

func NewRenderContext(config ...RenderContextFunc) *RenderContext {
	c := defaultContext()
	for _, con := range config {
		con(c)
	}
	return c
}

// RenderView renders a read-only HTML representation of data using this
// layout. Use RenderForm for an editable form.
func (f *FormLayout) RenderView(data any) string {
	return f.RenderMultiView([]any{data})
}

// RenderMultiView is like RenderView but merges fields from multiple
// structs into a single read-only view.
func (f *FormLayout) RenderMultiView(data []any) string {
	errors := make(map[string]string)
	fields := make([]DataField, 0)
	for _, d := range data {
		dataFields := FieldGenerator(d, errors)
		fields = append(fields, dataFields...)
	}
	m := FieldsToMap(fields)
	r := NewRenderContext(WithDisplayMode())
	f.renderFormToBuilder(r, "", m)
	return r.out.String()
}

// RenderForm renders an editable HTML form for data using this layout.
// Field values, types, and grouping are discovered by reflection from data.
func (f *FormLayout) RenderForm(data any) string {
	return f.RenderFormMulti([]any{data})
}

// RenderFormMulti is like RenderForm but merges fields from multiple
// structs into a single form.
func (f *FormLayout) RenderFormMulti(
	data []any,
) string {
	errors := make(map[string]string)
	return f.RenderFormWithErrors(errors, data)
}

// RenderFormWithErrors renders an editable form and displays per-field
// validation messages from errors (keyed by field name).
func (f *FormLayout) RenderFormWithErrors(
	errors map[string]string,
	data []any,
) string {
	fields := make([]DataField, 0)
	for _, d := range data {
		dataFields := FieldGenerator(d, errors)
		fields = append(fields, dataFields...)
	}
	m := FieldsToMap(fields)
	r := NewRenderContext()
	f.renderFormToBuilder(r, "", m)
	return r.out.String()
}

func (f *FormLayout) RenderElementWithErrors(
	name string,
	errors map[string]string,
	data ...any,
) string {
	e := f.findByName(name)
	if e == nil {
		return ""
	}
	fields := make([]DataField, 0)
	for _, d := range data {
		fields = append(fields, FieldGenerator(d, errors)...)
	}
	m := FieldsToMap(fields)
	var sb strings.Builder
	r := RenderContext{
		out: &sb,
	}
	// probably prefix taken from name instead of ""
	f.renderElement(*e, &r, "", m)
	return sb.String()
}

func (f *FormLayout) renderFormToBuilder(r *RenderContext, prefix string, m map[string]DataField,
) {
	for _, e := range f.elements {
		f.renderElement(e, r, prefix, m)
	}
}

func prefixed(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

func (f *FormLayout) renderElement(
	e FormElement,
	r *RenderContext,
	prefix string,
	m map[string]DataField,
) {
	switch e.ElementDisplayType {
	case "hidden":
		f.renderHidden(e, r, prefix, m)
	case "header":
		f.Theme.themeRenderHeader(r, e)
	case "group":
		f.renderGroup(e, r, prefix, m)
	case "input":
		f.renderInput(e, r, prefix, m)
	case "textarea":
		f.renderTextareaElement(e, r, prefix, m)
	case "yesno":
		f.renderYesNo(e, r, prefix, m)
	case "dropdown":
		f.renderDropdown(e, r, prefix, m)
	}
}

func (f *FormLayout) renderHidden(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	fieldName := prefixed(prefix, e.Name)
	field, ok := m[fieldName]
	if !ok {
		return
	}
	r.out.WriteString(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"%s\" />", Safe(fieldName), Safe(fmt.Sprintf("%v", field.Val()))))
}

func (f *FormLayout) renderGroup(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	if e.Config.ViewOnly && !r.ViewMode {
		return
	}
	newPrefix := prefixed(prefix, e.Name)
	if r.AnchorGroups {
		r.out.WriteString(fmt.Sprintf("<a name=\"formgroup-%s\"></a>", Safe(stringToAnchor(e.Config.Label))))
	}
	f.Theme.themeRenderGroup(r, m, newPrefix, e)
}

func (f *FormLayout) renderInput(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	fieldName := prefixed(prefix, e.Name)
	field, ok := m[fieldName]
	if !ok {
		return
	}
	if len(e.Config.Choices) > 0 {
		field.Choices = e.Config.Choices
	}
	if field.Multi {
		f.renderMultiInput(e, r, prefix, field)
		return
	}
	f.renderSingleInput(e, r, prefix, field)
}

func (f *FormLayout) renderMultiInput(e FormElement, r *RenderContext, prefix string, field DataField) {
	values := field.Value.([]string)
	for i := 0; i < len(field.Choices); i++ {
		choice := &field.Choices[i]
		if containsString(values, choice.Value) {
			choice.Checked = true
		}
	}
	f.Theme.themeRenderMulti(r, field, e, prefix)
}

func (f *FormLayout) renderSingleInput(e FormElement, r *RenderContext, prefix string, field DataField) {
	description := e.Config.Description
	switch {
	case field.Kind == "bool":
		f.Theme.themeRenderCheckbox(r, e, field, description, prefix)
	case len(field.Choices) > 0:
		f.Theme.themeRenderSelect(r, e, field, description, prefix)
	case field.Kind == "Time":
		f.Theme.themeRenderDateTime(r, e, field, prefix)
	case field.Kind == "Date":
		f.Theme.themeRenderDate(r, e, field, prefix)
	default:
		f.Theme.themeRenderInput(r, e, field, prefix)
	}
}

func (f *FormLayout) renderTextareaElement(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	fieldName := prefixed(prefix, e.Name)
	field, ok := m[fieldName]
	if !ok {
		return
	}
	f.Theme.themeRenderTextarea(r, e, field, e.Config.Description, prefix)
}

func (f *FormLayout) renderYesNo(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	fieldName := prefixed(prefix, e.Name)
	field, ok := m[fieldName]
	if !ok {
		return
	}
	if len(e.Config.Choices) > 0 {
		field.Choices = e.Config.Choices
	}
	f.Theme.themeRenderYesNo(r, e, field, e.Config.Description, prefix)
}

func (f *FormLayout) renderDropdown(e FormElement, r *RenderContext, prefix string, m map[string]DataField) {
	fieldName := prefixed(prefix, e.Name)
	field, ok := m[fieldName]
	if !ok {
		return
	}
	if len(e.Config.Choices) > 0 {
		field.Choices = e.Config.Choices
	}
	f.Theme.themeRenderSelect(r, e, field, e.Config.Description, prefix)
}

func renderCheckboxS(r *RenderContext, f DataField, config ElementOpts, prefix string, style string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := f.Name
	r.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" style=\"%s\" %s%s/>", Safe(name), Safe(style), checked, configToHtml(config)))
}

func renderCheckbox(r *RenderContext, f DataField, config ElementOpts, prefix string, class string) {
	checked := ""
	v, ok := f.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := f.Name
	r.out.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"%s\" %s%s/>", Safe(name), Safe(class), checked, configToHtml(config)))
}

func renderSelect(r *RenderContext, f DataField, config ElementOpts, prefix string, style string, e FormElement) {
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	// optgroup https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
	if len(e.Config.Groups) > 0 {
		r.out.WriteString(fmt.Sprintf("<select name=\"%s\" style=\"%s\"><option value=\"0\">-</option>", Safe(name), Safe(style)))
		for group, groupLabel := range e.Config.Groups {
			r.out.WriteString(fmt.Sprintf("<optgroup LABEL=\"%s\">", Safe(groupLabel)))
			for _, c := range f.Choices {
				if len(group) == 0 || c.Group == group {
					if c.IsSelected(f.Value) {
						r.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", Safe(c.Val()), Safe(c.L())))
					} else {
						r.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", Safe(c.Val()), Safe(c.L())))
					}
				}
			}
			r.out.WriteString("</optgroup>")
		}
		r.out.WriteString("</select>")
	} else {
		r.out.WriteString(fmt.Sprintf("<select name=\"%s\" style=\"%s\"><option value=\"0\">-</option>", Safe(name), Safe(style)))
		for _, c := range f.Choices {
			if c.IsSelected(f.Value) {
				r.out.WriteString(fmt.Sprintf("<option value=\"%s\" selected=\"selected\">%s</option>", Safe(c.Val()), Safe(c.L())))
			} else {
				r.out.WriteString(fmt.Sprintf("<option value=\"%s\">%s</option>", Safe(c.Val()), Safe(c.L())))
			}
		}
		r.out.WriteString("</select>")
	}
}

func renderTextareaS(
	r *RenderContext,
	f DataField,
	val any,
	config ElementOpts,
	style string,
	errorStyle string,
) {
	name := f.Name
	r.TEXTAREA(name, val, style, config)
	if f.HasError() {
		r.PS(f.Errors(), errorStyle)
	}
}

func renderTextInputS(
	r *RenderContext,
	f DataField,
	val any,
	config ElementOpts,
	style string,
	errorStyle string,
) {
	inputType := "text"
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	if f.SubKind == "email" {
		inputType = "email"
	} else if f.SubKind == "url" {
		inputType = "url"
	}

	r.INPUT(name, inputType, val, style, config)
	if f.HasError() {
		r.PS(f.Errors(), errorStyle)
	}
}

func renderDateTimeS(r *RenderContext, f DataField, val any, config ElementOpts, style string, errorStyle string) {
	inputType := "datetime-local"
	name := f.Name + ":time"
	r.INPUT(name, inputType, val, style, config)
	if f.HasError() {
		r.PS(f.Errors(), errorStyle)
	}
}

func renderDateS(r *RenderContext, f DataField, val any, config ElementOpts, style string, errorStyle string) {
	inputType := "date"
	name := f.Name + ":date"
	r.INPUT(name, inputType, val, style, config)
	if f.HasError() {
		r.PS(f.Errors(), errorStyle)
	}
}

func renderTextInput(r *RenderContext, f DataField, val any, config ElementOpts, prefix string, class string) {
	inputConstraints := ""
	inputType := "text"
	name := f.Name
	if f.Kind == "int" {
		name = name + ":int"
	}
	if f.SubKind == "email" {
		inputType = "email"
	} else if f.SubKind == "url" {
		inputType = "url"
	}

	//if f.Name == "Name" {
	//	sb.WriteString(fmt.Sprintf("<div class=\"field_%s\" hx-target-422=\"this\" hx-select=\".field_%s\" hx-target=\"this\" hx-swap=\"outerHTML\" hx-trigger=\"change\" hx-post=\"/ob/onboarding-dreamer?validate=yes\">\n", f.Name, f.Name))
	//	sb.WriteString(fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s value=\"%v\"%s class=\"%s\"/>",
	//		name, inputType,
	//		strings.TrimSpace(inputConstraints), val, configToHtml(config), class))
	//	if f.HasError() {
	//		sb.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", f.Errors()))
	//	}
	//	sb.WriteString("</div>")
	//
	//} else {
	r.out.WriteString(fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s value=\"%s\"%s class=\"%s\"/>",
		Safe(name), Safe(inputType),
		strings.TrimSpace(inputConstraints), Safe(fmt.Sprintf("%v", val)), configToHtml(config), Safe(class)))
	if f.HasError() {
		r.out.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", Safe(f.Errors())))
		//}
	}

}

func configToHtml(config ElementOpts) string {
	id := ""
	if len(config.Id) > 0 {
		id = fmt.Sprintf(" id=\"%s\"", Safe(config.Id))
	}
	placeholder := ""
	if len(config.Placeholder) > 0 {
		placeholder = fmt.Sprintf(" placeholder=\"%s\"", Safe(config.Placeholder))
	}
	configStr := fmt.Sprintf("%s%s", id, placeholder)
	return configStr
}

func containsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

/*

func SetChoices(setKey string, fields []FieldV, allValues []string) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, Choice{
					Label:    p,
					Value:    p,
					Selected: lo.Contains(values, p),
				})
			}

			fields[i].Choices = choices
			fields[i].ElementDisplayType = "string"
		}
	}
}

func SetKey(
	setKey string,
	fields []FieldV,
	allValues []string,
	group func(k string) string,
	LABEL func(l string) string,
) {
	for i := range fields {
		if fields[i].Name == setKey {
			var choices []Choice
			values := fields[i].Value.([]string)
			for _, p := range allValues {
				choices = append(choices, Choice{
					Group:    group(p),
					Label:    LABEL(p),
					Value:    p,
					Selected: lo.Contains(values, p),
				})
			}

			fields[i].Choices = choices
			fields[i].ElementDisplayType = "string"
		}
	}
}

*/
