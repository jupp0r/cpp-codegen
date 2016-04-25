package main

// see http://clang.llvm.org/doxygen/Index_8h_source.html
const CXTranslationUnit_None = 0x0
const CXTranslationUnit_DetailedPreprocessingRecord = 0x01
const CXTranslationUnit_Incomplete = 0x02
const CXTranslationUnit_PrecompiledPreamble = 0x04
const CXTranslationUnit_CacheCompletionResults = 0x08
const CXTranslationUnit_ForSerialization = 0x10
const CXTranslationUnit_CXXChainedPCH = 0x20
const CXTranslationUnit_SkipFunctionBodies = 0x40
const CXTranslationUnit_IncludeBriefCommentsInCodeCompletion = 0x80
const CXTranslationUnit_CreatePreambleOnFirstParse = 0x100
const CXTranslationUnit_KeepGoing = 0x200
