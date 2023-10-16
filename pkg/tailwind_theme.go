package pkg

import (
	"fmt"
	"strings"
)

type TailwindTheme struct {
}

func (t TailwindTheme) themeRenderInput(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div class=\"mt-6\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
	renderTextInput(sb, field, field.Val(), e.Config, prefix, class)
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderSelect(sb *strings.Builder, e FormElement, field DataField, prefix string) {
	sb.WriteString("<div class=\"mt-6\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	class := "mt-1 block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
	renderSelect(sb, field, e.Config, prefix, class)
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

func (t TailwindTheme) themeRenderCheckbox(sb *strings.Builder, e FormElement, field DataField, description string, prefix string) {
	sb.WriteString("<div class=\"mt-6 py-6 px-4 sm:p-6 lg:pb-8 relative flex items-start\">")
	sb.WriteString("<div class=\"flex h-5 items-center\">")
	class := "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
	renderCheckbox(sb, field, e.Config, prefix, class)
	sb.WriteString("</div>")
	sb.WriteString("<div class=\"ml-3 text-sm\">")
	if len(e.Config.Label) > 0 {
		sb.WriteString(fmt.Sprintf("<label class=\"block text-sm font-medium text-gray-700\">%s</label>", e.Config.Label))
	}
	sb.WriteString(fmt.Sprintf("<p class=\"text-gray-500\">%s</p>", description))
	sb.WriteString("</div>")
	sb.WriteString("</div>")
}

func (t TailwindTheme) themeRenderMulti(sb *strings.Builder, f DataField, e FormElement, prefix string) {
	sb.WriteString("<div class=\"mt-6\">")
	// Should this move to Field generation?
	if len(e.Config.Groups) > 0 {
		for _, group := range e.Config.Groups {
			t.renderMultiGroup(sb, f, group)
		}
	} else {
		t.renderMultiGroup(sb, f, "")
	}
	sb.WriteString("</div>")
}

func (t TailwindTheme) renderMultiGroup(sb *strings.Builder, f DataField, group string) {
	sb.WriteString("<div>")
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

func (t TailwindTheme) themeRenderGroup(sb *strings.Builder, data any, prefix string, e FormElement) {
	sb.WriteString("<div class=\"py-6 px-4 sm:p-6 lg:pb-8\">")
	sb.WriteString(fmt.Sprintf("<h2 class=\"text-lg leading-6 font-bold text-gray-900\">%s</h2>", e.Label))
	sb.WriteString(fmt.Sprintf("<p class=\"mt-1 text-sm text-gray-500\">%s</p>", e.Description))
	e.Config.SubLayout.renderFormToBuilder(sb, data, prefix)
	sb.WriteString("</div>")
}
