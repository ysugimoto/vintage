package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of urlencode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/urlencode/
func Test_Urlencode(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "hello world", expect: "hello%20world"},
	}

	for i, tt := range tests {
		ret, err := Urlencode(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
