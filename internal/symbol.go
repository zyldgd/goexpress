package internal

import "unicode"

type Operator int

const (
	OpNone          Operator = iota
	OpOpenParen              // (
	OpCloseParen             // )
	OpNot                    // !
	OpEq                     // ==
	OpNeq                    // !=
	OpGt                     // >
	OpLt                     // <
	OpGte                    // >=
	OpLte                    // <=
	OpAnd                    // &&
	OpOr                     // ||
	OpPlus                   // +
	OpMinus                  // -
	OpMultiply               // *
	OpDivide                 // /
	OpModulus                // %
	OpBitwiseAnd             // &
	OpBitwiseOr              // |
	OpBitwiseXor             // ^
	OpBitwiseLShift          // <<
	OpBitwiseRShift          // >>
	OpBitwiseNot             // ~
	OpAccess                 // .
	OpSeparate               // ,
)

var OperatorMap = map[string]Operator{
	"":   OpNone,
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

func getOperator(str string) Operator {
	return OperatorMap[str]
}

func OpPrecedence(Op Operator) int {
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

func (Op Operator) precedence() int {
	return OpPrecedence(Op)
}

func (Op Operator) precedenceTo(Op2 Operator) bool {
	return OpPrecedence(Op) < OpPrecedence(Op2)
}

func (Op Operator) String() string {
	switch Op {
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

func isVariableChar(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || c == '.'
}

func isNumeric(c rune) bool {
	return unicode.IsDigit(c)
}

func isQuote(c rune) bool {
	return c == '"'
}

func isLetter(c rune) bool {
	return unicode.IsLetter(c)
}

func isSpace(c rune) bool {
	return unicode.IsSpace(c)
}
