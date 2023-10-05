package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_trunc_Name = "math.trunc"

// Fastly built-in function implementation of math.trunc
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-trunc/
func Math_trunc[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	if math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1) {
		return val, nil
	}
	if val == 0 {
		return val, nil
	}
	return math.Trunc(val), nil
}
