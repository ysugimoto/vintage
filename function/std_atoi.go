package function

import (
	"strconv"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_atoi_Name = "std.atoi"

// Fastly built-in function implementation of std.atoi
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-atoi/
func Std_atoi[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (int64, error) {
	// If input string is empty, return 0 immediately without raising error
	if val == "" {
		return 0, nil
	}
	// In std.atoi spec, float string is rounded into int, not a raise error
	if idx := strings.Index(val, "."); idx != -1 {
		val = val[0:idx]
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_atoi_Name,
			"Failed to parse int value: %s, %s", val, err,
		)
	}
	return i, nil
}
