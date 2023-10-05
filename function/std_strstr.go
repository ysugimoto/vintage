package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strstr_Name = "std.strstr"

// Fastly built-in function implementation of std.strstr
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strstr/
func Std_strstr[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	haystack, needle string,
) (string, error) {

	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return "", nil
	}

	return haystack[idx:], nil
}
