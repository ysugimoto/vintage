package function

import (
	"testing"
)

// Fastly built-in function testing implementation of regsuball
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/regsuball/
func Test_Regsuball(t *testing.T) {
	tests := []struct {
		input       string
		pattern     string
		replacement string
		expect      string
	}{
		{input: "//foo///bar//baz", pattern: "/+", replacement: "/", expect: "/foo/bar/baz"},
		{input: "aaaa", pattern: "a", replacement: "aa", expect: "aaaaaaaa"},
		{input: "foo;bar;baz", pattern: "([^;]*)(;.*)?$", replacement: "\\1bar", expect: "foobar"},
	}

	for i, tt := range tests {
		ret, err := Regsuball(
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
