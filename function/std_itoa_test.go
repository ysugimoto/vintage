package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.itoa
// Arguments may be:
// - INTEGER, INTEGER
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa/
func Test_Std_itoa(t *testing.T) {
	tests := []struct {
		input  int64
		base   int64
		expect string
	}{
		{input: -10, expect: "-10"},
		{input: 42, base: 16, expect: "2a"},
	}

	for i, tt := range tests {
		var ret string
		var err error

		if tt.base > 0 {
			ret, err = Std_itoa(newTestRuntime(), tt.input, tt.base)
		} else {
			ret, err = Std_itoa(newTestRuntime(), tt.input)
		}

		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
