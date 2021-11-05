package internal

import "fmt"

type Evaluator interface {
	Evaluate(map[string]interface{}) (bool, error)
}

type postfixExpressionLexer struct {
	originTokens      []*Token
	postfixExpression []*Token
}

func (lexer *postfixExpressionLexer) toPostfixExpression() error {
	var stack1 = NewDefaultStack()
	var stack2 = NewDefaultStack()

	for _, token := range lexer.originTokens {
		switch token.Type {
		case TokenBoolean, TokenNumber, TokenVariable, TokenString: // 值类型
			stack1.Push(token)
		case TokenOperator:
			op, ok := token.Value.(Operator)
			if !ok {
				return fmt.Errorf("type not match")
			}
			switch op {
			case OpOpenParen:
				stack2.Push(token)
			case OpCloseParen:
				for !stack2.IsEmpty() {
					top := stack2.Pop().(Operator2)
					if OpOpenParen == top {
						continue loop
					} else {
						s1.Push(top)
					}
				}
			}

		}
	}

}

func (lexer *postfixExpressionLexer) Evaluate(map[string]interface{}) (bool, error) {
	return false, nil
}
