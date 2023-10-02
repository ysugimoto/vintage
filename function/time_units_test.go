package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of time.units
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-units/
func Test_Time_units(t *testing.T) {
	tests := []struct {
		unit    string
		time    time.Time
		expect  string
		isError bool
	}{
		{unit: "s", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620"},
		{unit: "ms", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620.000"},
		{unit: "us", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620.000000"},
		{unit: "ns", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620.000000000"},
		{unit: "z", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "", isError: true},
	}

	for i, tt := range tests {
		ret, err := Time_units(
			newTestRuntime(),
			tt.unit,
			tt.time,
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
