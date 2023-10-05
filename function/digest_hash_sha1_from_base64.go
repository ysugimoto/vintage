package function

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_sha1_from_base64_Name = "digest.hash_sha1_from_base64"

// Fastly built-in function implementation of digest.hash_sha1_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1-from-base64/
func Digest_hash_sha1_from_base64[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", errors.FunctionError(
			Digest_hash_sha1_from_base64_Name,
			"Failed to decode base64 string: %s", err,
		)
	}
	enc := sha1.Sum(decoded)
	return hex.EncodeToString(enc[:]), nil
}
