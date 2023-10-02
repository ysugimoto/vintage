package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_suffixof_Name = "std.suffixof"

// Fastly built-in function implementation of std.suffixof
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-suffixof/
func Std_suffixof[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, suffix string,
) (bool, error) {
	return strings.HasSuffix(input, suffix), nil
}
