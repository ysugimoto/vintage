package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_float_Name = "table.lookup_float"

// Fastly built-in function implementation of table.lookup_float
// Arguments may be:
// - TABLE, STRING, FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-float/
func Table_lookup_float[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultValue float64,
) (float64, error) {
	table, ok := ctx.Tables[id]
	if !ok {
		return defaultValue, errors.FunctionError(
			Table_lookup_float_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "FLOAT" {
		return defaultValue, errors.FunctionError(
			Table_lookup_float_Name,
			"table %s value type is not FLOAT", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(float64); ok {
			return cast, nil
		}
		return defaultValue, errors.FunctionError(
			Table_lookup_float_Name,
			"table %s item could not cast to FLOAT type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
