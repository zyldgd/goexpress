package main

import (
	"go/ast"
	"go/parser"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `abc.1sda`

	// Create the AST by parsing src.
	f, err := parser.ParseExpr(src)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(nil, f)

}
