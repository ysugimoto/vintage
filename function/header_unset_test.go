package function

import (
	"testing"

	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of header.unset
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-unset/
func Test_Header_unset(t *testing.T) {
	t.Run("Invalid arguments", func(t *testing.T) {
		tests := []struct {
			name    string
			isError bool
		}{
			{name: ""},
			{name: "Invalid%Header$<>"},
		}
		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{})
			Header_set(ctx, "req", tt.name, "value")
			err := Header_unset(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "req", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, v)
			}
		}
	})

	t.Run("unset for req", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header"},
			{name: "Content-Length", expect: "100"},
			{name: "OBJECT:foo"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Content-Length":  {"100"},
				"Object":          {"foo=valuefoo"},
			})
			err := Header_unset(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "req", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != tt.expect {
				t.Errorf("[%d] Unexpected value returned, expect=%s, got=%s", i, tt.expect, v)
			}
		}
	})

	t.Run("unset for bereq", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header"},
			{name: "Content-Length", expect: "100"},
			{name: "OBJECT:foo"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendRequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Content-Length":  {"100"},
				"Object":          {"foo=valuefoo"},
			})
			err := Header_unset(ctx, "bereq", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "bereq", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != tt.expect {
				t.Errorf("[%d] Unexpected value returned, expect=%s, got=%s", i, tt.expect, v)
			}
		}
	})

	t.Run("unset for beresp", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header"},
			{name: "Content-Length", expect: "100"},
			{name: "OBJECT:foo"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Content-Length":  {"100"},
				"Object":          {"foo=valuefoo"},
			})
			err := Header_unset(ctx, "beresp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "beresp", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != tt.expect {
				t.Errorf("[%d] Unexpected value returned, expect=%s, got=%s", i, tt.expect, v)
			}
		}
	})

	t.Run("unset for obj", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header"},
			{name: "Content-Length", expect: "100"},
			{name: "OBJECT:foo"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Content-Length":  {"100"},
				"Object":          {"foo=valuefoo"},
			})
			err := Header_unset(ctx, "obj", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "obj", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != tt.expect {
				t.Errorf("[%d] Unexpected value returned, expect=%s, got=%s", i, tt.expect, v)
			}
		}
	})

	t.Run("unset for resp", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header"},
			{name: "Content-Length", expect: "100"},
			{name: "OBJECT:foo"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Content-Length":  {"100"},
				"Object":          {"foo=valuefoo"},
			})
			err := Header_unset(ctx, "resp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_unset should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}

			v, err := Header_get(ctx, "resp", tt.name)
			if err != nil {
				t.Errorf("[%d] Header_unset should not return error but non-nil: %s", i, err)
				continue
			}
			if v != tt.expect {
				t.Errorf("[%d] Unexpected value returned, expect=%s, got=%s", i, tt.expect, v)
			}
		}
	})
}
