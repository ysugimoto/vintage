package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_is_version5_Name = "uuid.is_version5"

// Fastly built-in function implementation of uuid.is_version5
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-version5/
func Uuid_is_version5[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (bool, error) {
	if id, err := uuid.Parse(input); err != nil {
		return false, nil
	} else if id.Version() != 5 {
		return false, nil
	}
	return true, nil
}
