package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of utf8.strpad
// Arguments may be:
// - STRING, INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-strpad/
func Test_Utf8_strpad(t *testing.T) {
	tests := []struct {
		input  string
		width  int64
		pad    string
		expect string
	}{
		{input: "abc", width: -10, pad: "-=", expect: "abc-=-=-=-"},
		{input: "abc", width: 10, pad: "-=", expect: "-=-=-=-abc"},
		{input: "abcdefghijklmn", width: 10, pad: "-=", expect: "abcdefghijklmn"},
		{input: "abcdefghij", width: 10, pad: "-=", expect: "abcdefghij"},
		{input: "abc", width: 7, pad: "🌸🌼", expect: "🌸🌼🌸🌼abc"},
	}

	for i, tt := range tests {
		ret, err := Utf8_strpad(
			newTestRuntime(),
			tt.input,
			tt.width,
			tt.pad,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
