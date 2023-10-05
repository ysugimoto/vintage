package function

import (
	"path/filepath"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_dirname_Name = "std.dirname"

// Fastly built-in function implementation of std.dirname
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-dirname/
func Std_dirname[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val string,
) (string, error) {

	switch val {
	case ".", "", "..":
		return ".", nil
	case "/":
		return "/", nil
	default:
		return filepath.Dir(strings.TrimSuffix(val, "/")), nil
	}
}
