package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of substr
// Arguments may be:
// - STRING, INTEGER, INTEGER
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/substr/
func Test_Substr(t *testing.T) {
	tests := []struct {
		input   string
		offset  int64
		length  int64
		expect  string
		isError bool
	}{
		{input: "abcdefg", offset: 3, length: 0, expect: "defg"},
		{input: "abcdefg", offset: 0, length: 2, expect: "ab"},
		{input: "abcdefg", offset: 5, length: 3, expect: "fg"},
		{input: "abc", offset: 4, length: 2, expect: "", isError: true},
		{input: "abc", offset: 3, length: 2, expect: ""},
		{input: "abcdefg", offset: -3, length: 2, expect: "ef"},
		{input: "abcdefg", offset: 1, length: -3, expect: "bcd"},
		{input: "abcdefg", offset: -4, length: -3, expect: "d"},
	}

	for i, tt := range tests {
		var ret string
		var err error

		if tt.length != 0 {
			ret, err = Substr(newTestRuntime(), tt.input, tt.offset, tt.length)
		} else {
			ret, err = Substr(newTestRuntime(), tt.input, tt.offset)
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
