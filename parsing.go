package goexpress

import (
	"container/list"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Cond.And(Eq{}).

/*
${ls.variate1} == 1 || ${ls.variate2} != 2 && (${ls.variate1} == 2 && ${ls.variate2} != 1) && !(${ls.variate3} == 3)
 ${ls.variate1} == 1
  &&



*/

/*
ls.popular_tab = "Test1(OM2)"
OR
ls.popular_tab = "Test2(OM3)"

<recommend-cond>
 ${ls.variate1} == 1 && ${ls.variate2} != 2 || (${ls.variate1} == 2 && ${ls.variate2} != 1) && !(${ls.variate3} == 3)
</recommend-cond>

simple expression： Eq: == , Neq: !=  Not:!
 {value} {symbol} {value}


-------------------------------------------------------
parse(expression string) expression, err

Cond(expression string, params map[string]interface{}) bool, err

*/

/*

<recommend-cond>

</recommend-cond>

*/
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func (s *Stack) PushList(list ...interface{}) {
	if s != nil {
		for _, e := range list {
			s.list.PushBack(e)
		}
	}
}

func (s *Stack) Push(elem interface{}) {
	if s != nil {
		s.list.PushBack(elem)
	}
}

func (s *Stack) Pop() interface{} {
	if s != nil {
		e := s.list.Back()
		if e != nil {
			s.list.Remove(e)
			return e.Value
		}
	}
	return nil
}

func (s *Stack) Peak() interface{} {
	if s != nil {
		e := s.list.Back()
		if e != nil {
			return e.Value
		}
	}
	return nil
}

func (s *Stack) Len() int {
	if s != nil {
		return s.list.Len()
	}
	return 0
}

func (s *Stack) IsEmpty() bool {
	if s != nil {
		return s.list.Len() == 0
	}
	return true
}

// ------------------------------------------

const VarPrefix = "${"
const VarSuffix = "}"

// ------------------------------------------

type Operator string

const (
	OpEq         Operator = "=="
	OpNeq        Operator = "!="
	OpAnd        Operator = "&&"
	OpOr         Operator = "||"
	OpNot        Operator = "!"
	OpOpenParen  Operator = "("
	OpCloseParen Operator = ")"
)

// 值约小，优先级约高
func (op Operator) Priority() int {
	switch op {
	case OpOr:
		return 12
	case OpAnd:
		return 11
	case OpEq, OpNeq:
		return 7
	case OpNot:
		return 2
	case OpOpenParen, OpCloseParen:
		return 1
	}
	return 999
}

func (op Operator) len() int {
	return len(op)
}

func (op Operator) eq(str string) bool {
	return string(op) == str
}

//func (op Operator) string() string {
//	return []rune(op)
//}

func (op Operator) PriorTo(op2 Operator) bool {
	return op.Priority() < op2.Priority()
}

// ------------------------------------------
type Variable string
type LiteralValue struct {
	Literal string
	Value   interface{}
	Kind    reflect.Kind
}

func (v Variable) name() string {
	return string(v)
}

// ------------------------------------------
type Booler interface {
	Bool() bool
}

type BoolValue bool

func (e BoolValue) Bool() bool {
	return bool(e)
}

// ------------------------------------------
type SimpleExpression struct {
	Expression string
	Values     []BoolValue
	Operator   *Operator
}

func (e SimpleExpression) Bool() bool {
	if e.Operator == nil {
		return e.Values[0].Bool()
	}

	switch *e.Operator {
	case OpNot:
		return !e.Values[0].Bool()
	case OpAnd:
		return e.Values[0].Bool() && e.Values[1].Bool()
	case OpOr:
		return e.Values[0].Bool() || e.Values[1].Bool()
	}
	return false
}

// ------------------------------------------
const (
	StatusStart      = 0
	StatusOfValue    = 1
	StatusOfOperator = 2
)

type BoolExpression struct {
	RawExpression     string
	scanner           *Scanner
	postfixExpression []interface{}
}

type Scanner struct {
	expression []rune
	prefix     string
	suffix     string
	index      int
	elements   []interface{}
}

func (s *Scanner) append(e interface{}) error {
	//fmt.Printf("substr[%s] ------------\n", e)
	//if len(s.elements) == 0 {
	//	s.elements = append(s.elements, e)
	//} else {
	//	last := s.elements[len(s.elements)-1]
	//	lType := reflect.TypeOf(last).Kind()
	//	eType := reflect.TypeOf(e).Kind()
	//	switch last.(type) {
	//	case LiteralValue:
	//		if eType == lType {
	//			return fmt.Errorf("substr[%s] invalid", eType)
	//		}
	//	}
	//}
	s.elements = append(s.elements, e)
	return nil
}

func (s *Scanner) tryAppendOperator(op Operator) error {
	if op.len()+s.index > len(s.expression) {
		return fmt.Errorf("out of bound")
	} else {
		if s.index+op.len() > len(s.expression) || !op.eq(string(s.expression[s.index:s.index+op.len()])) {
			return fmt.Errorf("out of bound or operator[%s] not eq", op)
		}
	}

	if err := s.append(op); err != nil {
		return err
	}

	s.index += op.len() - 1
	return nil
}

func (s *Scanner) subScan() error {
	var i = s.index
	var startC = s.expression[s.index]
	var builder strings.Builder
	var quoted = startC == '"'

	if quoted {
		builder.WriteRune(startC)
		for i++; i < len(s.expression); i++ {
			c := s.expression[i]
			builder.WriteRune(c)
			if c == '\\' {
				if i+1 < len(s.expression) {
					next := s.expression[i+1]
					builder.WriteRune(next)
					i++
					continue
				} else {
					return fmt.Errorf("escape error")
				}
			} else if c == '"' {
				quoted = false
				str, err := strconv.Unquote(builder.String())
				if err != nil {
					return fmt.Errorf("unquote error: %s", err)
				}
				if err = s.append(LiteralValue{
					Literal: str,
					Value:   str,
					Kind:    reflect.String,
				}); err != nil {
					return err
				}
				break
			}
		}
		if quoted {
			return fmt.Errorf("not end with quote")
		}
	} else {
		for ; i < len(s.expression); i++ {
			c := s.expression[i]
			if c == '!' || c == '=' || c == '&' || c == '|' || c == '(' || c == ')' || c == ' ' {
				str := builder.String()
				if len(str) == 0 {
					return fmt.Errorf("empty str")
				} else {
					pl := len(s.prefix)
					sl := len(s.suffix)
					if len(str) > pl+sl && s.prefix == str[0:pl] && s.suffix == str[len(str)-sl:] {
						varStr := str[pl : len(str)-sl]
						if err := s.append(Variable(varStr)); err != nil {
							return err
						}
					} else {
						if err := s.append(value(str)); err != nil {
							return err
						}
					}
				}

				i-- // 回退
				break
			} else {
				builder.WriteRune(c)
			}
		}
	}

	s.index = i
	return nil
}

func (s *Scanner) scan() (err error) {
	for s.index = 0; s.index < len(s.expression) && err == nil; s.index++ {
		switch s.expression[s.index] {
		case '(':
			err = s.tryAppendOperator(OpOpenParen)
		case ')':
			err = s.tryAppendOperator(OpCloseParen)
		case '&':
			err = s.tryAppendOperator(OpAnd)
		case '|':
			err = s.tryAppendOperator(OpOr)
		case '!':
			if s.index+1 < len(s.expression) {
				if s.expression[s.index+1] == '=' {
					err = s.tryAppendOperator(OpNeq)
				} else {
					err = s.tryAppendOperator(OpNot)
				}
			} else {
				err = fmt.Errorf("out of bound")
			}
		case '=':
			err = s.tryAppendOperator(OpEq)
		case ' ': // 空格忽略
		default:
			err = s.subScan()
		}
	}

	if err != nil {
		return fmt.Errorf("syntax error at %d char, %s", s.index, err)
	}
	return err
}

func isVariateChar(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' || c == '_' || c == '.'

}

func isOperatorChar(c rune) bool {
	return c == '&' || c == '|' || c == '!'
}

func isParenChar(c rune) bool {
	return c == '(' || c == ')'
}

func value(str string) LiteralValue {
	lv := LiteralValue{
		Literal: str,
	}

	if str == "true" {
		lv.Value = true
		lv.Kind = reflect.Bool
	} else if str == "false" {
		lv.Value = false
		lv.Kind = reflect.Bool
	} else if vi, err := strconv.ParseInt(str, 10, 0); err == nil {
		lv.Value = vi
		lv.Kind = reflect.Int64
	} else if vf, err := strconv.ParseFloat(str, 64); err == nil {
		lv.Value = vf
		lv.Kind = reflect.Float64
	} else {
		lv.Value = str
		lv.Kind = reflect.String
	}

	return lv
}

// ${var1} == 1 && (${var2} != 2 ||  ${var3} == 3)

func (exp *BoolExpression) init() error {
	exp.scanner.expression = []rune(exp.RawExpression)
	exp.scanner.index = 0
	exp.scanner.prefix = VarPrefix
	exp.scanner.suffix = VarSuffix

	if err := exp.scanner.scan(); err != nil {
		return err
	}

	if err := exp.toPostfixExpression(); err != nil {
		return err
	}

	return nil
}

func (exp *BoolExpression) toPostfixExpression() error {
	elements := exp.scanner.elements

	s1 := NewStack()
	s2 := NewStack()

loop:
	for _, e := range elements {
		if _, ok := e.(Operator); !ok {
			s1.Push(e)
			continue
		}

		switch e {
		case OpNot, OpOr, OpAnd, OpEq, OpNeq:
		start:
			if s2.IsEmpty() || OpOpenParen == s2.Peak() {
				s2.Push(e)
			} else {
				top := s2.Peak().(Operator)
				if e.(Operator).PriorTo(top) {
					s2.Push(e)
				} else {
					s1.Push(s2.Pop())
					goto start
				}
			}
		case OpOpenParen:
			s2.Push(e)
		case OpCloseParen:
			for !s2.IsEmpty() {
				top := s2.Pop().(Operator)
				if OpOpenParen == top {
					continue loop
				} else {
					s1.Push(top)
				}
			}
			return fmt.Errorf("err1")
		}
	}

	if s1.IsEmpty() {
		return fmt.Errorf("err2")
	}

	for !s2.IsEmpty() {
		s1.Push(s2.Pop())
	}

	postfixExpression := make([]interface{}, 0, s1.Len())
	for !s1.IsEmpty() {
		postfixExpression = append(postfixExpression, s1.Pop())
	}
	exp.postfixExpression = postfixExpression

	return nil
}

func compare(left interface{}, op Operator, right interface{}) bool {
	//leftT := reflect.ValueOf(left)
	//rightT := reflect.TypeOf(right)
	//
	//if leftT.Kind() {
	//
	//}

	switch op {
	case OpEq:

	case OpNeq:

	}

	return false
}

func (exp *BoolExpression) Evaluate(params map[string]interface{}) bool {
	stackT := NewStack()

	for i := len(exp.postfixExpression) - 1; i >= 0; i-- {
		top := exp.postfixExpression[i]
		if _, ok := top.(Operator); !ok {
			stackT.Push(top)
		} else {
			switch top {
			case OpEq, OpNeq:
				var rightV interface{}
				var leftV interface{}
				rightV = stackT.Pop()
				if right, ok := rightV.(Variable); ok {
					rightV = params[right.name()]
				} else {
					rightV = rightV.(LiteralValue).Value
				}

				leftV = stackT.Pop()
				if left, ok := leftV.(Variable); ok {
					leftV = params[left.name()]
				} else {
					leftV = leftV.(LiteralValue).Value
				}

				if top == OpEq {
					stackT.Push(BoolValue(leftV == rightV))
				} else {
					stackT.Push(BoolValue(leftV != rightV))
				}

			case OpOr, OpAnd:
				if stackT.Len() >= 2 {
					rightV, ok := stackT.Pop().(Booler)
					if !ok {
						return false
					}
					leftV, ok := stackT.Pop().(Booler)
					if !ok {
						return false
					}
					var result bool
					if top == OpOr {
						result = leftV.Bool() || rightV.Bool()
					} else {
						result = leftV.Bool() && rightV.Bool()
					}
					stackT.Push(BoolValue(result))
				} else {
					return false
				}
			case OpNot:
				if stackT.Len() >= 1 {
					exp, ok := stackT.Pop().(Booler)
					if !ok {
						return false
					}
					result := !exp.Bool()
					stackT.Push(BoolValue(result))
				} else {
					return false
				}
			}
		}

	}

	if stackT.Len() != 1 {
		return false
	}

	result := stackT.Pop().(Booler)

	return result.Bool()
}

func NewBoolExpression(expression string) (*BoolExpression, error) {
	if len(expression) == 0 {
		return nil, fmt.Errorf("empty expression")
	}

	exp := &BoolExpression{
		RawExpression: expression,
		scanner:       &Scanner{},
	}
	//fmt.Println(exp.RawExpression)

	if err := exp.init(); err != nil {
		return nil, err
	}

	//fmt.Println(exp.scanner.elements)

	return exp, nil
}

func ToPostfixExpression(exs []interface{}) (*Stack, error) {
	s1 := NewStack()
	s2 := NewStack()

loop:
	for _, e := range exs {
		if _, ok := e.(Operator); ok {
			switch e {
			case OpNot, OpOr, OpAnd, OpEq, OpNeq:
			start:
				if s2.IsEmpty() || OpOpenParen == s2.Peak() {
					s2.Push(e)
				} else {
					top := s2.Peak().(Operator)
					if e.(Operator).PriorTo(top) {
						s2.Push(e)
					} else {
						s1.Push(s2.Pop())
						goto start
					}
				}
			case OpOpenParen:
				s2.Push(e)
			case OpCloseParen:
				for !s2.IsEmpty() {
					top := s2.Pop().(Operator)
					if OpOpenParen == top {
						continue loop
					} else {
						s1.Push(top)
					}
				}
				return nil, fmt.Errorf("err1")
			}
		} else {
			s1.Push(e)
		}
	}

	if s1.IsEmpty() {
		return nil, fmt.Errorf("err2")
	}

	for !s2.IsEmpty() {
		s1.Push(s2.Pop())
	}

	// 倒序
	s3 := NewStack()
	for !s1.IsEmpty() {
		s3.Push(s1.Pop())
	}

	return s3, nil
}

func Evaluate(exs []interface{}, params map[string]interface{}) (bool, error) {
	s, err := ToPostfixExpression(exs)
	if err != nil {
		return false, err
	}

	stackT := NewStack()
	for !s.IsEmpty() {
		top := s.Pop()
		if _, ok := top.(Operator); ok {
			switch top {
			case OpEq:

			case OpNeq:

			case OpOr, OpAnd:
				if stackT.Len() >= 2 {
					rightV, ok := stackT.Pop().(Booler)
					if !ok {
						return false, fmt.Errorf("can not be bool")
					}
					leftV, ok := stackT.Pop().(Booler)
					if !ok {
						return false, fmt.Errorf("can not be bool")
					}
					var result bool
					if top == OpOr {
						result = leftV.Bool() || rightV.Bool()
					} else {
						result = leftV.Bool() && rightV.Bool()
					}
					stackT.Push(BoolValue(result))
				} else {
					return false, fmt.Errorf("lack for params")
				}
			case OpNot:
				if stackT.Len() >= 1 {
					exp, ok := stackT.Pop().(Booler)
					if !ok {
						return false, fmt.Errorf("can not be bool")
					}
					result := !exp.Bool()
					stackT.Push(BoolValue(result))
				} else {
					return false, fmt.Errorf("lack for params")
				}
			}
		} else {
			stackT.Push(top)
		}
	}

	if stackT.Len() != 1 {
		return false, fmt.Errorf("evaluate err")
	}

	result := stackT.Pop().(Booler)

	return result.Bool(), nil

}

func main2() {
	ex := []interface{}{BoolValue(false), OpEq, BoolValue(false), OpAnd, OpOpenParen, BoolValue(false), OpOr, BoolValue(true), OpCloseParen, OpAnd, BoolValue(true), OpAnd, OpNot, BoolValue(false)}
	fmt.Println(ex)

	s, err := ToPostfixExpression(ex)
	if err != nil {
		fmt.Println(err)
	}
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}

	fmt.Println(Evaluate(ex, nil))

	//expression, err := govaluate.NewEvaluableExpression("ls_variate1 == 'exp1'")
	//if err != nil {
	// fmt.Println(err)
	//}
	//parameters := map[string]interface{}{"ls_variate1": "exp1"}
	//result, err := expression.Evaluate(parameters)
	//fmt.Println(result)
}

