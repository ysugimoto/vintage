package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_dns_Name = "uuid.dns"

// Fastly built-in function implementation of uuid.dns
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-dns/
func Uuid_dns[T core.EdgeRuntime](
	ctx *core.Runtime[T],
) (string, error) {
	// DNS namespace, namely constant string
	return uuid.NameSpaceDNS.String(), nil
}
