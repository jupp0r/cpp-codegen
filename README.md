# cpp codegen
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
