package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type scanner struct {
	chars  []rune
	index  int
	length int
	scan   func(s *scanner) (*Token, error)
}

func newScanner(str string) *scanner {
	cs := []rune(str)
	s := &scanner{
		chars:  cs,
		index:  0,
		length: len(cs),
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

func scanBegin(s *scanner) (*Token, error) {
	builder := strings.Builder{}
	builder.WriteRune('"')
	for s.hasNext() {
		c := s.next()
		builder.WriteRune(c)
		if c == '\\' {
			if s.hasNext() {
				builder.WriteRune(s.next())
			} else {
				return nil, fmt.Errorf("escape error")
			}
		} else if c == '"' {
			str, err := strconv.Unquote(builder.String())
			if err != nil {
				return nil, fmt.Errorf("unquote error: %s", err)
			}
			token := &Token{
				Value: str,
				Type:  TokenString,
			}
			return token, nil
		}
	}

	return nil, fmt.Errorf("not end with quote")
}

func scanString(s *scanner) (*Token, error) {
	builder := strings.Builder{}
	builder.WriteRune('"')
	for s.hasNext() {
		c := s.next()
		builder.WriteRune(c)
		if c == '\\' {
			if s.hasNext() {
				builder.WriteRune(s.next())
			} else {
				return nil, fmt.Errorf("escape error")
			}
		} else if c == '"' {
			str, err := strconv.Unquote(builder.String())
			if err != nil {
				return nil, fmt.Errorf("unquote error: %s", err)
			}
			token := &Token{
				Value: str,
				Type:  TokenString,
			}
			return token, nil
		}
	}

	return nil, fmt.Errorf("not end with quote")
}
