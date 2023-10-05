package function

import (
	"encoding/base64"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_base64url_decode_Name = "digest.base64url_decode"

// Fastly built-in function implementation of digest.base64url_decode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url-decode/
func Digest_base64url_decode[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	dec, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		return "", errors.FunctionError(
			Digest_base64url_decode_Name,
			"Failed to decode base64-url string, %s", err,
		)
	}

	return string(dec), nil
}
