package main

import (
	"fmt"

	"github.com/go-clang/v3.7/clang"
)

func visitAST(cursor clang.Cursor, parent clang.Cursor, model *Model) clang.ChildVisitResult {
	if cursor.IsNull() {
		fmt.Printf("cursor: <none>\n")
		return clang.ChildVisit_Continue
	}

	if !cursor.Location().IsFromMainFile() {
		return clang.ChildVisit_Continue
	}

	fmt.Printf("%s: %s (%s) Type: %s\n", cursor.Kind().Spelling(), cursor.Spelling(), cursor.USR(), cursor.Type().Spelling())

	switch cursor.Kind() {
	case clang.Cursor_Namespace:
		return clang.ChildVisit_Recurse
	case clang.Cursor_ClassDecl, clang.Cursor_StructDecl:
		return visitClassDecl(model, cursor)
	case clang.Cursor_CXXMethod:
		return visitMethod(model, cursor)
	}
	return clang.ChildVisit_Continue
}

func visitClassDecl(model *Model, cursor clang.Cursor) clang.ChildVisitResult {
	err := addClassToModel(model, cursor)
	if err != nil {
		return clang.ChildVisit_Continue
	}
	return clang.ChildVisit_Recurse
}

func visitMethod(model *Model, cursor clang.Cursor) clang.ChildVisitResult {
	if !cursor.CXXMethod_IsPureVirtual() {
		return clang.ChildVisit_Continue
	}

	if cursor.Kind() == clang.Cursor_Destructor {
		return clang.ChildVisit_Continue
	}

	classCursor := cursor.LexicalParent()
	if classCursor.Kind() != clang.Cursor_ClassDecl && classCursor.Kind() != clang.Cursor_StructDecl {
		return clang.ChildVisit_Continue
	}

	err := addMethodToModel(model, cursor)
	if err != nil {
		return clang.ChildVisit_Continue
	}
	return clang.ChildVisit_Continue
}
