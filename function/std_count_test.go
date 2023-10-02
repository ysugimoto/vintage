package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of std.count
// Arguments may be:
// - ID
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-count/
func Test_Std_count(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{input: "req.headers", expect: 2},
		{input: "bereq.headers", expect: 2},
		{input: "beresp.headers", expect: 2},
		{input: "obj.headers", expect: 2},
		{input: "resp.headers", expect: 2},
	}

	ctx := newTestRuntime()
	ctx.RequestHeader = core.NewHeader(map[string][]string{
		"X-Foo":  {"bar"},
		"X-Hoge": {"huga"},
	})
	ctx.BackendRequestHeader = core.NewHeader(map[string][]string{
		"X-Foo":  {"bar"},
		"X-Hoge": {"huga"},
	})
	ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
		"X-Foo":  {"bar"},
		"X-Hoge": {"huga"},
	})
	ctx.ResponseHeader = core.NewHeader(map[string][]string{
		"X-Foo":  {"bar"},
		"X-Hoge": {"huga"},
	})

	for i, tt := range tests {
		ret, err := Std_count(ctx, tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
