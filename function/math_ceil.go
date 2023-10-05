package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_ceil_Name = "math.ceil"

// Fastly built-in function implementation of math.ceil
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-ceil/
func Math_ceil[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1) || math.IsInf(val, 1):
		return val, nil
	case val == 0:
		return val, nil
	default:
		return math.Ceil(val), nil
	}
}
