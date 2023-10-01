package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.acosh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acosh/
func Test_Math_acosh(t *testing.T) {

	tests := []struct {
		input  float64
		expect float64
		err    string
	}{
		{input: 1.0, expect: 0},
		{input: 0.5, err: "EDOM"},
		{input: math.NaN(), err: "EDOM"},
		{input: math.Inf(1), err: "EDOM"},
		{input: math.Inf(-1), err: "EDOM"},
		{input: 1.2, expect: 0.6223625037147786},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ret, err := Math_acosh(ctx, tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
		if tt.err != "" {
			if diff := cmp.Diff(ctx.FastlyError, tt.err); diff != "" {
				t.Errorf("[%d] Error string unmatch, diff: %s", i, diff)
			}
			if !math.IsNaN(ret) {
				t.Errorf("[%d] If error exists, return value must be NaN", i)
			}
		}
	}
}
