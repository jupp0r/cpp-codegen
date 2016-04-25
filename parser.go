package main

import (
	"fmt"

	"github.com/go-clang/v3.7/clang"
)

func ParseModel(file string, args []string, parseOptions uint16) Model {
	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(file, args, nil, uint16(parseOptions))
	defer tu.Dispose()

	fmt.Printf("tu: %s\n", tu.Spelling())
	cursor := tu.TranslationUnitCursor()
	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(file).Name())

	model := NewModel()

	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		return visitAST(cursor, parent, &model)
	})

	return model
}

func addClassToModel(model *Model, cursor clang.Cursor) error {
	name := cursor.Spelling()
	iface := NewInterface()
	iface.Name = name

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
	model.Interfaces[name] = iface

	return nil
}

func addMethodToModel(model *Model, cursor clang.Cursor) error {
	className := cursor.LexicalParent().Spelling()

	method := method{Name: cursor.Spelling(), ReturnType: cursor.ResultType().Spelling(), Arguments: []argument{}}

	for i := int16(0); i < cursor.NumArguments(); i++ {
		argumentCursor := cursor.Argument(uint16(i))
		method.Arguments = append(
			method.Arguments,
			argument{
				Name: argumentCursor.Spelling(),
				Type: argumentCursor.Type().Spelling(),
			})
	}

	model.Interfaces[className].Methods[method.Name] = method
	return nil
}
