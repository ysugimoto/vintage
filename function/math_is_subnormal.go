package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_is_subnormal_Name = "math.is_subnormal"

// Fastly built-in function implementation of math.is_subnormal
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-subnormal/
func Math_is_subnormal[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (bool, error) {

	if math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1) {
		return false, nil
	}
	return lib.IsSubnormalFloat64(val), nil
}
