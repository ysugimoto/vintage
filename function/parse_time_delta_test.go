package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of parse_time_delta
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/parse-time-delta/
func Test_Parse_time_delta(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{input: "2d", expect: 172800},
		{input: "2D", expect: 172800},
		{input: "3h", expect: 10800},
		{input: "3H", expect: 10800},
		{input: "1m", expect: 60},
		{input: "1M", expect: 60},
		{input: "10s", expect: 10},
		{input: "10S", expect: 10},
	}

	for i, tt := range tests {
		ret, err := Parse_time_delta(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
