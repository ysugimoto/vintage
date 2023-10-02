package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.atof
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-atof/
func Test_Std_atof(t *testing.T) {
	tests := []struct {
		input  string
		expect float64
	}{
		{input: "21.95", expect: 21.95},
		{input: "0", expect: 0},
	}

	for i, tt := range tests {
		ret, err := Std_atof(
			newTestRuntime(),
			tt.input,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
