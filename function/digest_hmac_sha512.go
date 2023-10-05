package function

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hmac_sha512_Name = "digest.hmac_sha512"

// Fastly built-in function implementation of digest.hmac_sha512
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha512/
func Digest_hmac_sha512[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, text string,
) (string, error) {

	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(text))
	return hex.EncodeToString(mac.Sum(nil)), nil
}
