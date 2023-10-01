package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_crc32
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-crc32/
func Test_Digest_hash_crc32(t *testing.T) {
	ret, err := Digest_hash_crc32(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "181989fc"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
