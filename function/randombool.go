package function

import (
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

	if denominator <= 0 {
		return false, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rv := r.Float64()
	ratio := float64(numerator) / float64(denominator)

	return rv < ratio, nil
}
