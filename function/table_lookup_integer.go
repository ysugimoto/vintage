package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_integer_Name = "table.lookup_integer"

// Fastly built-in function implementation of table.lookup_integer
// Arguments may be:
// - TABLE, STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-integer/
func Table_lookup_integer[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultValue int64,
) (int64, error) {

	table, ok := ctx.Tables[id]
	if !ok {
		return defaultValue, errors.FunctionError(
			Table_lookup_integer_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "INTEGER" {
		return defaultValue, errors.FunctionError(
			Table_lookup_integer_Name,
			"table %s value type is not INTEGER", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(int64); ok {
			return cast, nil
		}
		return defaultValue, errors.FunctionError(
			Table_lookup_integer_Name,
			"table %s item could not cast to INTEGER type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
