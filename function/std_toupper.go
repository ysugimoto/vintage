package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_toupper_Name = "std.toupper"

// Fastly built-in function implementation of std.toupper
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-toupper/
func Std_toupper[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (string, error) {
	return strings.ToUpper(val), nil
}
