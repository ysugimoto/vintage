package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_x500_Name = "uuid.x500"

// Fastly built-in function implementation of uuid.x500
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-x500/
func Uuid_x500[T core.EdgeRuntime](
	ctx *core.Runtime[T],
) (string, error) {
	// URL namespace, namely constant string
	return uuid.NameSpaceX500.String(), nil
}
