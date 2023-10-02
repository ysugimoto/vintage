package function

import (
	"testing"
)

// Fastly built-in function testing implementation of ratelimit.check_rate
// Arguments may be:
// - STRING, ID, INTEGER, INTEGER, INTEGER, ID, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-check-rate/
func Test_Ratelimit_check_rate(t *testing.T) {
	t.Skip("Test Builtin function ratelimit.check_rate should be impelemented")
}