func main() {
	//a := strconv.Quote("\\n12我3")
	//fmt.Println(a)
	//a, err := strconv.Unquote("\"\\nas我d\"")
	//fmt.Println(err)
	//fmt.Println(a)

	//var a interface{} = 123
	//var b Variable = "123"
	////var c Variable = "true"
	//
	//fmt.Printf("%+v \n", reflect.ValueOf(a).IsValid())
	//fmt.Printf("%+v \n", reflect.ValueOf(a).Kind() == reflect.String)
	//fmt.Printf("%+v \n", reflect.ValueOf(b).)
	//fmt.Printf("%+v \n", reflect.ValueOf(a))
	//fmt.Printf("%+v \n", reflect.ValueOf(b))

	now := time.Now()
	for i := 0; i < 1000; i++ {
		exp, _ := NewBoolExpression(`(${var1}) == true && (${var2} != 2 || ${var3} == "abc")`)
		_ = exp.Evaluate(map[string]interface{}{"var1": true, "var2": 2, "var3": "abc"})
	}

	fmt.Println(time.Since(now))

}

func BenchmarkNewEvaluableExpression() {
	parameters1 := make(map[string]interface{}, 8)
	parameters1["gmv"] = 100
	parameters1["customerId"] = "80"
	parameters1["stayLength"] = 20

	for i := 0; i < 1000; i++ {
		_, _ = govaluate.NewEvaluableExpression("(gmv > 0) && (stayLength > 20) && customerId in ('80','code2','code3')")
	}
}
