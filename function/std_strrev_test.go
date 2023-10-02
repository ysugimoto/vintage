package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strrev
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrev/
func Test_Std_strrev(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "abc", expect: "cba"},
		{input: "abc日本語", expect: ""},
	}

	for i, tt := range tests {
		ret, err := Std_strrev(
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
