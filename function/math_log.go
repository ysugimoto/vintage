package function

import (
	"math"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_log_Name = "math.log"

// Fastly built-in function implementation of math.log
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-logexp/math-log/
func Math_log[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val float64,
) (float64, error) {

	if math.IsNaN(val) || math.IsInf(val, -1) || math.IsInf(val, 1) {
		return val, nil
	}
	return math.Log(val), nil
}
