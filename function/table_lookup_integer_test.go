package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_integer
// Arguments may be:
// - TABLE, STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-integer/
func Test_Table_lookup_integer(t *testing.T) {
	tests := []struct {
		input   string
		key     string
		expect  int64
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: 0, isError: true},
		{input: "example", key: "foo", expect: 10},
		{input: "example", key: "lorem", expect: 1000},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "INTEGER",
		vintage.TableItem("foo", int64(10)),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_integer(
			ctx,
			tt.input,
			tt.key,
			1000,
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
