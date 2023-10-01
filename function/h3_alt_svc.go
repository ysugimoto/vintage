package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const H3_alt_svc_Name = "h3.alt_svc"

// Fastly built-in function implementation of h3.alt_svc
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h3-alt-svc/
func H3_alt_svc[T core.EdgeRuntime](ctx *core.Runtime[T]) error {
	ctx.H3AltSvc = true
	return nil
}
