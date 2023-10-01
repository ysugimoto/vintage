package function

import (
	"math"

	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Math_atan2_Name = "math.atan2"

// Fastly built-in function implementation of math.atan2
// Arguments may be:
// - FLOAT, FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-atan2/
func Math_atan2[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	y, x float64,
) (float64, error) {
	switch {
	case math.IsNaN(y) || math.IsNaN(x):
		return math.NaN(), nil
	case math.IsInf(y, -1) || math.IsInf(y, 1):
		switch {
		case math.IsInf(x, -1):
			return 3 * (math.Pi / 4), nil
		case math.IsInf(x, 1):
			return math.Pi / 4, nil
		default:
			return math.Pi, nil
		}
	case math.Abs(y) > 0 && math.IsInf(x, -1):
		return math.Pi, nil
	case math.Abs(y) > 0 && math.IsInf(x, 1):
		return 0, nil
	case y == 0 && x < 0:
		return math.Pi, nil
	case y == 0 && x > 0:
		return 0, nil
	case y < 0 && x == 0:
		return -(math.Pi / 2), nil
	case y > 0 && x == 0:
		return math.Pi / 2, nil
	case y == 0 && lib.IsPositiveZero(x):
		// pole error will not occur.
		return 0, nil
	case y == 0 && x == 0:
		ctx.FastlyError = "EDOM"
		return 0, nil
	default:
		v := math.Atan2(y, x)
		// underflow
		if v < math.Inf(-1) {
			ctx.FastlyError = "EDOM"
			return y / x, nil
		}
		return v, nil
	}
}
