package main

import "testing"

func TestJoinSingleElement(t *testing.T) {
	result := join([]string{"foo"})
	if result != "foo" {
		t.Fatalf("error joining single value, expected 'foo', got %s", result)
	}
}

func TestJoinMultipleElements(t *testing.T) {
	result := join([]string{"foo", "bar", "baz"})
	expected := "foo, bar, baz"
	if result != expected {
		t.Fatalf("error joining single value, expected '%s', got '%s'", expected, result)
	}
}

func TestJoinArgumentTypes(t *testing.T) {
	arguments := []argument{
		argument{
			Name: "i",
			Type: "const int &",
		},
		argument{
			Name: "j",
			Type: "std::string",
		},
	}

	result := joinTypes(arguments)
	expected := "const int &, std::string"

	if result != expected {
		t.Fatalf("error joining argument types, expected '%s', got '%s'", expected, result)
	}
}
