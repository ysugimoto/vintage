package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_cosh_Name = "math.cosh"

// Fastly built-in function implementation of math.cosh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-cosh/
func Math_cosh[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1) || math.IsInf(val, 1):
		return math.Inf(1), nil
	case val == 0:
		return 1.0, nil
	default:
		v := math.Cosh(val)
		if v >= math.Inf(1) {
			return math.Inf(1), nil
		}
		return v, nil
	}
}
