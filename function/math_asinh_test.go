package function

import (
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.asinh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-asinh/
func Test_Math_asinh(t *testing.T) {

	subnormalValue := math.Float64frombits(0x0000000000000001 | (rand.Uint64() & 0x000fffffffffffff))
	tests := []struct {
		input  float64
		expect float64
		err    string
		isNaN  bool
	}{
		{input: math.NaN(), isNaN: true},
		{input: 0, expect: 0},
		{input: math.Inf(1), expect: math.Inf(1)},
		{input: math.Inf(-1), expect: math.Inf(-1)},
		{input: subnormalValue, expect: subnormalValue, err: "ERANGE"},
		{input: 0.5, expect: 0.48121182505960347},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ret, err := Math_asinh(ctx, tt.input)
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
		if tt.err != "" {
			if diff := cmp.Diff(ctx.FastlyError, tt.err); diff != "" {
				t.Errorf("[%d] Error string unmatch, diff: %s", i, diff)
			}
		}
	}
}
