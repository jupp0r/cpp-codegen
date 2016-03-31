package main

import (
	"flag"
	"fmt"

	"github.com/go-clang/v3.7/clang"
)

func main() {
	var (
		templateFile, interfaceFile, outFile string
	)
	flag.StringVar(&templateFile, "template", "", "template file for code generation")
	flag.StringVar(&interfaceFile, "interface", "", "interface header file")
	flag.StringVar(&outFile, "output", "", "output file")

	flag.Parse()

	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(interfaceFile, flag.Args(), nil, 0)
	defer tu.Dispose()

	fmt.Printf("tu: %s\n", tu.Spelling())
	cursor := tu.TranslationUnitCursor()
	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(interfaceFile).Name())

	model := NewModel()

	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		return visitAST(cursor, parent, &model)
	})

	fmt.Printf(":: model: %v\n", model)

	fmt.Printf(":: bye.\n")
}
