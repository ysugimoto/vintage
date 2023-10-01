package function

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hmac_sha1_Name = "digest.hmac_sha1"

// Fastly built-in function implementation of digest.hmac_sha1
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha1/
func Digest_hmac_sha1[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, text string,
) (string, error) {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(text))

	return hex.EncodeToString(mac.Sum(nil)), nil
}
