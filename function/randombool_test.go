package function

import (
	"testing"
)

// Fastly built-in function testing implementation of randombool
// Arguments may be:
// - INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool/
func Test_Randombool(t *testing.T) {
	tests := []struct {
		n int64
		d int64
	}{
		{n: 1, d: 10},
		{n: 3, d: 4},
		{n: 5, d: 10},
	}

	for i, tt := range tests {
		_, err := Randombool(
			newTestRuntime(),
			tt.n,
			tt.d,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
	}
}
