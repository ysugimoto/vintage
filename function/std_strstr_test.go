package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.strstr
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strstr/
func Test_Std_strstr(t *testing.T) {
	tests := []struct {
		input  string
		needle string
		expect string
	}{
		{input: "/foo?a=b", needle: "?", expect: "?a=b"},
		{input: "/foo?a=b", needle: "o?", expect: "o?a=b"},
	}

	for i, tt := range tests {
		ret, err := Std_strstr(
			newTestRuntime(),
			tt.input,
			tt.needle,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
