package function

import (
	"encoding/base64"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_base64_Name = "digest.base64"

// Fastly built-in function implementation of digest.base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64/
func Digest_base64[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}
