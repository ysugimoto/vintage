package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_is_infinite_Name = "math.is_infinite"

// Fastly built-in function implementation of math.is_infinite
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-infinite/
func Math_is_infinite[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (bool, error) {

	return math.IsInf(val, -1) || math.IsInf(val, 1), nil
}
