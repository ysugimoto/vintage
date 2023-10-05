package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_floor_Name = "math.floor"

// Fastly built-in function implementation of math.floor
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-floor/
func Math_floor[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1), math.IsInf(val, 1):
		return val, nil
	case val == 0:
		return val, nil
	default:
		return math.Floor(val), nil
	}
}
