package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.toupper
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-toupper/
func Test_Std_toupper(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "VerY", expect: "VERY"},
		{input: "012abc", expect: "012ABC"},
	}

	for i, tt := range tests {
		ret, err := Std_toupper(
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
