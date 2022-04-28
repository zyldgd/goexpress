package token

import (
	"unicode"
	"unicode/utf8"
)

type Token int

const (
	Illegal         Token = iota
	Integer               // 12345
	Float                 // 123.45
	Char                  // 'a'
	String                // "abc"
	True                  // true
	False                 // false
	OpOpenParen           // (
	OpCloseParen          // )
	OpNot                 // !
	OpEq                  // ==
	OpNeq                 // !=
	OpGt                  // >
	OpLt                  // <
	OpGte                 // >=
	OpLte                 // <=
	OpAnd                 // &&
	OpOr                  // ||
	OpPlus                // +
	OpMinus               // -
	OpMultiply            // *
	OpDivide              // /
	OpModulus             // %
	OpBitwiseAnd          // &
	OpBitwiseOr           // |
	OpBitwiseXor          // ^
	OpBitwiseLShift       // <<
	OpBitwiseRShift       // >>
	OpBitwiseNot          // ~
	OpAccess              // .
	OpSeparate            // ,
)

var OperatorMap = map[string]Token{
	"(":  OpOpenParen,
	")":  OpCloseParen,
	"!":  OpNot,
	"==": OpEq,
	"!=": OpNeq,
	">":  OpGt,
	"<":  OpLt,
	">=": OpGte,
	"<=": OpLte,
	"&&": OpAnd,
	"||": OpOr,
	"+":  OpPlus,
	"-":  OpMinus,
	"*":  OpMultiply,
	"/":  OpDivide,
	"%":  OpModulus,
	"&":  OpBitwiseAnd,
	"|":  OpBitwiseOr,
	"^":  OpBitwiseXor,
	"<<": OpBitwiseLShift,
	">>": OpBitwiseRShift,
	"~":  OpBitwiseNot,
	".":  OpAccess,
	",":  OpSeparate,
}

func GetOperator(str string) Token {
	return OperatorMap[str]
}

func OpPrecedence(Op Token) int {
	switch Op {
	case OpOpenParen, OpCloseParen, OpAccess:
		return 1
	case OpNot, OpBitwiseNot:
		return 2
	case OpMultiply, OpDivide, OpModulus:
		return 3
	case OpPlus, OpMinus:
		return 4
	case OpBitwiseLShift, OpBitwiseRShift:
		return 5
	case OpGt, OpLt, OpGte, OpLte:
		return 6
	case OpEq, OpNeq:
		return 7
	case OpBitwiseAnd:
		return 8
	case OpBitwiseXor:
		return 9
	case OpBitwiseOr:
		return 10
	case OpAnd:
		return 11
	case OpOr:
		return 12
	case OpSeparate:
		return 15
	}
	return 0
}

func (tok Token) precedence() int {
	return OpPrecedence(tok)
}

func (tok Token) precedenceTo(Op2 Token) bool {
	return OpPrecedence(tok) < OpPrecedence(Op2)
}

func (tok Token) String() string {
	switch tok {
	case Illegal:
		return "Illegal"
	case Integer:
		return "INTEGER"
	case Float:
		return "FLOAT"
	case Char:
		return "CHAR"
	case String:
		return "STRING"
	case True:
		return "TRUE"
	case False:
		return "FALSE"
	case OpOpenParen:
		return "("
	case OpCloseParen:
		return ")"
	case OpNot:
		return "!"
	case OpEq:
		return "=="
	case OpNeq:
		return "!="
	case OpGt:
		return ">"
	case OpLt:
		return "<"
	case OpGte:
		return ">="
	case OpLte:
		return "<="
	case OpAnd:
		return "&&"
	case OpOr:
		return "||"
	case OpPlus:
		return "+"
	case OpMinus:
		return "-"
	case OpMultiply:
		return "*"
	case OpDivide:
		return "/"
	case OpModulus:
		return "%"
	case OpBitwiseAnd:
		return "&"
	case OpBitwiseOr:
		return "|"
	case OpBitwiseXor:
		return "^"
	case OpBitwiseLShift:
		return "<<"
	case OpBitwiseRShift:
		return ">>"
	case OpBitwiseNot:
		return "~"
	case OpAccess:
		return "."
	case OpSeparate:
		return ","
	}
	return ""
}

// ------------------------------------------------------------------

func IsVariableChar(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || c == '.'
}

func IsQuote(c rune) bool {
	return c == '"'
}

func IsDecimal(c rune) bool {
	return '0' <= c && c <= '9'
}

func IsDigit(c rune) bool {
	return IsDecimal(c) || c >= utf8.RuneSelf && unicode.IsDigit(c)
}

func IsLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || (c >= utf8.RuneSelf && unicode.IsLetter(c))
}

func IsSpace(c rune) bool {
	return unicode.IsSpace(c)
}

func IsOperatorSymbol(c rune) bool {
	switch c {
	case '!', '&', '|', '=', '%', '^', '~', '+', '-', '*', '/':
		return true
	default:
		return false
	}
}
