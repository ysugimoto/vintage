package function

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha1_Name = "digest.hash_sha1"

// Fastly built-in function implementation of digest.hash_sha1
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1/
func Digest_hash_sha1[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := sha1.Sum([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
