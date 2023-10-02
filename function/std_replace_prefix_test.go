package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.replace_prefix
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace-prefix/
func Test_Std_replace_prefix(t *testing.T) {
	tests := []struct {
		input   string
		target  string
		replace string
		expect  string
	}{
		{input: "abcabc", target: "ab", replace: "", expect: "cabc"},
		{input: "0xABCD1234", target: "0x", replace: "", expect: "ABCD1234"},
	}

	for i, tt := range tests {
		ret, err := Std_replace_prefix(
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
