package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.secure_is_equal
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-secure-is-equal/
func Test_Digest_secure_is_equal(t *testing.T) {

	tests := []struct {
		s1     string
		s2     string
		expect bool
	}{
		{
			s1:     "thisiscomparestring",
			s2:     "thisiscomparestring",
			expect: true,
		},
		{
			s1:     "thisiscomparestring",
			s2:     "thisiscomparestrin",
			expect: false,
		},
	}

	for _, tt := range tests {
		ret, err := Digest_secure_is_equal(
			newTestRuntime(),
			tt.s1,
			tt.s2,
		)

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.expect {
			t.Errorf("return value unmach, expect=%t, got=%t", tt.expect, ret)
		}
	}
}
