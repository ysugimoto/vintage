package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha1_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1-from-base64/
func Test_Digest_hash_sha1_from_base64(t *testing.T) {
	ret, err := Digest_hash_sha1_from_base64(
		newTestRuntime(),
		"SGVsbG8sIHdvcmxkIQo=",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "09fac8dbfd27bd9b4d23a00eb648aa751789536d"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
