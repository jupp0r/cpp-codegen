package main

import "strings"

func join(items []string) string {
	return strings.Join(items, ", ")
}

func joinTypes(items []argument) string {
	types := []string{}
	for _, argument := range items {
		types = append(types, argument.Type)
	}
	return strings.Join(types, ", ")
}
