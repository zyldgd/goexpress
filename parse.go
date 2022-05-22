package main

import (
	"encoding/json"
	"github.com/zyldgd/goexpress/ast"
	"github.com/zyldgd/goexpress/token"
)

type parser struct {
	scanner *scanner
	tok     token.Token
	lit     string
	json.Number
}

func ParserAST(expr string) ast.Expr {
	p := &parser{
		scanner: NewScanner(expr),
	}
	p.next()
	e := p.ParseExpr()

	return e
}

// ----------------------------------

func (p *parser) ParseExpr() ast.Expr {
	e := p.parseBinaryExpr(99)
	return e
}

func (p *parser) next() {
	p.tok, p.lit = p.scanner.scan()
}

func (p *parser) parseOperand() ast.Expr {
	var e ast.Expr
	switch p.tok {
	case token.Integer, token.Float, token.Char, token.String:
		e = &ast.LiteralExpr{
			Kind:    p.tok,
			Literal: p.lit,
		}
		p.next()
	case token.ID:
		e = &ast.IDExpr{
			Name: p.lit,
		}
		p.next()
	case token.OpLParen:
		p.next()
		e = p.ParseExpr()
		p.next()
		e = &ast.ParenExpr{E: e}
	}

	switch p.tok {
	case token.OpLBracket:
		p.next()
		index := p.ParseExpr()
		p.next()
		e = &ast.IndexExpr{
			E:     e,
			Index: index,
		}
	case token.OpAccess:
		p.next()
		switch p.tok {
		case token.ID:
			e = &ast.AccessExpr{
				E: e,
				Access: ast.IDExpr{
					Name: p.lit,
				},
			}
		}
		p.next()
	}

	return e
}

func (p *parser) parseUnaryExpr() ast.Expr {
	switch p.tok {
	case token.OpAdd, token.OpMinus, token.OpNot, token.OpBitwiseXor, token.OpBitwiseNot:
		op := p.tok
		p.next()
		e := p.parseUnaryExpr()
		return &ast.UnaryExpr{Op: op, E: e}
	}

	return p.parseOperand()
}

func (p *parser) parseBinaryExpr(pre1 token.Precedence) ast.Expr {
	le := p.parseUnaryExpr()
	// 1 + 2 + 3
	for {
		op := p.tok
		pre := op.Precedence()
		if pre == 0 || !pre.PrecedenceWith(pre1) {
			break
		}
		p.next()
		re := p.parseBinaryExpr(pre)
		le = &ast.BinaryExpr{LE: le, Op: op, RE: re}
	}

	return le
}
