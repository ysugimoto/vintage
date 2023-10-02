package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.remove
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-remove/
func Test_Querystring_remove(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "/?foo=", expect: "/"},
		{input: "/path?a=b", expect: "/path"},
	}

	for i, tt := range tests {
		ret, err := Querystring_remove(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
