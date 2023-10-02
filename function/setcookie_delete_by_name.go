package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Setcookie_delete_by_name_Name = "setcookie.delete_by_name"

// Fastly built-in function implementation of setcookie.delete_by_name
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/setcookie-delete-by-name/
func Setcookie_delete_by_name[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	name string,
) (bool, error) {
	switch where {
	case "beresp":
		if ctx.BackendResponseHeader == nil {
			return false, errors.FunctionError(
				Setcookie_delete_by_name_Name,
				"beresp is not accessible",
			)
		}
		return ctx.BackendResponseHeader.DeleteSetCookie(name), nil
	case "resp":
		if ctx.ResponseHeader == nil {
			return false, errors.FunctionError(
				Setcookie_delete_by_name_Name,
				"resp is not accessible",
			)
		}
		return ctx.BackendResponseHeader.DeleteSetCookie(name), nil
	default:
		return false, errors.FunctionError(
			Setcookie_delete_by_name_Name, "Invalid ident: %s", where,
		)
	}

}
