package function

import (
	"encoding/base64"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_base64url_nopad_Name = "digest.base64url_nopad"

// Fastly built-in function implementation of digest.base64url_nopad
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url-nopad/
func Digest_base64url_nopad[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	return base64.RawURLEncoding.EncodeToString([]byte(input)), nil
}
