package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.tolower
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-tolower/
func Test_Std_tolower(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "VerY", expect: "very"},
		{input: "012abc", expect: "012abc"},
	}

	for i, tt := range tests {
		ret, err := Std_tolower(
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
