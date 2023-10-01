package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha1
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1/
func Test_Digest_hash_sha1(t *testing.T) {
	ret, err := Digest_hash_sha1(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "f7c3bc1d808e04732adf679965ccc34ca7ae3441"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
