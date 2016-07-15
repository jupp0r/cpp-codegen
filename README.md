# cpp codegen [![Build Status](https://travis-ci.org/jupp0r/cpp-codegen.svg?branch=master)](https://travis-ci.org/jupp0r/cpp-codegen) [![Windows Build status](https://ci.appveyor.com/api/projects/status/921lubl0gg04og10/branch/master?svg=true)](https://ci.appveyor.com/project/jupp0r/cpp-codegen/branch/master)

A Model View Controller for Source-To-Source
transformations. This tool reads C++ interfaces via libclang, and
transforms them according to specific recipes. One use case is the
automatic generation of mock objects for the (googletest
library)[https://github.com/google/googletest] (recipe provided).

# Usage

    Usage of ./cpp-codegen:
      -interface string
          interface header file
      -output string
          output file
      -template string
          template file for code generation

# License
MIT
