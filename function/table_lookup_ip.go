package function

import (
	"net"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_ip_Name = "table.lookup_ip"

// Fastly built-in function implementation of table.lookup_ip
// Arguments may be:
// - TABLE, STRING, IP
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-ip/
func Table_lookup_ip[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultValue net.IP,
) (net.IP, error) {
	table, ok := ctx.Tables[id]
	if !ok {
		return defaultValue, errors.FunctionError(
			Table_lookup_ip_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "IP" {
		return defaultValue, errors.FunctionError(
			Table_lookup_ip_Name,
			"table %s value type is not IP", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(net.IP); ok {
			return cast, nil
		}
		return defaultValue, errors.FunctionError(
			Table_lookup_ip_Name,
			"table %s item could not cast to IP type for key %s", id, key,
		)
	}
	return defaultValue, nil
}
