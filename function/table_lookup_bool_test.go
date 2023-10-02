package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_bool
// Arguments may be:
// - TABLE, STRING, BOOL
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-bool/
func Test_Table_lookup_bool(t *testing.T) {
	tests := []struct {
		input   string
		key     string
		expect  bool
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: false, isError: true},
		{input: "example", key: "foo", expect: true},
		{input: "example", key: "lorem", expect: true},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "BOOL",
		vintage.TableItem("foo", true),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_bool(
			ctx,
			tt.input,
			tt.key,
			true,
		)
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
