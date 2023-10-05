package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_asinh_Name = "math.asinh"

// Fastly built-in function implementation of math.asinh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-asinh/
func Math_asinh[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, 1), math.IsInf(val, -1):
		return val, nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = ErrERANGE
		return val, nil
	default:
		return math.Asinh(val), nil
	}
}
