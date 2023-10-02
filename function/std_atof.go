package function

import (
	"strconv"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_atof_Name = "std.atof"

// Fastly built-in function implementation of std.atof
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-atof/
func Std_atof[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (float64, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_atof_Name, "Failed to parse float value: %s, %w", val, err,
		)
	}
	return f, nil
}
