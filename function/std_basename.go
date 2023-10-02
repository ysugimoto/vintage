package function

import (
	"path/filepath"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_basename_Name = "std.basename"

// Fastly built-in function implementation of std.basename
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-basename/
func Std_basename[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (string, error) {
	switch val {
	case ".", "":
		return ".", nil
	case "..":
		return "..", nil
	case "/":
		return "/", nil
	default:
		return filepath.Base(strings.TrimSuffix(val, "/")), nil
	}
}
