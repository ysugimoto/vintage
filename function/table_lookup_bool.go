package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_bool_Name = "table.lookup_bool"

// Fastly built-in function implementation of table.lookup_bool
// Arguments may be:
// - TABLE, STRING, BOOL
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-bool/
func Table_lookup_bool[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultValue bool,
) (bool, error) {

	table, ok := ctx.Tables[id]
	if !ok {
		return defaultValue, errors.FunctionError(
			Table_lookup_bool_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "BOOL" {
		return defaultValue, errors.FunctionError(
			Table_lookup_bool_Name,
			"table %s value type is not BOOL", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(bool); ok {
			return cast, nil
		}
		return defaultValue, errors.FunctionError(
			Table_lookup_bool_Name,
			"table %s item could not cast to BOOL type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
