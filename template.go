package main

import (
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

// RenderModel uses the provided model and the given template file
// to perform code generation.
func RenderModel(model *Model, templateFile string, w io.Writer) error {

	helperFuncs := template.FuncMap{
		"join": func(items []string) string { return strings.Join(items, ",") },
		"joinTypes": func(items []argument) string {
			types := []string{}
			for _, argument := range items {
				types = append(types, argument.Type)
			}
			return strings.Join(types, ",")
		},
	}

	tname := filepath.Base(templateFile)
	tmpl, err := template.New(tname).Funcs(helperFuncs).ParseFiles(templateFile)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, tname, *model)
}
