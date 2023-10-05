package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_tanh_Name = "math.tanh"

// Fastly built-in function implementation of math.tanh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-tanh/
func Math_tanh[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1):
		return -1.0, nil
	case math.IsInf(val, 1):
		return 1.0, nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = "ERANGE"
		return val, nil
	default:
		return math.Tanh(val), nil
	}
}
