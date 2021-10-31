package internal

type TokenType int

const (
	TokenUnknown   TokenType = iota
	TokenPrefix              // &{
	TokenSuffix              // }
	TokenBoolean             // true  false
	TokenString              // "*****"
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
