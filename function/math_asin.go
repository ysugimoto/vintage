package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_asin_Name = "math.asin"

// Fastly built-in function implementation of math.asin
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-asin/
func Math_asin[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1), math.IsInf(val, 1):
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = "ERANGE"
		return val, nil
	case val < -1.0 || val > 1.0:
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	default:
		return math.Asin(val), nil
	}
}
