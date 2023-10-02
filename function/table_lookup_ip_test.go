package function

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of table.lookup_ip
// Arguments may be:
// - TABLE, STRING, IP
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-ip/
func Test_Table_lookup_ip(t *testing.T) {
	value := net.IPv4(10, 0, 0, 0)
	defaulValue := net.IPv4(10, 0, 0, 1)
	tests := []struct {
		input   string
		key     string
		expect  net.IP
		isError bool
	}{
		{input: "doesnotexist", key: "foo", expect: defaulValue, isError: true},
		{input: "example", key: "foo", expect: value},
		{input: "example", key: "lorem", expect: defaulValue},
	}

	ctx := newTestRuntime()
	ctx.Tables["example"] = vintage.NewTable("example", "IP",
		vintage.TableItem("foo", value),
	)

	for i, tt := range tests {
		ret, err := Table_lookup_ip(
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
