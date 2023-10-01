package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.hash_sha512
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha512/
func Test_Digest_hash_sha512(t *testing.T) {
	ret, err := Digest_hash_sha512(
		newTestRuntime(),
		"123456789",
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "d9e6762dd1c8eaf6d61b3c6192fc408d4d6d5f1176d0c29169bc24e71c3f274ad27fcd5811b313d681f7e55ec02d73d499c95455b6b5bb503acf574fba8ffe85"
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
