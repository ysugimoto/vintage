package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.sort
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-sort/
func Test_Querystring_sort(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "foo?b=1&a=2", expect: "foo?a=2&b=1"},
	}

	for i, tt := range tests {
		ret, err := Querystring_sort(
			newTestRuntime(),
			tt.input,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
