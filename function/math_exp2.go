package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_exp2_Name = "math.exp2"

// Fastly built-in function implementation of math.exp2
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-logexp/math-exp2/
func Math_exp2[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {
	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1), math.IsInf(val, 1):
		return val, nil
	}
	return math.Exp2(val), nil
}
