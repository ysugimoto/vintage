package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Setcookie_get_value_by_name_Name = "setcookie.get_value_by_name"

// Fastly built-in function implementation of setcookie.get_value_by_name
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/setcookie-get-value-by-name/
func Setcookie_get_value_by_name[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	name string,
) (string, error) {

	switch where {
	case "beresp":
		if ctx.BackendResponseHeader == nil {
			return "", errors.FunctionError(
				Setcookie_get_value_by_name_Name,
				"beresp is not accessible",
			)
		}
		return ctx.BackendResponseHeader.GetSetCookie(name), nil
	case "resp":
		if ctx.ResponseHeader == nil {
			return "", errors.FunctionError(
				Setcookie_get_value_by_name_Name,
				"resp is not accessible",
			)
		}
		return ctx.BackendResponseHeader.GetSetCookie(name), nil
	default:
		return "", errors.FunctionError(
			Setcookie_get_value_by_name_Name, "Invalid ident: %s", where,
		)
	}
}
