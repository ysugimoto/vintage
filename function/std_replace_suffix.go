package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_replace_suffix_Name = "std.replace_suffix"

// Fastly built-in function implementation of std.replace_suffix
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replace-suffix/
func Std_replace_suffix[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, target, replacement string,
) (string, error) {

	if strings.HasSuffix(input, target) {
		input = strings.TrimSuffix(input, target) + replacement
	}

	return input, nil
}
