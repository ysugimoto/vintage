package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_replaceall_Name = "std.replaceall"

// Fastly built-in function implementation of std.replaceall
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-replaceall/
func Std_replaceall[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, target, replacement string,
) (string, error) {

	return strings.ReplaceAll(input, target, replacement), nil
}
