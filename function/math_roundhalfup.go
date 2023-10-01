package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_roundhalfup_Name = "math.roundhalfup"

// Fastly built-in function implementation of math.roundhalfup
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-roundhalfup/
func Math_roundhalfup[T core.EdgeRuntime](
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
	if d := math.Abs(val - t); d >= 0.5 {
		return t + math.Copysign(1, val), nil
	}
	return t, nil
}
