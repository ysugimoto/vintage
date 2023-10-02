package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.clean
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-clean/
func Test_Querystring_clean(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "/path?name=value&&=value-only&name-only", expect: "/path?name=value&name-only"},
		{input: "/path?", expect: "/path"},
	}

	for i, tt := range tests {
		ret, err := Querystring_clean(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
