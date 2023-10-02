package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of time.hex_to_time
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-hex-to-time/
func Test_Time_hex_to_time(t *testing.T) {
	tests := []struct {
		divisor  int64
		dividend string
		expect   time.Time
	}{
		{divisor: 1, dividend: "43b9a355", expect: time.Date(2006, 1, 2, 22, 4, 5, 0, time.UTC)},
		{divisor: 2, dividend: "43b9a355", expect: time.Date(1988, 1, 2, 11, 2, 2, 0, time.UTC)},
	}

	for i, tt := range tests {
		ret, err := Time_hex_to_time(
			newTestRuntime(),
			tt.divisor,
			tt.dividend,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
