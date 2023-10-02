package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of time.add
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-add/
func Test_Time_add(t *testing.T) {
	now := time.Now()
	tests := []struct {
		duration time.Duration
		time     time.Time
		expect   time.Time
	}{
		{duration: time.Second, expect: now.Add(time.Second)},
	}

	for i, tt := range tests {
		ret, err := Time_add(newTestRuntime(), now, tt.duration)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
