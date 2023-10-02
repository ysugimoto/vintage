package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of time.is_after
// Arguments may be:
// - TIME, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-is-after/
func Test_Time_is_after(t *testing.T) {
	now := time.Now()
	tests := []struct {
		t2     time.Time
		expect bool
	}{
		{t2: now.Add(-time.Second), expect: true},
		{t2: now.Add(time.Second), expect: false},
	}

	for i, tt := range tests {
		ret, err := Time_is_after(
			newTestRuntime(),
			now,
			tt.t2,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
