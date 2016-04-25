package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-clang/v3.7/clang"
)

const testInterface = `
#pragma once
namespace one { namespace two { namespace three {
class TestInterface {
virtual ~TestInterface() = default;
virtual void method(int, string, const& TemplatedType<int>) = 0;
};
}}}
`

func TestParserWithSimpleInterface(t *testing.T) {
	var parseOptions uint16 = CXTranslationUnit_Incomplete | CXTranslationUnit_SkipFunctionBodies | CXTranslationUnit_KeepGoing

	file, err := ioutil.TempFile("", "cpp_codegen_test_parser")
	defer os.Remove(file.Name())
	if err != nil {
		t.Fatalf("Unable to create temporary file")
	}
	defer os.Remove(file.Name())

	_, err4 := file.Write([]byte(testInterface))
	if err4 != nil {
		t.Fatalf("Unable to write to file %s", file.Name())
	}

	err3 := file.Close()
	if err3 != nil {
		t.Fatalf("Unable to close file %s", file.Name())
	}

	// TODO @jupp: this is copy + pasted from main.go, extract this and hide behind an interface
	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	fmt.Printf("opened %s\n", file.Name())

	tu := idx.ParseTranslationUnit(file.Name(), []string{"-x", "c++"}, nil, parseOptions)
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()
	fmt.Printf("tu: %s\n", tu.Spelling())

	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(file.Name()).Name())

	model := NewModel()
	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		return visitAST(cursor, parent, &model)
	})

	classModel, ok := model.Interfaces["TestInterface"]
	if !ok {
		t.Fatalf("TestInterface not found in model interfaces=%v", model.Interfaces)
	}

	_, ok2 := classModel.Methods["method"]
	if !ok2 {
		t.Fatalf("method not found in model")
	}
}
