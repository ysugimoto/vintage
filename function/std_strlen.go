package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strlen_Name = "std.strlen"

// Fastly built-in function implementation of std.strlen
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strlen/
func Std_strlen[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (int64, error) {
	// Note: Fastly does not consider multibyte, so "日本語" in Japanese treat as 15 byte (3bytes per 1 word)
	// And, also does not consider surrogate-pair characeter probably
	return int64(len(val)), nil
}
