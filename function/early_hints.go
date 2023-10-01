package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Early_hints_Name = "early_hints"

// Fastly built-in function implementation of early_hints
// Arguments may be:
// - STRING, STRING, ...
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/early-hints/
func Early_hints[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	resource string,
	optional ...string,
) error {
	// Edge runtime does not support early_hints yet?
	return nil
}
