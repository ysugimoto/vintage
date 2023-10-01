package function

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Fastly_hash_Name = "fastly.hash"

// Fastly built-in function implementation of fastly.hash
// Arguments may be:
// - STRING, INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/fastly-hash/
func Fastly_hash[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key string,
	seed, from, to int64,
) (int64, error) {

	// Note: fastly.hash internal algorithm is not public.
	// So we implement hashing function as our own way:
	// sha256 hash algorithm
	// concatnation with key string and seed, and get random int between from and to
	enc := sha256.Sum256([]byte(key))
	sb := make([]byte, 64)
	n := binary.PutVarint(sb, seed)
	hash := append(enc[:], sb[:n]...)
	v, err := rand.Int(bytes.NewReader(hash), big.NewInt(from+to))
	if err != nil {
		return 0, errors.FunctionError(
			Fastly_hash_Name,
			"Failed to generate random hash: %w", err,
		)
	}

	return v.Int64() - from, nil
}
