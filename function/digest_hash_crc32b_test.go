package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_crc32b
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-crc32b/
func Test_Digest_hash_crc32b(t *testing.T) {
	ret, err := Digest_hash_crc32b(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "2639f4cb"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
