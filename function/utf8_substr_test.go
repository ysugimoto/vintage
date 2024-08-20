package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of utf8.substr
// Arguments may be:
// - STRING, INTEGER, INTEGER
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-substr/
func Test_Utf8_substr(t *testing.T) {
	tests := []struct {
		input   string
		offset  int64
		length  int64
		expect  string
		isError bool
	}{
		{input: "abあdefg", offset: 3, length: 0, expect: "defg"},
		{input: "abあdefg", offset: 0, length: 2, expect: "ab"},
		{input: "abあdefg", offset: 5, length: 3, expect: "fg"},
		{input: "abあ", offset: 4, length: 2, expect: "", isError: true},
		{input: "abあ", offset: 3, length: 2, expect: ""},
		{input: "abあdefg", offset: -3, length: 2, expect: "ef"},
		{input: "abあdefg", offset: 1, length: -3, expect: "bあd"},
	}

	for i, tt := range tests {
		var ret string
		var err error

		if tt.length != 0 {
			ret, err = Utf8_substr(newTestRuntime(), tt.input, tt.offset, tt.length)
		} else {
			ret, err = Utf8_substr(newTestRuntime(), tt.input, tt.offset)
		}

		if err != nil {
			if !tt.isError {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			continue
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
