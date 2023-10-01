package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_sqrt_Name = "math.sqrt"

// Fastly built-in function implementation of math.sqrt
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-sqrt/
func Math_sqrt[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, 1):
		return val, nil
	case math.IsInf(val, -1):
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	case val == 0:
		return val, nil
	case val < 0:
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	default:
		return math.Sqrt(val), nil
	}
}
