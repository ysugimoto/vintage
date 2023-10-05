package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_tan_Name = "math.tan"

// Fastly built-in function implementation of math.tan
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-tan/
func Math_tan[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1) || math.IsInf(val, 1):
		ctx.FastlyError = ErrEDOM
		return math.NaN(), nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = ErrERANGE
		return val, nil
	default:
		v := math.Tan(val)
		if v >= math.Inf(1) {
			ctx.FastlyError = ErrERANGE
			if val > 0 {
				return math.Inf(1), nil
			} else {
				return math.Inf(-1), nil
			}
		}
		return v, nil
	}
}
