package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_acosh_Name = "math.acosh"

// Fastly built-in function implementation of math.acosh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acosh/
func Math_acosh[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1), math.IsInf(val, 1):
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	case val == 1.0:
		return 0, nil
	case val < 1.0:
		ctx.FastlyError = "EDOM"
		return math.NaN(), nil
	default:
		return math.Acosh(val), nil
	}
}
