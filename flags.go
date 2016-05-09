package main

// see http://clang.llvm.org/doxygen/Index_8h_source.html

const cxTranslationUnitNone = 0x0
const cxTranslationUnitDetailedPreprocessingRecord = 0x01
const cxTranslationUnitIncomplete = 0x02
const cxTranslationUnitPrecompiledPreamble = 0x04
const cxTranslationUnitCacheCompletionResults = 0x08
const cxTranslationUnitForSerialization = 0x10
const cxTranslationUnitcxXChainedPCH = 0x20
const cxTranslationUnitSkipFunctionBodies = 0x40
const cxTranslationUnitIncludeBriefCommentsInCodeCompletion = 0x80
const cxTranslationUnitCreatePreambleOnFirstParse = 0x100
const cxTranslationUnitKeepGoing = 0x200
