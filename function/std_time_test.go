package function

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.time
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/std-time/
func Test_Std_time(t *testing.T) {
	strToTime := func(str, format string) time.Time {
		t, _ := time.Parse(format, str)
		return t
	}

	now := time.Now()

	tests := []struct {
		input  string
		expect time.Time
	}{
		{
			input:  "Mon, 02 Jan 2006 22:04:05 GMT",
			expect: strToTime("Mon, 02 Jan 2006 22:04:05 GMT", time.RFC1123),
		},
		{
			input:  "Monday, 02-Jan-06 22:04:05 GMT",
			expect: strToTime("Monday, 02-Jan-06 22:04:05 GMT", time.RFC850),
		},
		{
			input:  "Mon Jan 2 22:04:05 2006",
			expect: strToTime("Mon Jan 2 22:04:05 2006", time.ANSIC),
		},
		{
			input:  "2006-01-02 22:04:05",
			expect: strToTime("2006-01-02 22:04:05", "2006-01-02 15:04:05"),
		},
		{
			input:  "136239445.00",
			expect: time.Unix(136239445, 0),
		},
		{
			input:  "136239445",
			expect: time.Unix(136239445, 0),
		},
		{
			input:  "foobarbaz",
			expect: now,
		},
	}

	for i, tt := range tests {
		ret, err := Std_time(
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
