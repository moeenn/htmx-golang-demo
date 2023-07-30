package templates

import (
	"html/template"
)

var tmpl *template.Template

func Get() *template.Template {
	var err error
	if tmpl == nil {
		tmpl, err = template.ParseGlob("./pkg/templates/views/*.html")
		if err != nil {
			panic("failed to load template: " + err.Error())
		}
	}

	return tmpl
}
