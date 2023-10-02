package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_contains_Name = "table.contains"

// Fastly built-in function implementation of table.contains
// Arguments may be:
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-contains/
func Table_contains[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
) (bool, error) {
	table, ok := ctx.Tables[id]
	if !ok {
		return false, errors.FunctionError(
			Table_contains_Name,
			"table %s does not exist", id,
		)
	}

	_, ok = table.Items[key]
	return ok, nil
}
