package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Header_filter_Name = "header.filter"

// Fastly built-in function implementation of header.filter
// Arguments may be:
// - ID, STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-filter/
func Header_filter[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	headers ...string,
) error {
	for i := 1; i < len(headers); i++ {
		if !lib.IsValidHeader(headers[i]) {
			return errors.FunctionError(
				Header_filter_Name,
				"Invalid header name %s is not permitted", headers[i],
			)
		}
	}

	var err error
	switch where {
	case "req":
		if ctx.RequestHeader != nil {
			err = header_filter_delete(ctx.RequestHeader, headers)
		}
	case "resp":
		if ctx.ResponseHeader != nil {
			err = header_filter_delete(ctx.ResponseHeader, headers)
		}
	case "bereq":
		if ctx.BackendRequestHeader != nil {
			err = header_filter_delete(ctx.BackendRequestHeader, headers)
		}
	case "obj", "beresp":
		if ctx.BackendResponseHeader != nil {
			err = header_filter_delete(ctx.BackendResponseHeader, headers)
		}
	}

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func header_filter_delete(h *core.Header, names []string) error {
	for i := range names {
		if err := lib.CheckProtectedHeader(names[i]); err != nil {
			return errors.WithStack(err)
		}
		h.Unset(names[i])
	}
	return nil
}
