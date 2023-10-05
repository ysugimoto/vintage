package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strrep_Name = "std.strrep"

// Fastly built-in function implementation of std.strrep
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrep/
func Std_strrep[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	count int64,
) (string, error) {

	if count < 0 {
		count = 0
	}

	return strings.Repeat(input, int(count)), nil
}
