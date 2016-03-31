package main

import (
	"io"
	"text/template"
)

// RenderModel uses the provided model and the given template file
// to perform code generation.
func RenderModel(model *Model, templateFile string, w io.Writer) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, *model)
}
