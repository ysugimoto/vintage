package function

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_md5_Name = "digest.hash_md5"

// Fastly built-in function implementation of digest.hash_md5
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-md5/
func Digest_hash_md5[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	enc := md5.Sum([]byte(input))
	return hex.EncodeToString(enc[:]), nil
}
