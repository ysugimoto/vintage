package function

import (
	"testing"
)

// Fastly built-in function testing implementation of randombool_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool-seeded/
func Test_Randombool_seeded(t *testing.T) {
	tests := []struct {
		n int64
		d int64
		s int64
	}{
		{n: 1, d: 10, s: 1000000},
		{n: 3, d: 4, s: 1111111},
		{n: 5, d: 10, s: 2222222},
	}

	for i, tt := range tests {
		_, err := Randombool_seeded(
			newTestRuntime(),
			tt.n,
			tt.d,
			tt.s,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
	}
}
