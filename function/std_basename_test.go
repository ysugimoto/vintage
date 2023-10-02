package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.basename
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-basename/
func Test_Std_basename(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "", expect: "."},
		{input: "/usr/lib", expect: "lib"},
		{input: "/usr/", expect: "usr"},
		{input: "usr", expect: "usr"},
		{input: "/", expect: "/"},
		{input: ".", expect: "."},
		{input: "..", expect: ".."},
	}

	for i, tt := range tests {
		ret, err := Std_basename(
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
