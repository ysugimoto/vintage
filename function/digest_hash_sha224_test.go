package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha224
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha224/
func Test_Digest_hash_sha224(t *testing.T) {
	ret, err := Digest_hash_sha224(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "9b3e61bf29f17c75572fae2e86e17809a4513d07c8a18152acf34521"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
