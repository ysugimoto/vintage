package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_Name = "table.lookup"

// Fastly built-in function implementation of table.lookup
// Arguments may be:
// - TABLE, STRING, STRING
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup/
func Table_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	optional ...string,
) (string, error) {
	var defaultValue string
	if len(optional) > 0 {
		defaultValue = optional[0]
	}

	table, ok := ctx.Tables[id]
	if !ok {
		return "", errors.FunctionError(
			Table_lookup_Name,
			"table %s does not exist", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(string); ok {
			return cast, nil
		}
		return "", errors.FunctionError(
			Table_lookup_Name,
			"table %s item could not cast to STRING type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
