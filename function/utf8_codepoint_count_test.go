package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of utf8.codepoint_count
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-codepoint-count/
func Test_Utf8_codepoint_count(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{input: "hello, 世界", expect: 9},
		{input: "hello, world", expect: 12},
	}

	for i, tt := range tests {
		ret, err := Utf8_codepoint_count(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
