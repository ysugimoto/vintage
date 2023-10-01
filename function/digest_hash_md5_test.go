package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_md5
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-md5/
func Test_Digest_hash_md5(t *testing.T) {
	ret, err := Digest_hash_md5(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "25f9e794323b453885f5181f1b624d0b"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
