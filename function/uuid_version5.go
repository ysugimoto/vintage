package function

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Uuid_version5_Name = "uuid.version5"

// Fastly built-in function implementation of uuid.version5
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version5/
func Uuid_version5[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	ns, name string,
) (string, error) {

	space, err := uuid.Parse(ns)
	if err != nil {
		return "", errors.FunctionError(
			Uuid_version5_Name,
			"Failed to parse namespace of %s", ns,
		)
	}

	return uuid.NewSHA1(space, []byte(name)).String(), nil
}
