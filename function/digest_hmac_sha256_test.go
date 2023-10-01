package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_sha256
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha256/
func Test_Digest_hmac_sha256(t *testing.T) {
	ret, err := Digest_hmac_sha256(
		newTestRuntime(),
		"key",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "9e089ec13af881a8ac227a736c3e7c490ea3b4afca0c5f83dff6393b683a72e3"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
