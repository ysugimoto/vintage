package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of strftime
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/strftime/
func Test_Strftime(t *testing.T) {
	now := time.Date(2023, 3, 3, 1, 48, 10, 0, time.UTC)
	tests := []struct {
		input  string
		expect string
	}{
		{input: "%Y-%m-%d %H:%M", expect: "2023-03-03 01:48"},
		{input: "%a, %d %b %Y %T %z", expect: "Fri, 03 Mar 2023 01:48:10 +0000"},
		{input: "%Y-%m-%dT%H:%M:%SZ", expect: "2023-03-03T01:48:10Z"},
	}

	for i, tt := range tests {
		ret, err := Strftime(
			newTestRuntime(),
			tt.input,
			now,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
