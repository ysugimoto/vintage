package function

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hmac_md5_Name = "digest.hmac_md5"

// Fastly built-in function implementation of digest.hmac_md5
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-md5/
func Digest_hmac_md5[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, text string,
) (string, error) {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(text))

	return hex.EncodeToString(mac.Sum(nil)), nil
}
