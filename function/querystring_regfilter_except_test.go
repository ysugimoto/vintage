package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.regfilter_except
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-regfilter-except/
func Test_Querystring_regfilter_except(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "/path?a=b", expect: "/path"},
		{input: "/path?a=b&utm_source=foo", expect: "/path?utm_source=foo"},
	}

	for i, tt := range tests {
		ret, err := Querystring_regfilter_except(
			newTestRuntime(),
			tt.input,
			"utm_*",
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
