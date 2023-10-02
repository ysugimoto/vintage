package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.regfilter
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-regfilter/
func Test_Querystring_regfilter(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "/path?a=b", expect: "/path?a=b"},
		{input: "/path?a=b&utm_source=foo", expect: "/path?a=b"},
		{input: "/path?utm_source=foo", expect: "/path"},
	}

	for i, tt := range tests {
		ret, err := Querystring_regfilter(
			newTestRuntime(),
			tt.input,
			`utm_.*`,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
