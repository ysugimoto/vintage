package function

import (
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_rtime_Name = "table.lookup_rtime"

// Fastly built-in function implementation of table.lookup_rtime
// Arguments may be:
// - TABLE, STRING, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-rtime/
func Table_lookup_rtime[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultValue time.Duration,
) (time.Duration, error) {
	table, ok := ctx.Tables[id]
	if !ok {
		return defaultValue, errors.FunctionError(
			Table_lookup_rtime_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "RTIME" {
		return defaultValue, errors.FunctionError(
			Table_lookup_rtime_Name,
			"table %s value type is not RTIME", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(time.Duration); ok {
			return cast, nil
		}
		return defaultValue, errors.FunctionError(
			Table_lookup_rtime_Name,
			"table %s item could not cast to RTIME type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
