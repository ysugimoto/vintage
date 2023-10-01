package function

import (
	"testing"
)

// Fastly built-in function testing implementation of cstr_escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/cstr-escape/
func Test_Cstr_escape(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  `"`,
			expect: "\"",
		},
		{
			input:  string([]byte{0x08}),
			expect: "\\b",
		},
		{
			input:  string([]byte{0x09}),
			expect: "\\t",
		},
		{
			input:  string([]byte{0x0A}),
			expect: "\\n",
		},
		{
			input:  string([]byte{0x0B}),
			expect: "\\v",
		},
		{
			input:  string([]byte{0x0D}),
			expect: "\\r",
		},
		{
			input:  string([]byte{0x10}),
			expect: "\\x10",
		},
		{
			input:  "abc",
			expect: "abc",
		},
	}

	for _, tt := range tests {
		ret, err := Cstr_escape(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.expect {
			t.Errorf("Return value unmatch, expect=%s, got=%s", tt.expect, ret)
		}
	}
}
