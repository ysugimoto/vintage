package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_is_version4_Name = "uuid.is_version4"

// Fastly built-in function implementation of uuid.is_version4
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-version4/
func Uuid_is_version4[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (bool, error) {
	if id, err := uuid.Parse(input); err != nil {
		return false, nil
	} else if id.Version() != 4 {
		return false, nil
	}
	return true, nil
}
