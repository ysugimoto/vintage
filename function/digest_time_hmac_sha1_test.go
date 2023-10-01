package function

import (
	"crypto/sha1"
	"encoding/base64"
	"testing"
	"time"
)

// Fastly built-in function testing implementation of digest.time_hmac_sha1
// Arguments may be:
// - STRING, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-time-hmac-sha1/
func Test_Digest_time_hmac_sha1(t *testing.T) {
	secret := base64.StdEncoding.EncodeToString([]byte("12345678901234567890"))
	ret, err := digest_time_hmac_sha1(
		time.Date(1970, 1, 1, 0, 0, 59, 0, time.UTC),
		secret,
		30,
		0,
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	enc := sha1.Sum([]byte("287082"))
	expect := base64.StdEncoding.EncodeToString(enc[:])
	if ret != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, ret)
	}
}
