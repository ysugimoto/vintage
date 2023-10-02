package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_version4_Name = "uuid.version4"

// Fastly built-in function implementation of uuid.version4
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version4/
func Uuid_version4[T core.EdgeRuntime](
	ctx *core.Runtime[T],
) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.FunctionError(
			Uuid_version4_Name, "Failed to create random",
		)
	}

	return id.String(), nil
}
