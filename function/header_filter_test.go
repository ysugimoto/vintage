package function

import (
	"strings"
	"testing"

	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of header.filter
// Arguments may be:
// - ID, STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-filter/
func Test_Header_filter(t *testing.T) {
	t.Run("Invalid argument", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "", isError: true},
			{name: "Invalid%Header$<>", isError: true},
		}
		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			err := Header_filter(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
				}
			}
		}
	})

	t.Run("filter from req", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "X-Custom-Header", isFiltered: true},
			{name: "X-Not-Found"},
			{name: "Content-Length", isError: true},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
			})

			err := Header_filter(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isFiltered {
				if actual != "" {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if actual == "" {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})

	t.Run("filter from bereq", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "X-Custom-Header", isFiltered: true},
			{name: "X-Not-Found"},
			{name: "Content-Length", isError: true},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendRequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
			})

			err := Header_filter(ctx, "bereq", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			actual := ctx.BackendRequestHeader.Get("X-Custom-Header")
			if tt.isFiltered {
				if actual != "" {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if actual == "" {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})

	t.Run("filter from beresp", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "X-Custom-Header", isFiltered: true},
			{name: "X-Not-Found"},
			{name: "Content-Length", isError: true},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
			})

			err := Header_filter(ctx, "beresp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			actual := ctx.BackendResponseHeader.Get("X-Custom-Header")
			if tt.isFiltered {
				if actual != "" {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if actual == "" {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})

	t.Run("filter from obj", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "X-Custom-Header", isFiltered: true},
			{name: "X-Not-Found"},
			{name: "Content-Length", isError: true},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
			})

			err := Header_filter(ctx, "obj", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			actual := ctx.BackendResponseHeader.Get("X-Custom-Header")
			if tt.isFiltered {
				if actual != "" {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if actual == "" {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})

	t.Run("filter from response", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "X-Custom-Header", isFiltered: true},
			{name: "X-Not-Found"},
			{name: "Content-Length", isError: true},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
			})

			err := Header_filter(ctx, "resp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			actual := ctx.ResponseHeader.Get("X-Custom-Header")
			if tt.isFiltered {
				if actual != "" {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if actual == "" {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})

	t.Run("filter for object-like header", func(t *testing.T) {
		tests := []struct {
			name       string
			isFiltered bool
			isError    bool
		}{
			{name: "Object:foo", isFiltered: true},
			{name: "Object:bar", isFiltered: false},
			{name: "Object:baz", isFiltered: false},
		}
		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			err := Header_filter(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_filter should not return error but non-nil: %s", i, err)
					continue
				}
			}

			var exists bool
			for _, v := range ctx.RequestHeader.MH.Values("Object") {
				spl := strings.SplitN(v, "=", 2)
				if spl[0] == "foo" {
					exists = true
					break
				}
			}

			if tt.isFiltered {
				if exists {
					t.Errorf("[%d] Could not be filtered header", i)
				}
			} else {
				if !exists {
					t.Errorf("[%d] Unexpected header has been filtered", i)
				}
			}
		}
	})
}
