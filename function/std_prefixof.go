package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_prefixof_Name = "std.prefixof"

// Fastly built-in function implementation of std.prefixof
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-prefixof/
func Std_prefixof[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val, beginsWith string,
) (bool, error) {
	return strings.HasPrefix(val, beginsWith), nil
}
