package function

import (
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Header_unset_Name = "header.unset"

// Fastly built-in function implementation of header.unset
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-unset/
func Header_unset[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	name string,
) error {
	if !lib.IsValidHeader(name) {
		return nil
	}
	if err := lib.CheckProtectedHeader(name); err != nil {
		return nil
	}

	switch where {
	case "req":
		if ctx.RequestHeader != nil {
			ctx.RequestHeader.Unset(name)
		}
	case "resp":
		if ctx.ResponseHeader != nil {
			ctx.ResponseHeader.Unset(name)
		}
	case "obj", "beresp":
		if ctx.BackendResponseHeader != nil {
			ctx.BackendResponseHeader.Unset(name)
		}
	case "bereq":
		if ctx.BackendRequestHeader != nil {
			ctx.BackendRequestHeader.Unset(name)
		}
	}

	return nil
}
