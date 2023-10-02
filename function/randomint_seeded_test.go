package function

import (
	"testing"
)

// Fastly built-in function testing implementation of randomint_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomint-seeded/
func Test_Randomint_seeded(t *testing.T) {
	tests := []struct {
		from int64
		to   int64
		seed int64
	}{
		{from: 0, to: 99, seed: 1000000},
		{from: -1, to: 0, seed: 1000000},
	}

	for i, tt := range tests {
		for j := 0; j < 10000; j++ {
			ret, err := Randomint_seeded(
				newTestRuntime(),
				tt.from,
				tt.to,
				tt.seed,
			)
			if err != nil {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			if ret < tt.from || ret > tt.to {
				t.Errorf(
					"[%d] Unexpected return value, value %d is not in range from %d to %d",
					i, ret, tt.from, tt.to,
				)
			}
		}
	}
}
