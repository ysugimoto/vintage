package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_url_Name = "uuid.url"

// Fastly built-in function implementation of uuid.url
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-url/
func Uuid_url[T core.EdgeRuntime](
	ctx *core.Runtime[T],
) (string, error) {
	// URL namespace, namely constant string
	return uuid.NameSpaceURL.String(), nil
}
