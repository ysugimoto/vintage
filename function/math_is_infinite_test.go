package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.is_infinite
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-infinite/
func Test_Math_is_infinite(t *testing.T) {
	tests := []struct {
		input  float64
		expect bool
	}{
		{input: math.NaN(), expect: false},
		{input: math.Inf(-1), expect: true},
		{input: math.Inf(1), expect: true},
		{input: 1.2, expect: false},
	}

	for i, tt := range tests {
		ret, err := Math_is_infinite(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
