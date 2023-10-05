package function

import (
	"unicode/utf8"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Utf8_codepoint_count_Name = "utf8.codepoint_count"

// Fastly built-in function implementation of utf8.codepoint_count
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-codepoint-count/
func Utf8_codepoint_count[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (int64, error) {

	return int64(utf8.RuneCountInString(input)), nil
}
