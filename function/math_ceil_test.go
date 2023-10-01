package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.ceil
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-ceil/
func Test_Math_ceil(t *testing.T) {
	tests := []struct {
		input  float64
		expect float64
		err    string
		isNaN  bool
	}{
		{input: math.NaN(), isNaN: true},
		{input: math.Inf(-1), expect: math.Inf(-1)},
		{input: math.Inf(1), expect: math.Inf(1)},
		{input: 0.2, expect: 1},
		{input: 0.8, expect: 1},
	}

	for i, tt := range tests {
		ret, err := Math_ceil(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if tt.isNaN {
			if !math.IsNaN(ret) {
				t.Errorf("[%d] Return value should be NaN", i)
			}
			continue
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
