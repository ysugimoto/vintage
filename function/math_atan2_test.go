package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.atan2
// Arguments may be:
// - FLOAT, FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-atan2/
func Test_Math_atan2(t *testing.T) {
	tests := []struct {
		y      float64
		x      float64
		expect float64
		err    string
		isNaN  bool
	}{
		{y: math.NaN(), x: 1.0, isNaN: true},
		{y: 1.0, x: math.NaN(), isNaN: true},

		{y: 1.0, x: math.Inf(-1), expect: math.Pi},
		{y: -1.0, x: math.Inf(-1), expect: math.Pi},
		{y: 1.0, x: math.Inf(1), expect: 0},
		{y: -1.0, x: math.Inf(1), expect: 0},

		{y: math.Inf(1), x: math.Inf(-1), expect: 3 * math.Pi / 4},
		{y: math.Inf(1), x: math.Inf(1), expect: math.Pi / 4},
		{y: math.Inf(1), x: 1.0, expect: math.Pi},
		{y: math.Inf(-1), x: math.Inf(-1), expect: 3 * math.Pi / 4},
		{y: math.Inf(-1), x: math.Inf(1), expect: math.Pi / 4},
		{y: math.Inf(-1), x: 1.0, expect: math.Pi},

		{y: 0, x: -1.0, expect: math.Pi},
		{y: 0, x: 1.0, expect: 0},

		{y: -1.0, x: 0, expect: -math.Pi / 2},
		{y: 1.0, x: 0, expect: math.Pi / 2},

		{y: 0, x: 0, expect: 0, err: "EDOM"},
		{y: 1.0, x: 0.5, expect: 1.1071487177940904},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ret, err := Math_atan2(ctx, tt.y, tt.x)
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
		if diff := cmp.Diff(ctx.FastlyError, tt.err); diff != "" {
			t.Errorf("[%d] Error string unmatch, diff: %s", i, diff)
		}
	}
}
