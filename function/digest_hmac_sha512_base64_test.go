package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_sha512_base64
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha512-base64/
func Test_Digest_hmac_sha512_base64(t *testing.T) {
	ret, err := Digest_hmac_sha512_base64(
		newTestRuntime(),
		"key",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "A613yBfzJmnMzzjayRXU5VoWgzscpoWXmp31IaBSNZeAkAQ8PaQC2tNl7TmsBa9IZKgERRhh9LTfdoCDTG1PlQ=="
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
