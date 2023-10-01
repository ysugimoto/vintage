package function

import (
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Header_get_Name = "header.get"

// Fastly built-in function implementation of header.get
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-get/
func Header_get[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	name string,
) (string, error) {
	if !lib.IsValidHeader(name) {
		return "", nil
	}

	switch where {
	case "req":
		if ctx.RequestHeader != nil {
			return ctx.RequestHeader.Get(name), nil
		}
	case "resp":
		if ctx.ResponseHeader != nil {
			return ctx.ResponseHeader.Get(name), nil
		}
	case "obj", "beresp":
		if ctx.BackendResponseHeader != nil {
			return ctx.BackendResponseHeader.Get(name), nil
		}
	case "bereq":
		if ctx.BackendRequestHeader != nil {
			return ctx.BackendRequestHeader.Get(name), nil
		}
	}

	return "", nil
}
