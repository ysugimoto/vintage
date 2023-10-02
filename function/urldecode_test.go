package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of urldecode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/urldecode/
func Test_Urldecode(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "hello%20world+!", expect: "hello world !"},
		{input: "hello%2520world+!", expect: "hello world !"},
	}

	for i, tt := range tests {
		ret, err := Urldecode(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
