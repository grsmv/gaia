package main

import (
    "../../gaia"
)

func main() {
   /* gaia.SourceSyntaxTree ("../../gaia/source_files/possible-comands-and-syntax.lisp", false) */
   gaia.SourceSyntaxTree ("../fixtures/000_simple_example.lisp", false)
}

// vim: noai:ts=4:sw=4
