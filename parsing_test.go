package goexpress

import (
	"github.com/Knetic/govaluate"
	"testing"
	"time"
)

func TestParsing2(test *testing.T) {
	exs, _ := govaluate.NewEvaluableExpression("(gmv > 0) && (stayLength > 20) && customerId in ('80','code2','code3')")

	exs.Eval(govaluate.MapParameters{"gmv": 1, "stayLength": 12, "customerId": "code2"})
}

func TestParsing(test *testing.T) {
	test.Logf("failed to parse original var test:")

	now := time.Now()
	for i := 0; i < 1000; i++ {
		_, _ = govaluate.NewEvaluableExpression("(gmv > 0) && (stayLength > 20) && customerId in ('80','code2','code3')")
	}
	test.Log(time.Since(now))

	now2 := time.Now()
	for i := 0; i < 1000; i++ {
		exp, _ := NewBoolExpression(`(${var1}) == true && (${var2} != 2 || ${var3} == "abc")`)
		_ = exp
		// _ = exp.Evaluate(map[string]interface{}{"var1": true, "var2": 2, "var3": "abc"})
	}

	test.Log(time.Since(now2))
}
