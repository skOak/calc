package compute

import (
	"fmt"
	"math"
	"testing"

	"github.com/skOak/calc/variables"
)

type mockSource struct {
	props []string
	val   float64
}

func mockValueSourceFunc(v string) float64 {
	m := &mockSource{
		props: make([]string, 1000*1000),
		val:   1.1,
	}
	return m.val
}

const DELTA = 0.000001

func TestCompute(t *testing.T) {
	for i := 0; i < 100; i++ {
		for expression, expected := range exps {
			v, err := Evaluate(expression, variables.ValueSourceFunc(mockValueSourceFunc))
			if err != nil || math.Abs(v-expected) > DELTA {
				t.Fatal(fmt.Sprintf("calc err, expected: %v, v: %v", expected, v))
			}
		}
	}
}

var exps = map[string]float64{
	"1+__a+__b+__c": 4.3,
}

func BenchmarkEvaluate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for expression, expected := range exps {
			v, err := Evaluate(expression, variables.ValueSourceFunc(mockValueSourceFunc))
			if err != nil || math.Abs(v-expected) > DELTA {
				panic(fmt.Sprintf("calc err, expected: %v, v: %v", expected, v))
			}
		}
	}
}
