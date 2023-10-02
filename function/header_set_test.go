package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of header.set
// Arguments may be:
// - ID, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-set/
func Test_Header_set(t *testing.T) {

	t.Run("Invalid arguments", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "", expect: ""},
			{name: "Invalid%Header$<>", expect: ""},
		}
		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "req", tt.name, "value")
			if err != nil {
				t.Errorf("[%d] Unexpected error return: %s", i, err)
			}

			v, err := Header_get(ctx, "req", tt.name)
			if err != nil {
				t.Errorf("[%d] Unexpected error return: %s", i, err)
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("set for req", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header", expect: "value"},
			{name: "X-Not-Found", expect: "value"},
			{name: "OBJECT:foo", expect: "value"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "req", tt.name, "value")
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			v, err := Header_get(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("set for bereq", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header", expect: "value"},
			{name: "X-Not-Found", expect: "value"},
			{name: "OBJECT:foo", expect: "value"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendRequestHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "bereq", tt.name, "value")
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			v, err := Header_get(ctx, "bereq", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("set for beresp", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header", expect: "value"},
			{name: "X-Not-Found", expect: "value"},
			{name: "OBJECT:foo", expect: "value"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "beresp", tt.name, "value")
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			v, err := Header_get(ctx, "beresp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("set for obj", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header", expect: "value"},
			{name: "X-Not-Found", expect: "value"},
			{name: "OBJECT:foo", expect: "value"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "obj", tt.name, "value")
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			v, err := Header_get(ctx, "obj", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("set for response", func(t *testing.T) {
		tests := []struct {
			name    string
			expect  string
			isError bool
		}{
			{name: "X-Custom-Header", expect: "value"},
			{name: "X-Not-Found", expect: "value"},
			{name: "OBJECT:foo", expect: "value"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{})

			err := Header_set(ctx, "resp", tt.name, "value")
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			v, err := Header_get(ctx, "resp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_set should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_set should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})
}
