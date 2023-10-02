package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_rtime
// Arguments may be:
// - TABLE, STRING, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-rtime/
func Test_Table_lookup_rtime(t *testing.T) {
	value := time.Second
	defaulValue := time.Microsecond
	tests := []struct {
		input   string
		key     string
		expect  time.Duration
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: defaulValue, isError: true},
		{input: "example", key: "foo", expect: value},
		{input: "example", key: "lorem", expect: defaulValue},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "RTIME",
		vintage.TableItem("foo", value),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_rtime(
			ctx,
			tt.input,
			tt.key,
			defaulValue,
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
