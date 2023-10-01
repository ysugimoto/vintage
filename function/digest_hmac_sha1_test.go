package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_sha1
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha1/
func Test_Digest_hmac_sha1(t *testing.T) {
	ret, err := Digest_hmac_sha1(
		newTestRuntime(),
		"key",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "8513bb355076cce2ae5eb9f399ab5cafdba7c8a2"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
