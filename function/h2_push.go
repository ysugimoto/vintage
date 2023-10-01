package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const H2_push_Name = "h2.push"

// Fastly built-in function implementation of h2.push
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h2-push/
func H2_push[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	resource string,
	optional ...string,
) error {
	// Fastly document does not explain about second argument "as", so we ignore them for now.
	ctx.PushResources = append(ctx.PushResources, resource)

	return nil
}
