package function

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha224_Name = "digest.hash_sha224"

// Fastly built-in function implementation of digest.hash_sha224
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha224/
func Digest_hash_sha224[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := sha256.Sum224([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
