package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_cos_Name = "math.cos"

// Fastly built-in function implementation of math.cos
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-cos/
func Math_cos[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1) || math.IsInf(val, 1):
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	case val == 0:
		return 1.0, nil
	default:
		return math.Cos(val), nil
	}
}
