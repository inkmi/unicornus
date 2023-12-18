package pkg

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

type TailwindTheme struct {
}

func (t TailwindTheme) themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string, errors map[string]string) {
	sb.WriteString("<div class=\"mt-6\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
	renderTextInput(sb, field, field.Val(), e.Config, prefix, class, errors)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, description string, prefix string, errors map[string]string) {
	sb.WriteString("<div class=\"mt-6\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
	renderSelect(sb, field, e.Config, prefix, class, e)
	if errorMsg, hasError := errors[field.Name]; hasError {
		sb.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", errorMsg))
	}
	sb.WriteString("</div>")
}

/*
<div class="relative flex items-start">
	    <div class="flex h-5 items-center">
	        {{ yield checkbox(  key="EditCompanyCto.CeoTechie" ) }}
	    </div>
	    <div class="ml-3 text-sm">
	        <label for="EditCompanyCto.CeoTechie" class="font-medium text-gray-700">CEO is a techie</label>
	        <p  class="text-gray-500">CEO has not studied economics but is a coder</p>
	    </div>
	</div>
*/

func (t TailwindTheme) themeRenderYesNo(sb *strings.Builder, e FormElement, field DataField, description string, prefix string, errors map[string]string) {
	id := generateRandomID(10)
	checked := ""
	v, ok := field.Val().(bool)
	if ok {
		if v {
			checked = "checked"
		}
	}
	name := field.Name
	sb.WriteString(fmt.Sprintf(`
<div class="mt-6">
  <div class="block pb-3 font-medium text-gray-900">%s</div>
  <label for="%s" class="inline-flex cursor-pointer items-center space-x-4 text-gray-900">
    <span class="text-sm font-medium text-gray-700">%s</span>
    <span class="relative">`, e.Config.Label, id, "No"))
	sb.WriteString(fmt.Sprintf("<input id=\"%s\" type=\"checkbox\" name=\"%s\" class=\"hidden peer\" %s%s/>", id, name, checked,
		configToHtml(e.Config)))
	sb.WriteString(fmt.Sprintf(`
<div class="h-7 w-11 rounded-full shadow-inner bg-gray-200 peer-checked:bg-indigo-500"></div>
      <div class="absolute inset-y-0 left-0 m-1 h-5 w-5 rounded-full 
      ring-1 ring-gray-800
      bg-gray-100 shadow peer-checked:left-auto peer-checked:right-0 peer-checked:bg-gray-100 peer-checked:ring-1 peer-checked:ring-indigo-800"></div>
    </span>
    <span class="text-sm font-medium text-gray-700">%s</span>
  </label>
</div>`, "Yes"))
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

func (t TailwindTheme) themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, description string, prefix string, errors map[string]string) {
	sb.WriteString("<div class=\"py-2 px-4 sm:p-2 lg:pb-4 relative flex items-start\">")
	sb.WriteString("<div class=\"flex h-5 items-center\">")
	class := "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
	renderCheckbox(sb, field, e.Config, prefix, class)
	sb.WriteString("</div>")
	sb.WriteString("<div class=\"ml-3 text-sm\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	sb.WriteString(fmt.Sprintf("<p class=\"text-gray-500\">%s</p>", description))
	if errorMsg, hasError := errors[field.Name]; hasError {
		sb.WriteString(fmt.Sprintf("<p class=\"mt-2 text-sm text-red-600\">%s</p>", errorMsg))
	}
	sb.WriteString("</div>")
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderMulti(sb *strings.Builder, f DataField, e FormElement, prefix string) {
	sb.WriteString("<div class=\"mt-6\">")
	// Should this move to Field generation?
	if len(e.Config.Groups) > 0 {
		for group, name := range e.Config.Groups {
			t.renderMultiGroup(sb, f, group, name)
		}
	} else {
		t.renderMultiGroup(sb, f, "", "")
	}
	sb.WriteString("</div>")
}

func (t TailwindTheme) renderMultiGroup(sb *strings.Builder, f DataField, group string, groupName string) {
	sb.WriteString("<div class=\"mt-6\">")
	if len(groupName) > 0 {
		sb.WriteString(fmt.Sprintf("<h3 class=\"font-bold text-gray-900\">%s</h3>", groupName))
	}
	sb.WriteString("<fieldset class=\"space-y-1\">")
	// range copies slice
	for _, c := range f.Choices {
		if len(group) == 0 || c.Group == group {
			name := f.Name + "#" + c.Val()
			sb.WriteString("<div class=\"relative flex items-start\">")
			sb.WriteString("<div class=\"flex h-5 items-center\">")
			if c.Checked {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" checked class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500\">", name))
			} else {
				sb.WriteString(fmt.Sprintf("<input type=\"checkbox\" name=\"%s\" class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500\">", name))
			}
			sb.WriteString("</div>")
			sb.WriteString("<div class=\"ml-3 text-sm\">")
			sb.WriteString(fmt.Sprintf(`<label class="font-medium text-gray-700">%s</label>`, c.L()))
			sb.WriteString("</div>")
			sb.WriteString("</div>")
		}
	}
	sb.WriteString("</fieldset>")
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderHeader(sb *strings.Builder, e FormElement) {
	sb.WriteString(fmt.Sprintf("<h2>%s</h2>", e.Name))
}

func (t TailwindTheme) themeRenderGroup(sb *strings.Builder, m map[string]DataField, prefix string, e FormElement, errors map[string]string) {
	sb.WriteString("<div class=\"py-6\">")
	sb.WriteString(fmt.Sprintf("<h2 class=\"text-lg leading-6 font-bold text-gray-900\">%s</h2>", e.Label))
	sb.WriteString(fmt.Sprintf("<p class=\"mt-1 text-sm text-gray-500\">%s</p>", e.Description))
	e.Config.SubLayout.renderFormToBuilder(sb, prefix, errors, m)
	sb.WriteString("</div>")
}
