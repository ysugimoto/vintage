package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_replace_prefix_Name = "std.replace_prefix"

// Fastly built-in function implementation of std.replace_prefix
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace-prefix/
func Std_replace_prefix[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, target, replacement string,
) (string, error) {
	if strings.HasPrefix(input, target) {
		input = replacement + strings.TrimPrefix(input, target)
	}

	return input, nil
}
