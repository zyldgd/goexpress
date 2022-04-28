package main

import (
	"github.com/zyldgd/goexpress/token"
)

// -----------------------------------------------------------------------------------

type Expr interface {
	string()
}

type BasicLiteral struct {
	Kind    token.Token
	Literal string // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}

type BinaryExpr struct {
	LE Expr
	RE Expr
}

type ParenExpr struct {
	E Expr
}

type UnaryExpr struct {
	Op token.Token
	E  Expr
}

func (*BasicLiteral) string() {}
func (*BinaryExpr) string()   {}
func (*ParenExpr) string()    {}
func (*UnaryExpr) string()    {}

// -----------------------------------------------------------------------------------

func scan(expr string) {
	exprChars := []rune(expr)
	for i := 0; i < len(exprChars); i++ {
		//c := exprChars[i]

	}
}
