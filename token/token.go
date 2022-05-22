package token

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Token int

const (
	Illegal         Token = iota
	EOF                   // eof
	ID                    // Identifier
	Integer               // 12345
	Float                 // 123.45
	Char                  // 'a'
	String                // "abc"
	OpLParen              // (
	OpRParen              // )
	OpLBracket            // [
	OpRBracket            // ]
	OpNot                 // !
	OpEq                  // ==
	OpNeq                 // !=
	OpGt                  // >
	OpLt                  // <
	OpGte                 // >=
	OpLte                 // <=
	OpAnd                 // &&
	OpOr                  // ||
	OpAdd                 // +
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
	"(":  OpLParen,
	")":  OpRParen,
	"[":  OpLBracket,
	"]":  OpRBracket,
	"!":  OpNot,
	"==": OpEq,
	"!=": OpNeq,
	">":  OpGt,
	"<":  OpLt,
	">=": OpGte,
	"<=": OpLte,
	"&&": OpAnd,
	"||": OpOr,
	"+":  OpAdd,
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

func (tok Token) Precedence() Precedence {
	return OpPrecedence(tok)
}

func (tok Token) PrecedenceWith(Op2 Token) bool {
	return OpPrecedence(tok) < OpPrecedence(Op2)
}

func (tok Token) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", tok.String())
	return []byte(s), nil
}

func (tok Token) String() string {
	switch tok {
	case EOF:
		return "EOF"
	case ID:
		return "Identifier"
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
	case OpLParen:
		return "("
	case OpRParen:
		return ")"
	case OpLBracket:
		return "["
	case OpRBracket:
		return "]"
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
	case OpAdd:
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
