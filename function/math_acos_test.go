package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.acos
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acos/
func Test_Math_acos(t *testing.T) {

	tests := []struct {
		input  float64
		expect float64
		err    string
		isNaN  bool
	}{
		{input: 1.0, expect: 0},
		{input: math.NaN(), err: "EDOM"},
		{input: math.Inf(1), err: "EDOM"},
		{input: math.Inf(-1), err: "EDOM"},
		{input: 1.2, err: "EDOM"},
		{input: 0.5, expect: 1.0471975511965976},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ret, err := Math_acos(ctx, tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
		if tt.err != "" {
			if diff := cmp.Diff(ctx.FastlyError, tt.err); diff != "" {
				t.Errorf("[%d] Error string value unmatch, diff: %s", i, diff)
			}
			if !math.IsNaN(ret) {
				t.Errorf("[%d] If error exists, return value must be NaN", i)
			}
		}
	}
}
