package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_is_finite_Name = "math.is_finite"

// Fastly built-in function implementation of math.is_finite
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-finite/
func Math_is_finite[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (bool, error) {
	return !math.IsInf(val, -1) && !math.IsInf(val, 1) && !math.IsNaN(val), nil
}
