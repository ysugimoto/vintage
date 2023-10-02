package function

import (
	"strconv"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_itoa_Name = "std.itoa"

// Fastly built-in function implementation of std.itoa
// Arguments may be:
// - INTEGER, INTEGER
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa/
func Std_itoa[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input int64,
	optional ...int64,
) (string, error) {
	var base int64 = 10
	if len(optional) > 0 {
		base = optional[0]
		if base < 2 || base > 36 {
			ctx.FastlyError = "EINVAL"
			return "", errors.FunctionError(
				Std_itoa_Name, "Invalid base value: %d", base,
			)
		}
	}

	return strconv.FormatInt(input, int(base)), nil
}
