package main

import (
	"fmt"
	"github.com/zyldgd/goexpress/token"
	"strconv"
	"strings"
)

type scanner struct {
	source []rune
	index  int
	char   rune // next one
}

func NewScanner(e string) *scanner {
	if len(e) == 0 {
		return nil
	}

	l := &scanner{
		source: []rune(e),
		index:  0,
	}

	l.char = l.source[0]

	return l
}

func (s *scanner) String() string {
	str := strconv.Quote(string(s.source))
	str = str[1 : len(str)-1]
	space := strings.Repeat(" ", s.index)
	return fmt.Sprintf("\n-------------------------------------------------------\n"+
		"%5d-index : %s↓\n"+
		"%5d-source:[%s]\n"+
		"-------------------------------------------------------", s.index, space, len(str), str)
}

// read next one
func (s *scanner) next() {
	s.index++
	if s.index < len(s.source) {
		s.char = s.source[s.index]
		return
	}
	s.char = -1
}

func (s *scanner) nextChar() rune {
	idx := s.index + 1
	if idx < len(s.source) {
		return s.source[idx]
	}
	return -1
}

func (s *scanner) skip() {
	for token.IsSpace(s.char) {
		s.next()
	}
}

// -------------------------------------------------------------------------------------

func (s *scanner) scan() (token.Token, string) {
	s.skip()
	if s.index >= len(s.source) {
		return token.EOF, ""
	}

	tok, lit := token.Illegal, ""

	switch {
	case token.IsDecimal(s.char):
		tok, lit = s.scanNumber()
	case token.IsLetter(s.char) || '_' == s.char:
		tok, lit = s.scanIdentifier()
	case '"' == s.char:
		tok, lit = s.scanString()
	case '\'' == s.char:
		tok, lit = s.scanChar()
	default:
		switch s.char {
		case '+':
			tok, lit = token.OpAdd, "+"
		case '-':
			tok, lit = token.OpMinus, "-"
		case '*':
			tok, lit = token.OpMultiply, "*"
		case '/':
			tok, lit = token.OpDivide, "/"
		case '%':
			tok, lit = token.OpModulus, "%"
		case '(':
			tok, lit = token.OpLParen, "("
		case ')':
			tok, lit = token.OpRParen, ")"
		case '[':
			tok, lit = token.OpLBracket, "["
		case ']':
			tok, lit = token.OpRBracket, "]"
		case '.':
			tok, lit = token.OpAccess, "."
		case '!':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = token.OpNeq, "!="
			} else {
				tok, lit = token.OpNot, "!"
			}
		case '=':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = token.OpEq, "=="
			} else {
				tok, lit = token.Illegal, "" // Illegal
			}
		case '&':
			if '&' == s.nextChar() {
				s.next()
				tok, lit = token.OpAnd, "&&"
			} else {
				tok, lit = token.OpBitwiseAnd, "&"
			}
		case '|':
			if '|' == s.nextChar() {
				s.next()
				tok, lit = token.OpOr, "||"
			} else {
				tok, lit = token.OpBitwiseOr, "|"
			}
		case '^':
			tok, lit = token.OpBitwiseXor, "^"
		case '~':
			tok, lit = token.OpBitwiseNot, "~"
		case '<':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = token.OpLte, "<="
			} else if '<' == s.nextChar() {
				s.next()
				tok, lit = token.OpBitwiseLShift, "<<"
			} else {
				tok, lit = token.OpLt, "<"
			}
		case '>':
			if '=' == s.nextChar() {
				s.next()
				tok, lit = token.OpGte, ">="
			} else if '>' == s.nextChar() {
				s.next()
				tok, lit = token.OpBitwiseRShift, ">>"
			} else {
				tok, lit = token.OpGt, ">"
			}
		}

		s.next()
	}

	return tok, lit
}

// scanNumber fun look for Integer and Float
func (s *scanner) scanNumber() (token.Token, string) {
	start := s.index
	tok := token.Integer
	for token.IsDecimal(s.char) {
		s.next()
	}

	if s.char == '.' {
		tok = token.Float
		s.next()
		if !token.IsDecimal(s.char) {
			return token.Illegal, ""
		}
		for token.IsDecimal(s.char) {
			s.next()
		}
	}
	return tok, string(s.source[start:s.index])
}

// scanEscape parses an escape-sequence where rune is the accepted escaped quote
func (s *scanner) scanEscape() bool {
	s.next()
	switch s.char {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"', '0':
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

// scanString fun look for string
func (s *scanner) scanString() (token.Token, string) {
	start := s.index // start with "
	tok := token.String

	for {
		s.next()
		if s.char == '"' {
			break
		} else if s.char == '\\' {
			if !s.scanEscape() {
				return token.Illegal, ""
			}
		}
	}

	if s.char != '"' || start == s.index {
		return token.Illegal, ""
	}
	s.next()

	return tok, string(s.source[start:s.index])
}

func (s *scanner) scanChar() (token.Token, string) {
	start := s.index // start with '
	tok := token.Char

	s.next()
	if s.char == '\\' {
		if !s.scanEscape() {
			return token.Illegal, ""
		}
	} else if s.char < 0 {
		return token.Illegal, ""
	}
	s.next()
	if s.char != '\'' {
		return token.Illegal, ""
	}

	s.next()
	return tok, string(s.source[start:s.index])
}

func (s *scanner) scanIdentifier() (token.Token, string) {
	start := s.index

	for token.IsLetter(s.char) || token.IsDecimal(s.char) || '_' == s.char {
		s.next()
	}

	return token.ID, string(s.source[start:s.index])
}
