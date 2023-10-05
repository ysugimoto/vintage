package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_secure_is_equal_Name = "digest.secure_is_equal"

// Fastly built-in function implementation of digest.secure_is_equal
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-secure-is-equal/
func Digest_secure_is_equal[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	s1, s2 string,
) (bool, error) {

	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		return false, nil
	}

	for i := range r1 {
		if r1[i] != r2[i] {
			return false, nil
		}
	}

	return true, nil
}
