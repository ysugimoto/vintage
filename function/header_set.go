package function

import (
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Header_set_Name = "header.set"

// Fastly built-in function implementation of header.set
// Arguments may be:
// - ID, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-set/
func Header_set[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	name, value string,
) error {
	// Invalid header and protected header are no effect
	if !lib.IsValidHeader(name) {
		return nil
		// return errors.FunctionError(
		// 	Header_set_Name,
		// 	"Invalid header name provided: %s", name,
		// )
	}
	if err := lib.CheckProtectedHeader(name); err != nil {
		return nil
		// return errors.FunctionError(
		// 	Header_set_Name,
		// 	"Header %s is protected", name,
		// )
	}

	switch where {
	case "req":
		if ctx.RequestHeader != nil {
			ctx.RequestHeader.Set(name, value)
		}
	case "resp":
		if ctx.ResponseHeader != nil {
			ctx.ResponseHeader.Set(name, value)
		}
	case "obj", "beresp":
		if ctx.BackendResponseHeader != nil {
			ctx.BackendResponseHeader.Set(name, value)
		}
	case "bereq":
		if ctx.BackendRequestHeader != nil {
			ctx.BackendRequestHeader.Set(name, value)
		}
	}

	return nil
}
