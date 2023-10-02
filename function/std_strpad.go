package function

import (
	"math"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strpad_Name = "std.strpad"

// Fastly built-in function implementation of std.strpad
// Arguments may be:
// - STRING, INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strpad/
func Std_strpad[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	width int64,
	pad string,
) (string, error) {
	w := int(math.Abs(float64(width)))
	if len(input) >= w {
		return input, nil
	}

	rep := strings.Repeat(pad, w-len(input))
	if width < 0 {
		return input + rep[0:w-len(input)], nil
	}
	return rep[0:w-len(input)] + input, nil
}
