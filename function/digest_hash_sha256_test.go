package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha256
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha256/
func Test_Digest_hash_sha256(t *testing.T) {
	ret, err := Digest_hash_sha256(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
