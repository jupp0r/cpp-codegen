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

	fmt.Printf("%s: %s (%s)\n", cursor.Kind().Spelling(), cursor.Spelling(), cursor.USR())

	switch cursor.Kind() {
	case clang.Cursor_Namespace:
		return clang.ChildVisit_Recurse
	case clang.Cursor_ClassDecl, clang.Cursor_StructDecl:
		err := addClassToModel(model, cursor)
		if err != nil {
			return clang.ChildVisit_Continue
		}
		return clang.ChildVisit_Recurse
	case clang.Cursor_CXXMethod:
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
		return clang.ChildVisit_Recurse
	case clang.Cursor_ParmDecl:
		argument := argument{Name: cursor.Spelling(), Type: cursor.Type().Spelling()}

		parentMethod := parent.Spelling()
		parentClass := parent.LexicalParent().Spelling()

		method := model.Interfaces[parentClass].Methods[parentMethod]
		methodArguments := append(model.Interfaces[parentClass].Methods[parentMethod].Arguments, argument)
		method.Arguments = methodArguments
		model.Interfaces[parentClass].Methods[parentMethod] = method

		return clang.ChildVisit_Continue
	}
	return clang.ChildVisit_Continue
}
