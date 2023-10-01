package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/value"
	"github.com/ysugimoto/vintage/runtime/core"
	// "github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of header.get
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-get/
func Test_Header_get(t *testing.T) {

	t.Run("Invalid arguments", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: ""}, expect: ""},
			{name: &value.String{Value: "Invalid%Header$<>"}, expect: ""},
		}
		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from req", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: "value"},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: "valuefoo"},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from bereq", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: "value"},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: "valuefoo"},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendRequestHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "bereq", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from beresp", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: "value"},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: "valuefoo"},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "beresp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from obj", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: "value"},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: "valuefoo"},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "obj", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from response", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: "value"},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: "valuefoo"},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "resp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})

	t.Run("get from invalid id", func(t *testing.T) {
		tests := []struct {
			name    value.Value
			expect  string
			isError bool
		}{
			{name: &value.String{Value: "X-Custom-Header"}, expect: ""},
			{name: &value.String{Value: "X-Not-Found"}, expect: ""},
			{name: &value.String{Value: "OBJECT:foo"}, expect: ""},
			{name: &value.String{Value: "OBJECT:baz"}, expect: ""},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{
				"X-Custom-Header": {"value"},
				"Object":          {"foo=valuefoo", "bar=valuebar"},
			})

			v, err := Header_get(ctx, "foo", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_get should return error but non-nil", i)
				}
				continue
			} else {
				if err != nil {
					t.Errorf("[%d] Header_get should not return error but non-nil: %s", i, err)
					continue
				}
			}

			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Errorf("[%d] Unexpected value returned, diff=%s", i, diff)
			}
		}
	})
}
