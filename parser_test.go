package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
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
class TestInterface2 {
virtual ~TestInterface() = default;
virtual void method(int foo, std::string, const std::vector<int>& t) = 0;
};
}}}
`

func TestClassParsing(t *testing.T) {
	model := createTestModel(t)

	_, ok := model.Interfaces["TestInterface"]

	if !ok {
		t.Fatalf("TestInterface not found in model interfaces=%v", model.Interfaces)
	}
}

func TestClassNameParsing(t *testing.T) {
	model := createTestModel(t)
	classModel := model.Interfaces["TestInterface"]

	if classModel.Name != "TestInterface" {
		t.Fatalf("class name mismatch, expected TestInterface, got %s", classModel.Name)
	}
}

func TestNamespaceParsing(t *testing.T) {
	model := createTestModel(t)
	classModel := model.Interfaces["TestInterface"]

	if !reflect.DeepEqual(classModel.Namespaces, []string{"one", "two", "three"}) {
		t.Fatalf("namespaces don't match, expected [one, two, three], got %v", classModel.Namespaces)
	}
}

func TestMethodExistsInModel(t *testing.T) {
	model := createTestModel(t)
	classModel := model.Interfaces["TestInterface"]

	_, ok := classModel.Methods["method"]
	if !ok {
		t.Fatalf("method not found in model")
	}
}

func TestMethodArgumentParsing(t *testing.T) {
	methodModel := createMethodModel(t)

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

func TestSecondClassIsParsed(t *testing.T) {
	model := createTestModel(t)

	classModel := model.Interfaces["TestInterface"]

	secondClassModel := model.Interfaces["TestInterface2"]

	if !reflect.DeepEqual(secondClassModel.Methods, classModel.Methods) {
		t.Fatalf("method models should be equal: %v == %v", secondClassModel.Methods, classModel.Methods)
	}

	if !reflect.DeepEqual(secondClassModel.Namespaces, classModel.Namespaces) {
		t.Fatalf("namespace models should be equal: %v == %v", secondClassModel.Namespaces, classModel.Namespaces)
	}
}

func createTestModel(t *testing.T) Model {
	var parseOptions uint16 = CXTranslationUnit_Incomplete | CXTranslationUnit_SkipFunctionBodies | CXTranslationUnit_KeepGoing

	file, err := ioutil.TempFile("", "cpp_codegen_test_parser")
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

	return ParseModel(file.Name(), []string{"-x", "c++"}, parseOptions)
}

func createMethodModel(t *testing.T) method {
	model := createTestModel(t)
	classModel := model.Interfaces["TestInterface"]

	methodModel, ok := classModel.Methods["method"]
	if !ok {
		t.Fatalf("method not found in model")
	}

	return methodModel
}
