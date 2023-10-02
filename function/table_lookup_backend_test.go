package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_backend
// Arguments may be:
// - TABLE, STRING, BACKEND
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-backend/
func Test_Table_lookup_backend(t *testing.T) {
	backend := vintage.NewBackend("example")
	defaultBackend := vintage.NewBackend("default")
	tests := []struct {
		input   string
		key     string
		expect  *vintage.Backend
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: nil, isError: true},
		{input: "example", key: "foo", expect: backend},
		{input: "example", key: "lorem", expect: defaultBackend},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "BACKEND",
		vintage.TableItem("foo", backend),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_backend(
			ctx,
			tt.input,
			tt.key,
			defaultBackend,
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
