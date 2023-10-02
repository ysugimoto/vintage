package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.replace
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace/
func Test_Std_replace(t *testing.T) {
	tests := []struct {
		input   string
		target  string
		replace string
		expect  string
	}{
		{input: "abcabc", target: "b", replace: "", expect: "acabc"},
		{input: "aa", target: "a", replace: "", expect: "a"},
	}

	for i, tt := range tests {
		ret, err := Std_replace(
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
