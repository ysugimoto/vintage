package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.replaceall
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replaceall/
func Test_Std_replaceall(t *testing.T) {
	tests := []struct {
		input   string
		target  string
		replace string
		expect  string
	}{
		{input: "abcabc", target: "b", replace: "", expect: "acac"},
		{input: "/foo+bar/a+b", target: "+", replace: "%2520", expect: "/foo%2520bar/a%2520b"},
	}

	for i, tt := range tests {
		ret, err := Std_replaceall(
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
