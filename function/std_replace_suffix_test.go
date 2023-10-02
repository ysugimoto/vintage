package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.replace_suffix
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace-suffix/
func Test_Std_replace_suffix(t *testing.T) {
	tests := []struct {
		input   string
		target  string
		replace string
		expect  string
	}{
		{input: "abcabc", target: "bc", replace: "", expect: "abca"},
		{input: "/foo/bar/", target: "/", replace: "", expect: "/foo/bar"},
	}

	for i, tt := range tests {
		ret, err := Std_replace_suffix(
			newTestRuntime(),
			tt.input,
			tt.target,
			tt.replace,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
