package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_sha1_base64
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha1-base64/
func Test_Digest_hmac_sha1_base64(t *testing.T) {
	ret, err := Digest_hmac_sha1_base64(
		newTestRuntime(),
		"key",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "hRO7NVB2zOKuXrnzmatcr9unyKI="
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
