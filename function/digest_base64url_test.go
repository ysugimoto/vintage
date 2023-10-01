package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.base64url
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url/
func Test_Digest_base64url(t *testing.T) {
	ret, err := Digest_base64url(
		newTestRuntime(),
		"Καλώς ορίσατε",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "zprOsc67z47PgiDOv8-Bzq_Pg86xz4TOtQ=="
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
