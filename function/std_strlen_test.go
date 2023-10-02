package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strlen
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strlen/
func Test_Std_strlen(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{input: "Hello world!", expect: 12},
		{input: "Hello 日本語!", expect: 16},
	}

	for i, tt := range tests {
		ret, err := Std_strlen(
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
