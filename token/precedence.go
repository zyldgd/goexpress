package token

type Precedence int

func OpPrecedence(Op Token) Precedence {
	switch Op {
	//case OpLParen, OpRParen, OpAccess:
	//	return 1
	case OpNot, OpBitwiseNot:
		return 2
	case OpMultiply, OpDivide, OpModulus:
		return 3
	case OpAdd, OpMinus:
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

func (p Precedence) PrecedenceWith(p2 Precedence) bool {
	return p < p2
}
