package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strtol
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strtol/
func Test_Std_strtol(t *testing.T) {
	tests := []struct {
		input  string
		base   int64
		expect int64
	}{
		{input: "123", base: 0, expect: 123},
		{input: "123", base: 10, expect: 123},
		{input: "0123", base: 0, expect: 83},
		{input: "0123", base: 8, expect: 83},
		{input: "0xABC", base: 0, expect: 2748},
		{input: "0xABC", base: 16, expect: 2748},
		{input: "0xABC", base: 24, expect: 6036},
		{input: "0xABC", base: 36, expect: 1553016},
	}

	for i, tt := range tests {
		ret, err := Std_strtol(
			newTestRuntime(),
			tt.input,
			tt.base,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
