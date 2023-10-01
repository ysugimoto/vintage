package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.sqrt
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-sqrt/
func Test_Math_sqrt(t *testing.T) {
	tests := []struct {
		input  float64
		expect float64
		err    string
		isNaN  bool
	}{
		{input: math.NaN(), isNaN: true},
		{input: math.Inf(-1), isNaN: true, err: "EDOM"},
		{input: math.Inf(1), expect: math.Inf(1)},
		{input: 0, expect: 0},
		{input: -1, isNaN: true, err: "EDOM"},
		{input: 0.5, expect: 0.7071067811865476},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ret, err := Math_sqrt(ctx, tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if tt.isNaN {
			if !math.IsNaN(ret) {
				t.Errorf("[%d] Return value should be NaN", i)
			}
		} else {
			if diff := cmp.Diff(ret, tt.expect); diff != "" {
				t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
			}
		}
		if diff := cmp.Diff(ctx.FastlyError, tt.expect); diff != "" {
			t.Errorf("[%d] Error string unmatch, diff: %s", i, diff)
		}
	}
}
