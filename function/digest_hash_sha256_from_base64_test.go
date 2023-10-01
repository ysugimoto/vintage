package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha256_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha256-from-base64/
func Test_Digest_hash_sha256_from_base64(t *testing.T) {
	ret, err := Digest_hash_sha256_from_base64(
		newTestRuntime(),
		"SGVsbG8sIHdvcmxkIQo=",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "d9014c4624844aa5bac314773d6b689ad467fa4e1d1a50a1b8a99d5a95f72ff5"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
