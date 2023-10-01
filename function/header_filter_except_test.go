package function

import (
	"testing"

	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of header.filter_except
// Arguments may be:
// - ID, STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/headers/header-filter-except/
func Test_Header_filter_except(t *testing.T) {

	t.Run("filter except from req", func(t *testing.T) {
		tests := []struct {
			name       string
			isExepcted bool
			isError    bool
		}{
			{name: "X-Custom-Header", isExepcted: true},
			{name: "X-Not-Found"},
			{name: "Fastly-FF"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.RequestHeader = core.NewHeader(map[string][]string{
				"Fastly-FF":           {"test"},
				"X-Custom-Header":     {"value"},
				"X-Additional-Header": {"value"},
			})

			err := Header_filter_except(ctx, "req", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter_except should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_filter_except should not error but non-nil: %s", i, err)
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isExepcted {
				if actual == "" {
					t.Errorf("[%d] Could not be excepted header", i)
				}
			} else {
				if actual != "" {
					t.Errorf("[%d] Unexpected header has been expected: %s", i, actual)
				}
			}

			if ctx.ResponseHeader.Get("Fastly-FF") == "" {
				t.Errorf("[%d] Protected header must not be removed", i)
			}
		}
	})

	t.Run("filter except from bereq", func(t *testing.T) {
		tests := []struct {
			name       string
			isExepcted bool
			isError    bool
		}{
			{name: "X-Custom-Header", isExepcted: true},
			{name: "X-Not-Found"},
			{name: "Fastly-FF"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendRequestHeader = core.NewHeader(map[string][]string{
				"Fastly-FF":           {"test"},
				"X-Custom-Header":     {"value"},
				"X-Additional-Header": {"value"},
			})

			err := Header_filter_except(ctx, "bereq", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter_except should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_filter_except should not error but non-nil: %s", i, err)
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isExepcted {
				if actual == "" {
					t.Errorf("[%d] Could not be excepted header", i)
				}
			} else {
				if actual != "" {
					t.Errorf("[%d] Unexpected header has been expected: %s", i, actual)
				}
			}

			if ctx.ResponseHeader.Get("Fastly-FF") == "" {
				t.Errorf("[%d] Protected header must not be removed", i)
			}
		}
	})

	t.Run("filter except from obj", func(t *testing.T) {
		tests := []struct {
			name       string
			isExepcted bool
			isError    bool
		}{
			{name: "X-Custom-Header", isExepcted: true},
			{name: "X-Not-Found"},
			{name: "Fastly-FF"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"Fastly-FF":           {"test"},
				"X-Custom-Header":     {"value"},
				"X-Additional-Header": {"value"},
			})

			err := Header_filter_except(ctx, "obj", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter_except should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_filter_except should not error but non-nil: %s", i, err)
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isExepcted {
				if actual == "" {
					t.Errorf("[%d] Could not be excepted header", i)
				}
			} else {
				if actual != "" {
					t.Errorf("[%d] Unexpected header has been expected: %s", i, actual)
				}
			}

			if ctx.ResponseHeader.Get("Fastly-FF") == "" {
				t.Errorf("[%d] Protected header must not be removed", i)
			}
		}
	})

	t.Run("filter except from beresp", func(t *testing.T) {
		tests := []struct {
			name       string
			isExepcted bool
			isError    bool
		}{
			{name: "X-Custom-Header", isExepcted: true},
			{name: "X-Not-Found"},
			{name: "Fastly-FF"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
				"Fastly-FF":           {"test"},
				"X-Custom-Header":     {"value"},
				"X-Additional-Header": {"value"},
			})

			err := Header_filter_except(ctx, "beresp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter_except should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_filter_except should not error but non-nil: %s", i, err)
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isExepcted {
				if actual == "" {
					t.Errorf("[%d] Could not be excepted header", i)
				}
			} else {
				if actual != "" {
					t.Errorf("[%d] Unexpected header has been expected: %s", i, actual)
				}
			}

			if ctx.ResponseHeader.Get("Fastly-FF") == "" {
				t.Errorf("[%d] Protected header must not be removed", i)
			}
		}
	})

	t.Run("filter except from resp", func(t *testing.T) {
		tests := []struct {
			name       string
			isExepcted bool
			isError    bool
		}{
			{name: "X-Custom-Header", isExepcted: true},
			{name: "X-Not-Found"},
			{name: "Fastly-FF"},
		}

		for i, tt := range tests {
			ctx := newTestRuntime()
			ctx.ResponseHeader = core.NewHeader(map[string][]string{
				"Fastly-FF":           {"test"},
				"X-Custom-Header":     {"value"},
				"X-Additional-Header": {"value"},
			})

			err := Header_filter_except(ctx, "resp", tt.name)
			if tt.isError {
				if err == nil {
					t.Errorf("[%d] Header_filter_except should return error but nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("[%d] Header_filter_except should not error but non-nil: %s", i, err)
			}

			actual := ctx.RequestHeader.Get("X-Custom-Header")
			if tt.isExepcted {
				if actual == "" {
					t.Errorf("[%d] Could not be excepted header", i)
				}
			} else {
				if actual != "" {
					t.Errorf("[%d] Unexpected header has been expected: %s", i, actual)
				}
			}

			if ctx.ResponseHeader.Get("Fastly-FF") == "" {
				t.Errorf("[%d] Protected header must not be removed", i)
			}
		}
	})
}
