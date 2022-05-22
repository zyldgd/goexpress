package main

import (
	"errors"
	"fmt"
	"github.com/zyldgd/goexpress/ast"
	"github.com/zyldgd/goexpress/token"
	"strconv"
)

type Expression struct {
	Expr   ast.Expr
	Params map[string]interface{}
}

func NewExpression(expression string) *Expression {
	expr := ParserAST(expression)
	return &Expression{
		Expr: expr,
	}
}

func (e *Expression) Calc(params map[string]interface{}) *Result {
	e.Params = params
	if e.Params == nil {
		e.Params = make(map[string]interface{}, 2)
	}
	e.Params["true"] = true
	e.Params["false"] = false

	return e.calc()
}

func (e *Expression) calc() *Result {
	result, err := e.calcExpr(e.Expr)
	if err != nil {
		return nil
	}

	return result
}

func (e *Expression) calcExpr(expr ast.Expr) (*Result, error) {
	switch ex := expr.(type) {
	case *ast.LiteralExpr:
		return e.calcLiteralExpr(ex)
	case *ast.IDExpr:
		return e.calcIDExpr(ex)
	case *ast.UnaryExpr:
		return e.calcUnaryExpr(ex)
	case *ast.BinaryExpr:
		return e.calcBinaryExpr(ex)
	case *ast.ParenExpr:
		return e.calcParenExpr(ex)
	}
	return nil, errors.New("token error")
}

func (e *Expression) calcIDExpr(expr *ast.IDExpr) (*Result, error) {
	if val, find := e.Params[expr.Name]; find {
		switch v := val.(type) {
		case int:
			result := &Result{
				kind: token.Integer,
				data: v,
			}
			return result, nil
		case int8:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case int16:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case int32:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case uint:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case uint8:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case uint16:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case uint32:
			result := &Result{
				kind: token.Integer,
				data: int(v),
			}
			return result, nil
		case float64:
			result := &Result{
				kind: token.Float,
				data: float32(v),
			}
			return result, nil
		case float32:
			result := &Result{
				kind: token.Float,
				data: v,
			}
			return result, nil
		case string:
			result := &Result{
				kind: token.String,
				data: v,
			}
			return result, nil
		case bool:
			result := &Result{
				kind: token.ID,
				data: v,
			}
			return result, nil
			//case func():

		}
	}
	return nil, errors.New("token error")
}

func (e *Expression) calcParenExpr(expr *ast.ParenExpr) (*Result, error) {
	return e.calcExpr(expr.E)
}

