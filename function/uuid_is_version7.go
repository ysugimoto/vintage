package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_is_version7_Name = "uuid.is_version7"

// Fastly built-in function implementation of uuid.is_version7
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-version7/
func Uuid_is_version7[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (bool, error) {

	if id, err := uuid.Parse(input); err != nil {
		return false, nil
	} else if id.Version() != 7 {
		return false, nil
	}
	return true, nil
}
