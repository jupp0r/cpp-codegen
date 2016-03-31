package main

import "github.com/go-clang/v3.7/clang"

func addClassToModel(model *Model, cursor clang.Cursor) error {
	name := cursor.Spelling()
	iface := NewInterface()
	iface.Name = name
	model.Interfaces[name] = iface

	namespaceCursor := cursor.LexicalParent()
	for namespaceCursor.Kind() == clang.Cursor_Namespace {
		iface.Namespaces = append(iface.Namespaces, namespaceCursor.Spelling())
		namespaceCursor = namespaceCursor.LexicalParent()
	}

	var reversedNamespaces []string
	for i := len(iface.Namespaces) - 1; i >= 0; i-- {
		reversedNamespaces = append(reversedNamespaces, iface.Namespaces[i])
	}
	iface.Namespaces = reversedNamespaces

	return nil
}

func addMethodToModel(model *Model, cursor clang.Cursor) error {
	className := cursor.LexicalParent().Spelling()

	method := method{Name: cursor.Spelling(), Arguments: []argument{}}
	model.Interfaces[className].Methods[method.Name] = method
	return nil
}
