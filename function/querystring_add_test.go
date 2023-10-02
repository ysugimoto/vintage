package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.add
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-add/
func Test_Querystring_add(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "foo", expect: "foo?a=b"},
		{input: "foo?bar=baz", expect: "foo?bar=baz&a=b"},
	}

	for i, tt := range tests {
		ret, err := Querystring_add(
			newTestRuntime(),
			tt.input, "a", "b",
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
