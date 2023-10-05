package function

import (
	"math"
	"math/rand"
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Randombool_Name = "randombool"

// Fastly built-in function implementation of randombool
// Arguments may be:
// - INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool/
func Randombool[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	numerator, denominator int64,
) (bool, error) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rand.Int63n(math.MaxInt64)

	return r/math.MaxInt64 < numerator/denominator, nil
}
