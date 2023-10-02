package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strrep
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrep/
func Test_Std_strrep(t *testing.T) {
	tests := []struct {
		input  string
		count  int64
		expect string
	}{
		{input: "abc", count: 3, expect: "abcabcabc"},
		{input: "abc", count: -1, expect: ""},
	}

	for i, tt := range tests {
		ret, err := Std_strrep(
			newTestRuntime(),
			tt.input,
			tt.count,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
