package function

import (
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of math.is_normal
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-normal/
func Test_Math_is_normal(t *testing.T) {
	subnormalValue := math.Float64frombits(0x0000000000000001 | (rand.Uint64() & 0x000fffffffffffff))
	tests := []struct {
		input  float64
		expect bool
	}{
		{input: math.NaN(), expect: true},
		{input: math.Inf(-1), expect: true},
		{input: math.Inf(1), expect: true},
		{input: 1.2, expect: true},
		{input: subnormalValue, expect: false},
	}

	for i, tt := range tests {
		ret, err := Math_is_normal(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
