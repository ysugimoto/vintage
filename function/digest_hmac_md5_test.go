package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_md5
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-md5/
func Test_Digest_hmac_md5(t *testing.T) {
	ret, err := Digest_hmac_md5(
		newTestRuntime(),
		"secret",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "bf8d470185ab817f3c92bb5cef1fd7d5"
	if ret != expect {
		t.Errorf("return value unmatch, expect=%s, got=%s", expect, ret)
	}
}
