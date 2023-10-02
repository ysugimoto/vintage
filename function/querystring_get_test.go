package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of querystring.get
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-get/
func Test_Querystring_get(t *testing.T) {
	tests := []struct {
		input  string
		second string
		expect string
	}{
		{input: "/?foo=", second: "foo", expect: ""},
		{input: "", second: "", expect: ""},
		{input: "/?foo=&foo=bar", second: "foo", expect: ""},
		{input: "/?a=1", second: "b", expect: ""},
		{input: "/?foo", second: "foo", expect: ""},
		{input: "/?a=1&b=2&c=3&d=4&b=5", second: "b", expect: "2"},
	}

	for i, tt := range tests {
		ret, err := Querystring_get(newTestRuntime(), tt.input, tt.second)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
