package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_count_Name = "std.count"

// Fastly built-in function implementation of std.count
// Arguments may be:
// - ID
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-count/
func Std_count[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	headers string, // IDENT
) (int64, error) {

	switch headers {
	case "req.headers":
		return int64(len(ctx.RequestHeader.MH)), nil
	case "bereq.headers":
		return int64(len(ctx.BackendRequestHeader.MH)), nil
	case "beresp.headers", "obj.headers":
		return int64(len(ctx.BackendResponseHeader.MH)), nil
	case "resp.headers":
		return int64(len(ctx.ResponseHeader.MH)), nil
	default:
		return 0, errors.FunctionError(
			Std_count_Name,
			"Unexpected or Unsupported collection ident %s is provided", headers,
		)
	}
}
