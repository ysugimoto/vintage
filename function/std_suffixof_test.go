package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.suffixof
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-suffixof/
func Test_Std_suffixof(t *testing.T) {
	tests := []struct {
		input  string
		suffix string
		expect bool
	}{
		{input: "greenhouse", suffix: "house", expect: true},
		{input: "greenhousa", suffix: "house", expect: false},
	}

	for i, tt := range tests {
		ret, err := Std_suffixof(
			newTestRuntime(),
			tt.input,
			tt.suffix,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
