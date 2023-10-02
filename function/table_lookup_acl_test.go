package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_acl
// Arguments may be:
// - TABLE, STRING, ACL
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-acl/
func Test_Table_lookup_acl(t *testing.T) {
	acl := vintage.NewAcl("example")
	defaultAcl := vintage.NewAcl("default")
	tests := []struct {
		input   string
		key     string
		expect  *vintage.Acl
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: nil, isError: true},
		{input: "example", key: "foo", expect: acl},
		{input: "example", key: "lorem", expect: defaultAcl},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "ACL",
		vintage.TableItem("foo", acl),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_acl(
			ctx,
			tt.input,
			tt.key,
			defaultAcl,
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
