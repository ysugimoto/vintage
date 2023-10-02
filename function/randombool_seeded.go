package function

import (
	"math"
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
	rand.New(rand.NewSource(seed))
	r := rand.Int63n(math.MaxInt64)

	return r/math.MaxInt64 < numerator/denominator, nil
}
