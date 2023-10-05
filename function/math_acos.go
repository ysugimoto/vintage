package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_acos_Name = "math.acos"

// Fastly built-in function implementation of math.acos
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acos/
func Math_acos[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1), math.IsInf(val, 1):
		ctx.FastlyError = ErrEDOM
		return math.NaN(), nil
	case val == 1.0:
		return 0, nil
	case val < -1.0 || val > 1.0:
		ctx.FastlyError = ErrEDOM
		return math.NaN(), nil
	default:
		return math.Acos(val), nil
	}
}
