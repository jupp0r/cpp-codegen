package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/go-clang/v3.7/clang"
)

const testInterface = `
#pragma once
#include <string>
#include <vector>
namespace one { namespace two { namespace three {
class TestInterface {
virtual ~TestInterface() = default;
virtual void method(int foo, std::string, const std::vector<int>& t) = 0;
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

	tu := idx.ParseTranslationUnit(file.Name(), []string{"-x", "c++"}, nil, parseOptions)
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()

	model := NewModel()
	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		return visitAST(cursor, parent, &model)
	})

	fmt.Printf("parsed model %v", model)

	classModel, ok := model.Interfaces["TestInterface"]
	if !ok {
		t.Fatalf("TestInterface not found in model interfaces=%v", model.Interfaces)
	}

	if classModel.Name != "TestInterface" {
		t.Fatalf("class name mismatch, expected TestInterface, got %s", classModel.Name)
	}

	if !reflect.DeepEqual(classModel.Namespaces, []string{"one", "two", "three"}) {
		t.Fatalf("namespaces don't match, expected [one, two, three], got %v", classModel.Namespaces)
	}

	methodModel, ok := classModel.Methods["method"]
	if !ok {
		t.Fatalf("method not found in model")
	}

	if methodModel.Arguments[0].Name != "foo" {
		t.Fatalf("expected foo for argument name 0, got %s", methodModel.Arguments[0].Name)
	}

	if methodModel.Arguments[0].Type != "int" {
		t.Fatalf("wrong type for argument, expected int, got %s", methodModel.Arguments[0].Type)
	}

	if methodModel.Arguments[1].Name != "" {
		t.Fatalf("expected no name for %s argument", methodModel.Arguments[1].Name)
	}

	if methodModel.Arguments[1].Type != "std::string" {
		t.Fatalf("wrong type for argument 1, expected std::string, got %s", methodModel.Arguments[1].Type)
	}

	if methodModel.Arguments[2].Name != "t" {
		t.Fatalf("expected t for argument name 2, got %s", methodModel.Arguments[2].Name)
	}

	if methodModel.Arguments[2].Type != "const std::vector<int> &" {
		t.Fatalf("wrong type for argument2 , expected 'const std::vector<int> &', got %s", methodModel.Arguments[2].Type)
	}
}