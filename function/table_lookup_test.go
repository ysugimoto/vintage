package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup
// Arguments may be:
// - TABLE, STRING, STRING
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup/
func Test_Table_lookup(t *testing.T) {
	tests := []struct {
		input        string
		key          string
		defaultValue string
		expect       string
		isError      bool
	}{
		{input: "doesnotexist", key: "foo", expect: "", isError: true},
		{input: "doesnotexist", key: "foo", defaultValue: "fallback", expect: "fallback", isError: true},
		{input: "example", key: "foo", expect: "bar"},
		{input: "example", key: "other", defaultValue: "fallback", expect: "fallback"},
		{input: "example", key: "lorem", expect: ""},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "RTIME",
		vintage.TableItem("foo", "bar"),
	)

	for i, tt := range tests {
		var ret string
		var err error

		if tt.defaultValue != "" {
			ret, err = Table_lookup(ctx, tt.input, tt.key, tt.defaultValue)
		} else {
			ret, err = Table_lookup(ctx, tt.input, tt.key)
		}

		if err != nil {
			if !tt.isError {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			continue
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
