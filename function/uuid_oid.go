package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_oid_Name = "uuid.oid"

// Fastly built-in function implementation of uuid.oid
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-oid/
func Uuid_oid[T core.EdgeRuntime](
	ctx *core.Runtime[T],
) (string, error) {
	// ISO OID namespace, namely constant string
	return uuid.NameSpaceOID.String(), nil
}
