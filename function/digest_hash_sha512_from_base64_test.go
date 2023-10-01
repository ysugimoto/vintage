package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha512_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha512-from-base64/
func Test_Digest_hash_sha512_from_base64(t *testing.T) {
	ret, err := Digest_hash_sha512_from_base64(
		newTestRuntime(),
		"SGVsbG8sIHdvcmxkIQo=",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "09e1e2a84c92b56c8280f4a1203c7cffd61b162cfe987278d4d6be9afbf38c0e8934cdadf83751f4e99d111352bffefc958e5a4852c8a7a29c95742ce59288a8"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
