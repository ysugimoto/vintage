package function

import (
	"encoding/base64"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_base64url_Name = "digest.base64url"

// Fastly built-in function implementation of digest.base64url
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url/
func Digest_base64url[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	return base64.URLEncoding.EncodeToString([]byte(input)), nil
}
