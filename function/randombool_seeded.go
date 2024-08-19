package function

import (
	"math/rand"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Randombool_seeded_Name = "randombool_seeded"

// Fastly built-in function implementation of randombool_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool-seeded/
func Randombool_seeded[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	numerator, denominator, seed int64,
) (bool, error) {

	if denominator <= 0 {
		return false, nil
	}

	r := rand.New(rand.NewSource(seed))
	rv := r.Float64()
	ratio := float64(numerator) / float64(denominator)

	return rv < ratio, nil
}
