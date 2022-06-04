package main

import (
	"encoding/json"
	"strconv"
)

type parser struct {
	scanner *scanner
	tok     Token
	lit     string
	json.Number
}

func ParserAST(expr string) Expr {
	p := &parser{
		scanner: NewScanner(expr),
	}
	p.next()
	e := p.ParseExpr()

	return e
}

// ----------------------------------

func (p *parser) ParseExpr() Expr {
	e := p.parseBinaryExpr(99)
	return e
}

func (p *parser) next() {
	p.tok, p.lit = p.scanner.scan()
}

func (p *parser) parseLiteral() Expr {
	var e Expr

	switch p.tok {
	case Integer:
		if data, err := strconv.Atoi(p.lit); err == nil {
			e = &LiteralExpr{
				Kind:    Integer,
				Literal: p.lit,
				Date:    data,
			}
		}
		p.next()
	case Float:
		if data, err := strconv.ParseFloat(p.lit, 32); err == nil {
			e = &LiteralExpr{
				Kind:    Float,
				Literal: p.lit,
				Date:    data,
			}
		}
		p.next()
	case Char:
		if data, _, _, err := strconv.UnquoteChar(p.lit, byte('\'')); err == nil {
			e = &LiteralExpr{
				Kind:    Integer,
				Literal: p.lit,
				Date:    data,
			}
		}
		p.next()
	case String:
		if data, err := strconv.Unquote(p.lit); err == nil {
			e = &LiteralExpr{
				Kind:    Float,
				Literal: p.lit,
				Date:    data,
			}
		}
		p.next()
	}
	return e
}

func (p *parser) parseOperand() Expr {
	var e Expr
	switch p.tok {
	case Ident:
		e = &IdentExpr{
			Name: p.lit,
		}
		p.next()
	case OpLParen:
		p.next()
		e = p.ParseExpr()
		p.next()
		e = &ParenExpr{E: e}
	default:
		e = p.parseLiteral()
	}

	switch p.tok {
	case OpLBracket:
		p.next()
		index := p.ParseExpr()
		p.next()
		e = &IndexExpr{
			E:     e,
			Index: index,
		}
	case OpAccess:
		p.next()
		switch p.tok {
		case Ident:
			e = &AccessExpr{
				E: e,
				Access: IdentExpr{
					Name: p.lit,
				},
			}
		}
		p.next()
	}

	return e
}

func (p *parser) parseUnaryExpr() Expr {
	switch p.tok {
	case OpAdd, OpMinus, OpNot, OpBitwiseXor, OpBitwiseNot:
		op := p.tok
		p.next()
		e := p.parseUnaryExpr()
		return &UnaryExpr{Op: op, E: e}
	}

	return p.parseOperand()
}

func (p *parser) parseBinaryExpr(p0 Precedence) Expr {
	le := p.parseUnaryExpr()
	// 1 + 2 + 3
	for {
		op := p.tok
		p1 := op.Precedence()
		if p1 == 0 || !p1.PrecedenceWith(p0) {
			break
		}
		p.next()
		re := p.parseBinaryExpr(p1)
		le = &BinaryExpr{LE: le, Op: op, RE: re}
	}

	return le
}
