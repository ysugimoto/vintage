package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.contains
// Arguments may be:
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-contains/
func Test_Table_contains(t *testing.T) {
	tests := []struct {
		input   string
		key     string
		expect  bool
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: false, isError: true},
		{input: "example", key: "foo", expect: true},
		{input: "example", key: "lorem", expect: false},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "STRING",
		vintage.TableItem("foo", "bar"),
	)

	for i, tt := range tests {
		ret, err := Table_contains(
			ctx,
			tt.input,
			tt.key,
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
