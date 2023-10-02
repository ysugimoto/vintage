package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_tolower_Name = "std.tolower"

// Fastly built-in function implementation of std.tolower
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-tolower/
func Std_tolower[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (string, error) {
	return strings.ToLower(val), nil
}
