package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of subfield
// Arguments may be:
// - STRING, STRING, STRING
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/subfield/
func Test_Subfield(t *testing.T) {
	tests := []struct {
		input  string
		field  string
		sep    string
		expect string
	}{
		{input: "foo=bar,lorem=ipsum", field: "foo", expect: "bar"},
		{input: "foo=bar&lorem=ipsum", field: "foo", sep: "&", expect: "bar"},
		{input: "foo=bar&lorem=ipsum", field: "foo", sep: "%", expect: "bar&lorem=ipsum"},
	}

	for i, tt := range tests {
		var ret string
		var err error
		if tt.sep != "" {
			ret, err = Subfield(newTestRuntime(), tt.input, tt.field, tt.sep)
		} else {
			ret, err = Subfield(newTestRuntime(), tt.input, tt.field)
		}
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
