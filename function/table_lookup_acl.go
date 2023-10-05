package function

import (
	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_acl_Name = "table.lookup_acl"

// Fastly built-in function implementation of table.lookup_acl
// Arguments may be:
// - TABLE, STRING, ACL
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-acl/
func Table_lookup_acl[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	defaultAcl *vintage.Acl,
) (*vintage.Acl, error) {

	table, ok := ctx.Tables[id]
	if !ok {
		return nil, errors.FunctionError(
			Table_lookup_acl_Name,
			"table %s does not exist", id,
		)
	}
	if table.Type != "ACL" {
		return nil, errors.FunctionError(
			Table_lookup_acl_Name,
			"table %s value type is not ACL", id,
		)
	}

	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(*vintage.Acl); ok {
			return cast, nil
		}
		return nil, errors.FunctionError(
			Table_lookup_acl_Name,
			"table %s item could not cast to ACL type for key %s", id, key,
		)
	}
	return defaultAcl, nil
}
