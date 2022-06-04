package main

import (
	"encoding/json"
	"fmt"
)

type Expr interface {
	String() string
}
type (
	LiteralExpr struct {
		Kind    main.Token  `json:"kind"`
		Literal string      `json:"literal"`
		Date    interface{} `json:"data"`
	}

	AccessExpr struct {
		E      Expr      `json:"e"`
		Access IdentExpr `json:"access"`
	}

	IndexExpr struct {
		E     Expr `json:"e"`
		Index Expr `json:"index"`
	}

	IdentExpr struct {
		Name string `json:"name"`
	}

	BinaryExpr struct {
		LE Expr       `json:"le"`
		Op main.Token `json:"op"`
		RE Expr       `json:"re"`
	}

	ParenExpr struct {
		E Expr `json:"e"`
	}

	UnaryExpr struct {
		Op main.Token `json:"op"`
		E  Expr       `json:"e"`
	}
)

func (e *LiteralExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *AccessExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *IndexExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *IdentExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *BinaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *ParenExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *UnaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func PrintAst(e Expr) {
	b, _ := json.MarshalIndent(e, "", "    ")

	fmt.Printf("ast:\n%+v\n", string(b))
}

// -----------------------------------------------------------------------------------

func scan(expr string) {
	exprChars := []rune(expr)
	for i := 0; i < len(exprChars); i++ {
		//c := exprChars[i]

	}
}
