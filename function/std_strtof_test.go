package function

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strtof
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strtof/
func Test_Std_strtof(t *testing.T) {
	tests := []struct {
		input  string
		base   int64
		expect float64
	}{
		{input: "1.2", base: 10, expect: 1.200},
		{input: "1.2", base: 0, expect: 1.200},
		{input: "-1.2e-3", base: 10, expect: -0.001},
		{input: "-1.2e-3", base: 0, expect: -0.001},
		{input: "0xA.B", base: 16, expect: 10.688},
		{input: "0xA.B", base: 0, expect: 10.688},
		{input: "0xA.Bp-3", base: 16, expect: 1.336},
		{input: "0xA.Bp-3", base: 0, expect: 1.336},
	}

	for i, tt := range tests {
		ret, err := Std_strtof(
			newTestRuntime(),
			tt.input,
			tt.base,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(math.Round(ret*1000)/1000, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
