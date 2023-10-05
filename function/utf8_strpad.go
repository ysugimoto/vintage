package function

import (
	"math"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Utf8_strpad_Name = "utf8.strpad"

// Fastly built-in function implementation of utf8.strpad
// Arguments may be:
// - STRING, INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-strpad/
func Utf8_strpad[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	width int64,
	pad string,
) (string, error) {

	w := int(math.Abs(float64(width)))
	if len(input) >= w {
		return input, nil
	}

	p := []rune(strings.Repeat(pad, w-len(input)))
	if width < 0 {
		return input + string(p[0:w-len(input)]), nil
	}
	return string(p[0:w-len(input)]) + input, nil
}
