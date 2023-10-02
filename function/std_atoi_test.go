package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.atoi
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-atoi/
func Test_Std_atoi(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{input: "21.95", expect: 21},
		{input: "-100", expect: -100},
		{input: "0", expect: 0},
		{input: "", expect: 0},
	}

	for i, tt := range tests {
		ret, err := Std_atoi(
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
