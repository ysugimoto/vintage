package function

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha256_Name = "digest.hash_sha256"

// Fastly built-in function implementation of digest.hash_sha256
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha256/
func Digest_hash_sha256[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := sha256.Sum256([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
