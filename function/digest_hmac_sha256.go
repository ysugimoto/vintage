package function

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hmac_sha256_Name = "digest.hmac_sha256"

// Fastly built-in function implementation of digest.hmac_sha256
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha256/
func Digest_hmac_sha256[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, text string,
) (string, error) {

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(text))
	return hex.EncodeToString(mac.Sum(nil)), nil
}
