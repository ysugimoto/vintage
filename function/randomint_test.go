package function

import (
	"testing"
)

// Fastly built-in function testing implementation of randomint
// Arguments may be:
// - INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomint/
func Test_Randomint(t *testing.T) {
	tests := []struct {
		from int64
		to   int64
	}{
		{from: 0, to: 99},
		{from: -1, to: 0},
	}

	for i, tt := range tests {
		// for randomize tests, try enough large loop
		for j := 0; j < 10000; j++ {
			ret, err := Randomint(
				newTestRuntime(),
				tt.from,
				tt.to,
			)
			if err != nil {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			if ret < tt.from || ret > tt.to {
				t.Errorf(
					"[%d] Unexpected return value %d value is not in range from %d to %d",
					i, ret, tt.from, tt.to,
				)
			}
		}
	}
}
