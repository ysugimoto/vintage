package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.prefixof
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-prefixof/
func Test_Std_prefixof(t *testing.T) {
	tests := []struct {
		input  string
		prefix string
		expect bool
	}{
		{input: "greenhouse", prefix: "green", expect: true},
		{input: "bluehouse", prefix: "green", expect: false},
	}

	for i, tt := range tests {
		ret, err := Std_prefixof(
			newTestRuntime(),
			tt.input,
			tt.prefix,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