func (e *Expression) calcBinaryExpr(expr *ast.BinaryExpr) (*Result, error) {
	switch expr.Op {
	case token.OpAdd:
		l, err := e.calcExpr(expr.LE)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.RE)
		if err != nil {
			return nil, err
		}

		if l.kind == token.Integer && r.kind == token.Integer {
			data := l.data.(int) + r.data.(int)
			result := &Result{
				kind: token.Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Float {
			data := l.data.(float32) + r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Integer && r.kind == token.Float {
			data := float32(l.data.(int)) + r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Integer {
			data := l.data.(float32) + float32(r.data.(int))
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.String && r.kind == token.String {
			data := l.data.(string) + r.data.(string)
			result := &Result{
				kind: token.String,
				data: data,
			}
			return result, nil
		}
	case token.OpMinus:
		l, err := e.calcExpr(expr.LE)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.RE)
		if err != nil {
			return nil, err
		}

		if l.kind == token.Integer && r.kind == token.Integer {
			data := l.data.(int) - r.data.(int)
			result := &Result{
				kind: token.Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Float {
			data := l.data.(float32) - r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Integer && r.kind == token.Float {
			data := float32(l.data.(int)) + r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Integer {
			data := l.data.(float32) - float32(r.data.(int))
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		}
	case token.OpMultiply:
		l, err := e.calcExpr(expr.LE)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.RE)
		if err != nil {
			return nil, err
		}

		if l.kind == token.Integer && r.kind == token.Integer {
			data := l.data.(int) * r.data.(int)
			result := &Result{
				kind: token.Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Float {
			data := l.data.(float32) * r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Integer && r.kind == token.Float {
			data := float32(l.data.(int)) * r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Integer {
			data := l.data.(float32) * float32(r.data.(int))
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		}
	case token.OpDivide:
		l, err := e.calcExpr(expr.LE)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.RE)
		if err != nil {
			return nil, err
		}

		if l.kind == token.Integer && r.kind == token.Integer {
			data := l.data.(int) / r.data.(int)
			result := &Result{
				kind: token.Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Float {
			data := l.data.(float32) / r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Integer && r.kind == token.Float {
			data := float32(l.data.(int)) / r.data.(float32)
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		} else if l.kind == token.Float && r.kind == token.Integer {
			data := l.data.(float32) / float32(r.data.(int))
			result := &Result{
				kind: token.Float,
				data: data,
			}
			return result, nil
		}
	case token.OpModulus:
		l, err := e.calcExpr(expr.LE)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.RE)
		if err != nil {
			return nil, err
		}
		if l.kind == token.Integer && r.kind == token.Integer {
			data := l.data.(int) % r.data.(int)
			result := &Result{
				kind: token.Integer,
				data: data,
			}
			return result, nil
		}
	}

	return nil, errors.New("token error")
}

func (e *Expression) calcLiteralExpr(expr *ast.LiteralExpr) (*Result, error) {
	switch expr.Kind {
	case token.Integer:
		if data, err := strconv.Atoi(expr.Literal); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.Kind,
				data: data,
			}
			return result, nil
		}
	case token.Float:
		if data, err := strconv.ParseFloat(expr.Literal, 32); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.Kind,
				data: float32(data),
			}
			return result, nil
		}
	case token.Char:
		if len(expr.Literal) == 3 {
			result := &Result{
				kind: expr.Kind,
				data: rune(expr.Literal[1]),
			}
			return result, nil
		} else {
			return nil, errors.New("parse char error")
		}
	case token.String:
		if data, err := strconv.Unquote(expr.Literal); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.Kind,
				data: data,
			}
			return result, nil
		}
	}

	return nil, errors.New("token error")
}

func (e *Expression) calcUnaryExpr(expr *ast.UnaryExpr) (*Result, error) {
	switch expr.Op {
	case token.OpAdd:
		if result, err := e.calcExpr(expr.E); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case token.Integer:
				data := +(result.data.(int))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			case token.Float:
				data := +(result.data.(float32))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			default:
				return nil, errors.New("token error")
			}
		}
	case token.OpMinus:
		if result, err := e.calcExpr(expr.E); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case token.Integer:
				data := -(result.data.(int))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			case token.Float:
				data := -(result.data.(float32))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			default:
				return nil, errors.New("token error")
			}
		}

	case token.OpNot:
		if result, err := e.calcExpr(expr.E); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case token.ID:
				if data, ok := result.data.(bool); ok {
					result := &Result{
						kind: result.kind,
						data: data,
					}
					return result, nil
				}
			default:
				return nil, errors.New("token error")
			}
		}
	case token.OpBitwiseXor:
		// TODO
	case token.OpBitwiseNot:
		// TODO
	}
	return nil, errors.New("token error")
}

// --------------------------------------------------------------------------

type Result struct {
	kind token.Token
	data interface{}
}

//
//func (e result) Int() (int, error) {
//
//}
//
//func (e result) Float() (float32, error) {
//
//}
//
//func (e result) Bool() (bool, error) {
//
//}
//
//func (e result) String() (string, error) {
//
//}
//
//func (e result) Char() (rune, error) {
//
//}

func main() {
	expr := NewExpression("a + 1 * b")
	result := expr.Calc(map[string]interface{}{"a": 89.9, "b": 2})
	fmt.Printf("%+v \n", result)
}
