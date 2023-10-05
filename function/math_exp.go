package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_exp_Name = "math.exp"

// Fastly built-in function implementation of math.exp
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-logexp/math-exp/
func Math_exp[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	switch {
	case math.IsNaN(val):
		return val, nil
	case math.IsInf(val, -1):
		return val, nil
	case math.IsInf(val, 1):
		return val, nil
	}
	return math.Exp(val), nil
}
