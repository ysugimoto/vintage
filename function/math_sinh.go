package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_sinh_Name = "math.sinh"

// Fastly built-in function implementation of math.sinh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-sinh/
func Math_sinh[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1):
		return val, nil
	case val == 0:
		return val, nil
	case lib.IsSubnormalFloat64(val):
		ctx.FastlyError = "ERANGE"
		return val, nil
	}
	v := math.Sinh(val)
	if v >= math.Inf(1) {
		if val > 0 {
			return math.Inf(1), nil
		} else {
			return math.Inf(-1), nil
		}
	}
	return v, nil
}
