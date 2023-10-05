package function

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha384_Name = "digest.hash_sha384"

// Fastly built-in function implementation of digest.hash_sha384
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha384/
func Digest_hash_sha384[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := sha512.Sum384([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
