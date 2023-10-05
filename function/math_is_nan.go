package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_is_nan_Name = "math.is_nan"

// Fastly built-in function implementation of math.is_nan
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-nan/
func Math_is_nan[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (bool, error) {

	return math.IsNaN(val), nil
}
