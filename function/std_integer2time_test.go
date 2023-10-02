package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage"
)

// Fastly built-in function testing implementation of std.integer2time
// Arguments may be:
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/std-integer2time/
func Test_Std_integer2time(t *testing.T) {
	tests := []struct {
		input  int64
		expect string
	}{
		{input: 1136239445, expect: "Mon, 02 Jan 2006 22:04:05 GMT"},
		{input: 1677350800, expect: "Sat, 25 Feb 2023 18:46:40 GMT"},
	}

	for i, tt := range tests {
		ret, err := Std_integer2time(
			newTestRuntime(),
			tt.input,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(vintage.ToString(ret), tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
