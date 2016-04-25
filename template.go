package main

import (
	"io"
	"path/filepath"
	"text/template"
)

// RenderModel uses the provided model and the given template file
// to perform code generation.
func RenderModel(model *Model, templateFile string, w io.Writer) error {
	helperFuncs := template.FuncMap{
		"join":      join,
		"joinTypes": joinTypes,
	}

	tname := filepath.Base(templateFile)
	tmpl, err := template.New(tname).Funcs(helperFuncs).ParseFiles(templateFile)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, tname, *model)
}
