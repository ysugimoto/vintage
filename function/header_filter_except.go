package function

import (
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Header_filter_except_Name = "header.filter_except"

// Fastly built-in function implementation of header.filter_except
// Arguments may be:
// - ID, STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-filter-except/
func Header_filter_except[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	where string, // IDENT
	headers ...string,
) error {
	filter := make(map[string]struct{})
	for i := 0; i < len(headers); i++ {
		if !lib.IsValidHeader(headers[i]) {
			return errors.FunctionError(
				Header_filter_except_Name,
				"Invalid header name %s is not permitted", headers[i],
			)
		}
		filter[strings.ToLower(headers[i])] = struct{}{}
	}

	switch where {
	case "req":
		if ctx.RequestHeader != nil {
			header_filter_except_delete(ctx.RequestHeader, filter)
		}
	case "resp":
		if ctx.ResponseHeader != nil {
			header_filter_except_delete(ctx.ResponseHeader, filter)
		}
	case "obj", "beresp":
		if ctx.BackendResponseHeader != nil {
			header_filter_except_delete(ctx.BackendResponseHeader, filter)
		}
	case "bereq":
		if ctx.BackendRequestHeader != nil {
			header_filter_except_delete(ctx.BackendRequestHeader, filter)
		}
	}
	return nil
}

func header_filter_except_delete(h *core.Header, filter map[string]struct{}) {
	for key, val := range h.MH {
		if err := lib.CheckProtectedHeader(key); err != nil {
			continue
		}

		if _, ok := filter[strings.ToLower(key)]; ok {
			continue
		}
		var filtered []string
		for i := range val {
			if strings.Contains(val[i], "=") {
				spl := strings.SplitN(val[i], "=", 2)
				if _, ok := filter[strings.ToLower(key+":"+spl[0])]; ok {
					filtered = append(filtered, val[i])
				}
			}
		}
		if len(filtered) == 0 {
			h.Unset(key)
		} else {
			h.MH[key] = filtered
		}
	}
}
