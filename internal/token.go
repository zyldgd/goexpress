package internal

type TokenType int

const (
	TokenUnknown   TokenType = iota
	TokenPrefix              // &{
	TokenSuffix              // }
	TokenBoolean             // true  false
	TokenString              // "*****"
	TokenNumber              // 123.45
	TokenPattern             // ( )
	TokenVariable            // var1
	TokenSeparator           // ,
	TokenOperator            // Operator
)

type Token struct {
	Value interface{}
	Type  TokenType
}

// --------------------------------------------------------------
