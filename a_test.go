package main

import (
	"strconv"
	"testing"
)

func TestScanNumber(t *testing.T) {
	scan := NewScanner("1231.23a\a")
	tok, lit := scan.scanNumber()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)
}

func TestScanString(t *testing.T) {
	quote := strconv.Quote("  -- \"---A-\t-\n-asdasdadsd-")

	scan := NewScanner(quote[0:1])
	tok, lit := scan.scanString()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)
}

func TestWalk(t *testing.T) {

	scan := NewScanner("0123456789")
	scan.walk(100)
	t.Logf("%+v", scan)
	scan.walk(-1)
	t.Logf("%+v", scan)
}

//func main() {
//	a := 1 - +1
//	_ = a
//	src := "1+-2345-.0"
//	f, err := parser.ParseExprFrom(token.NewFileSet(), "", []byte(src), 0)
//	if err != nil {
//		panic(err)
//	}
//
//	_ = ast.Print(nil, f)
//}
