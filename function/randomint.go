package function

import (
	"math/rand"
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Randomint_Name = "randomint"

// Fastly built-in function implementation of randomint
// Arguments may be:
// - INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomint/
func Randomint[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	from, to int64,
) (int64, error) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rand.Int63n(to - from + 1)

	return r + from, nil
}
