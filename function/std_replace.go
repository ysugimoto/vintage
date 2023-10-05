package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_replace_Name = "std.replace"

// Fastly built-in function implementation of std.replace
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace/
func Std_replace[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, target, replacement string,
) (string, error) {

	return strings.Replace(input, target, replacement, 1), nil
}
