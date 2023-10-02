package function

import (
	"testing"
)

// Fastly built-in function testing implementation of regsub
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/regsub/
func Test_Regsub(t *testing.T) {
	tests := []struct {
		input       string
		pattern     string
		replacement string
		expect      string
	}{
		{input: "www.example.com", pattern: "www\\.", replacement: "", expect: "example.com"},
		{input: "/foo/bar/", pattern: "/$", replacement: "", expect: "/foo/bar"},
		{input: "aaaa", pattern: "a", replacement: "aa", expect: "aaaaa"},
		{input: "foo;bar;baz", pattern: "([^;]*)(;.*)?$", replacement: "\\1bar", expect: "foobar"},
	}

	for i, tt := range tests {
		ret, err := Regsub(
			newTestRuntime(),
			tt.input,
			tt.pattern,
			tt.replacement,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret != tt.expect {
			t.Errorf("[%d] Return value unmatch, expect=%s, got=%s", i, tt.expect, ret)
		}
	}
}
