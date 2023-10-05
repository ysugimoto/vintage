package function

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hmac_md5_base64_Name = "digest.hmac_md5_base64"

// Fastly built-in function implementation of digest.hmac_md5_base64
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-md5-base64/
func Digest_hmac_md5_base64[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, text string,
) (string, error) {

	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}
