package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_sin_Name = "math.sin"

// Fastly built-in function implementation of math.sin
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-sin/
func Math_sin[T core.EdgeRuntime](
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
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = "ERANGE"
		return val, nil
	default:
		return math.Sin(val), nil
	}
}
