package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of xml_escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/xml-escape/
func Test_Xml_escape(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "abc123", expect: "abc123"},
		{input: "romeo&juliet", expect: "romeo&amp;juliet"},
		{input: "0 < 1", expect: "0 &lt; 1"},
		{input: "isn't", expect: "isn&apos;t"},
	}

	for i, tt := range tests {
		ret, err := Xml_escape(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}