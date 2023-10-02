package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_collect_Name = "std.collect"

// Fastly built-in function implementation of std.collect
// Arguments may be:
// - ID
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-collect/
func Std_collect[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	optional ...string,
) error {
	// TODO: std.collection has no effect because Golang's HTTP header is automatically collected
	return nil
}
