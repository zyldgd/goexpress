package internal

import (
	"fmt"
	"strconv"
)

type scanner struct {
	chars  []rune
	index  int
	length int
	tokens []*Token
	lexer  scan
}

type scan func(s *scanner) (scan, error)

func newScanner(str string) *scanner {
	cs := []rune(str)
	s := &scanner{
		chars:  cs,
		index:  0,
		length: len(cs),
		tokens: make([]*Token, 0),
		lexer:  scanBegin,
	}
	return s
}

func (s *scanner) hasNext() bool {
	return s.index < s.length
}

func (s *scanner) next() rune {
	c := s.chars[s.index]
	s.index++
	return c
}

func (s *scanner) appendToken(t *Token) {
	s.tokens = append(s.tokens, t)
}

func (s *scanner) rollback(step int) {
	if s.index-step >= 0 {
		s.index -= step
	}
}

func (s *scanner) skip() (end bool) {
	for s.hasNext() {
		if !isSpace(s.next()) {
			break
		}
	}

	return s.hasNext()
}

func (s *scanner) searchToken() (*Token, error) {
	s.skip()

	for s.hasNext() {
		c := s.next()
		if isLetter(c) {

		}
	}

	return nil, nil
}

// -------------------------------------------------------

func scanBegin(s *scanner) (scan, error) {
	if s.hasNext() {
		switch s.next() {
		case '"':
			return scanString, nil
		}
	}

	return nil, fmt.Errorf("unknown lexer")
}

func scanString(s *scanner) (scan, error) {
	start := s.index - 1
	for s.hasNext() {
		c := s.next()
		if c == '\\' {
			if !s.hasNext() {
				return nil, fmt.Errorf("escape error")
			}
			s.next()
		} else if c == '"' {
			str, err := strconv.Unquote(string(s.chars[start:s.index]))
			if err != nil {
				return nil, fmt.Errorf("unquote error: %s", err)
			}
			s.appendToken(&Token{
				Value: str,
				Type:  TokenString,
			})

			return scanBegin, nil
		}
	}

	return nil, fmt.Errorf("not end with quote")
}

func scanVariable(s *scanner) (scan, error) {
	// todo
	start := s.index - 1
	found := false
	for s.hasNext() {
		if !isVariableChar(s.next()) {
			break
		}
		found = true
	}

	if found {
		v := string(s.chars[start:s.index])
		s.appendToken(&Token{
			Value: v,
			Type:  TokenVariable,
		})

		return scanBegin, nil
	}

	return nil, fmt.Errorf("unknown str")
}

func scanNumber(s *scanner) (scan, error) {
	// todo
	s.rollback(1)
	start := s.index
	found := false
	floating := false
	for s.hasNext() {
		c := s.next()
		if isNumeric(c) {

		}

		if !isNumeric(s.next()) {
			break
		}
		found = true
	}

	if found {
		v := string(s.chars[start:s.index])
		s.appendToken(&Token{
			Value: v,
			Type:  TokenNumber,
		})

		return scanBegin, nil
	}

	return nil, fmt.Errorf("unknown str")
}
