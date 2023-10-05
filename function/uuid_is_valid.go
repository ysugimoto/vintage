package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_is_valid_Name = "uuid.is_valid"

// Fastly built-in function implementation of uuid.is_valid
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-valid/
func Uuid_is_valid[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (bool, error) {

	if _, err := uuid.Parse(input); err != nil {
		return false, nil
	}
	return true, nil
}
