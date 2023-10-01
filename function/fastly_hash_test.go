package function

import (
	"testing"
)

// Fastly built-in function testing implementation of fastly.hash
// Arguments may be:
// - STRING, INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/fastly-hash/
func Test_Fastly_hash(t *testing.T) {

	// This is hashing test, continue to several times
	for i := 0; i < 1000; i++ {
		ret, err := Fastly_hash(
			newTestRuntime(),
			"example",
			0,
			10,
			100,
		)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret < 10 || ret > 100 {
			t.Errorf("return value is not in expected range, got=%d", ret)
		}
	}
}
