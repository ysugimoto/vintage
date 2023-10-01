package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_roundeven_Name = "math.roundeven"

// Fastly built-in function implementation of math.roundeven
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-roundeven/
func Math_roundeven[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	if math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1) {
		return val, nil
	}
	if val == 0 {
		return val, nil
	}
	return math.RoundToEven(val), nil
}
