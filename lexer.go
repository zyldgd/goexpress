package main

import (
	"fmt"
	"github.com/zyldgd/goexpress/token"
)

type lexer struct {
	source []rune
	index  int
	char   rune // current char
}

func NewScanner(e string) *lexer {
	l := &lexer{
		source: []rune(e),
		index:  0,
	}

	l.walk(0)

	return l
}

func (l *lexer) String() string {
	return fmt.Sprintf("source:{%s} index:%d", string(l.source), l.index)
}

func (l *lexer) walk(step int) bool {
	stop := step + l.index
	valid := false
	if stop < 0 {
		l.index = 0
	} else if stop < len(l.source) {
		l.index = stop
		valid = true
	} else {
		l.index = len(l.source) - 1
	}

	l.char = l.source[l.index]

	return valid
}

func (l *lexer) integer() bool {
	if !token.IsDecimal(l.char) {
		return false
	}
	for l.walk(1) {
		if !token.IsDecimal(l.char) {
			break
		}
	}
	return true
}

func (l *lexer) scanNumber() (token.Token, string) {
	start := l.index
	tok := token.Integer
	l.integer()

	if l.char == '.' {
		tok = token.Float
		if !l.walk(1) || !l.integer() {
			tok = token.Illegal
		}
	}

	lit := l.source[start:l.index]
	return tok, string(lit)
}

func (l *lexer) scanEscape() bool {
	switch l.char {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"':
		l.walk(1)
		return true
	default:
		//msg := "unknown escape sequence"
		//if l.char < 0 {
		//	msg = "escape sequence not terminated"
		//}
		// s.error(offs, msg)
		return false
	}
}

func (l *lexer) scanString() (token.Token, string) {
	start := l.index
	tok := token.String

	for l.walk(1) {
		if l.char == '"' {
			break
		} else if l.char == '\\' {
			l.scanEscape()
		}
	}

	if l.char != '"' || start == l.index {
		tok = token.Illegal
	}

	lit := l.source[start : l.index+1]
	return tok, string(lit)
}

func (l *lexer) scanFor(str string) bool {
	step := len(str)
	if step == 0 {
		return true
	}

	start := l.index
	l.walk(step)

	return str == string(l.source[start:l.index+1])
}

func (l *lexer) scanBool() (token.Token, string) {
	//start := l.index
	tok := token.Illegal
	if l.char == 't' {
		if find := l.scanFor("true"); find {
			return token.True, "true"
		}
	} else if l.char == 'f' {
		if find := l.scanFor("false"); find {
			return token.False, "false"
		}
	}
	return tok, ""
}
