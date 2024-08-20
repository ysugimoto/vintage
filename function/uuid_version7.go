package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_version7_Name = "uuid.version7"

// Fastly built-in function implementation of uuid.version7
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version7/
func Uuid_version7[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	ns, name string,
) (string, error) {

	id, err := uuid.NewV7()
	if err != nil {
		return "", errors.FunctionError(
			Uuid_version7_Name,
			"Failed to create uuid v7",
		)
	}

	return id.String(), nil
}
