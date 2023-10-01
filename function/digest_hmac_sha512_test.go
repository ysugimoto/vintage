package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hmac_sha512
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha512/
func Test_Digest_hmac_sha512(t *testing.T) {
	ret, err := Digest_hmac_sha512(
		newTestRuntime(),
		"key",
		"input",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "03ad77c817f32669cccf38dac915d4e55a16833b1ca685979a9df521a05235978090043c3da402dad365ed39ac05af4864a804451861f4b4df7680834c6d4f95"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
