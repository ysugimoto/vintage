package function

import (
	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_backend_Name = "table.lookup_backend"

// Fastly built-in function implementation of table.lookup_backend
// Arguments may be:
// - TABLE, STRING, BACKEND
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-backend/
func Table_lookup_backend[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultBackend *vintage.Backend,
) (*vintage.Backend, error) {

	table, ok := ctx.Tables[id]
	if !ok {
		return defaultBackend, errors.FunctionError(
			Table_lookup_backend_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "BACKEND" {
		return defaultBackend, errors.FunctionError(
			Table_lookup_backend_Name,
			"table %s value type is not BACKEND", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(*vintage.Backend); ok {
			return cast, nil
		}
		return nil, errors.FunctionError(
			Table_lookup_backend_Name,
			"table %s item could not cast to BACKEND type for key %s", id, key,
		)
	}
	return defaultBackend, nil
}
