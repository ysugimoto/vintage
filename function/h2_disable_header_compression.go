package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const H2_disable_header_compression_Name = "h2.disable_header_compression"

// Fastly built-in function implementation of h2.disable_header_compression
// Arguments may be:
// - STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h2-disable-header-compression/
func H2_disable_header_compression[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	optional ...string,
) error {

	ctx.DisableCompressionHeaders = append(ctx.DisableCompressionHeaders, optional...)
	return nil
}
