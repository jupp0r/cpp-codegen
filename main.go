package main

import "flag"

func main() {
	var parseOptions uint16 = CXTranslationUnit_Incomplete | CXTranslationUnit_SkipFunctionBodies | CXTranslationUnit_KeepGoing

	var (
		templateFile, interfaceFile, outFile string
	)
	flag.StringVar(&templateFile, "template", "", "template file for code generation")
	flag.StringVar(&interfaceFile, "interface", "", "interface header file")
	flag.StringVar(&outFile, "output", "", "output file")

	flag.Parse()

	ParseModel(interfaceFile, flag.Args(), parseOptions)
}
