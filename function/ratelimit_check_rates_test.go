package function

import (
	"testing"
)

// Fastly built-in function testing implementation of ratelimit.check_rates
// Arguments may be:
// - STRING, ID, INTEGER, INTEGER, INTEGER, ID, INTEGER, INTEGER, INTEGER, ID, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-check-rates/
func Test_Ratelimit_check_rates(t *testing.T) {
	t.Skip("Test Builtin function ratelimit.check_rates should be impelemented")
}
