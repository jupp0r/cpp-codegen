package main

import (
	"flag"
	"os"
)

func main() {
	var parseOptions uint16 = cxTranslationUnitIncomplete | cxTranslationUnitSkipFunctionBodies | cxTranslationUnitKeepGoing

	var (
		templateFile, interfaceFile, outFile string
	)
	flag.StringVar(&templateFile, "template", "", "template file for code generation")
	flag.StringVar(&interfaceFile, "interface", "", "interface header file")
	flag.StringVar(&outFile, "output", "", "output file")

	flag.Parse()

	model := ParseModel(interfaceFile, flag.Args(), parseOptions)

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err2 := RenderModel(&model, templateFile, out)
	if err2 != nil {
		panic(err2)
	}
}
