package function

import (
	"math/rand"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Randomint_seeded_Name = "randomint_seeded"

// Fastly built-in function implementation of randomint_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomint-seeded/
func Randomint_seeded[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	from, to, seed int64,
) (int64, error) {
	rand.New(rand.NewSource(seed))
	r := rand.Int63n(to - from)

	return r + from, nil
}
