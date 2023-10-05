package function

import (
	"unicode/utf8"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Utf8_is_valid_Name = "utf8.is_valid"

// Fastly built-in function implementation of utf8.is_valid
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-is-valid/
func Utf8_is_valid[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (bool, error) {

	return utf8.ValidString(input), nil
}
