package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_float
// Arguments may be:
// - TABLE, STRING, FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-float/
func Test_Table_lookup_float(t *testing.T) {
	tests := []struct {
		input   string
		key     string
		expect  float64
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: 0, isError: true},
		{input: "example", key: "foo", expect: 10.01},
		{input: "example", key: "lorem", expect: 0.05},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "FLOAT",
		vintage.TableItem("foo", 10.01),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_float(
			ctx,
			tt.input,
			tt.key,
			0.05,
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
