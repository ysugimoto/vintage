package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_roundhalfdown_Name = "math.roundhalfdown"

// Fastly built-in function implementation of math.roundhalfdown
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-roundhalfdown/
func Math_roundhalfdown[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	if math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1) {
		return val, nil
	}
	if val == 0 {
		return val, nil
	}
	t := math.Trunc(val)
	if d := math.Abs(val - t); d <= 0.5 {
		return t, nil
	}
	return t + math.Copysign(1, val), nil
}
