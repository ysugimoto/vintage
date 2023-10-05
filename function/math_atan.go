package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_atan_Name = "math.atan"

// Fastly built-in function implementation of math.atan
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-atan/
func Math_atan[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1):
		return -math.Pi / 2, nil
	case math.IsInf(val, 1):
		return math.Pi / 2, nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = ErrERANGE
		return val, nil
	default:
		return math.Atan(val), nil
	}
}
