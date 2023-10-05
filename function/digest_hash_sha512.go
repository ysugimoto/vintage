package function

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha512_Name = "digest.hash_sha512"

// Fastly built-in function implementation of digest.hash_sha512
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha512/
func Digest_hash_sha512[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := sha512.Sum512([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
