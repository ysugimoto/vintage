package function

import (
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Ratelimit_check_rates_Name = "ratelimit.check_rates"

// Fastly built-in function implementation of ratelimit.check_rates
// Arguments may be:
// - STRING, ID, INTEGER, INTEGER, INTEGER, ID, INTEGER, INTEGER, INTEGER, ID, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-check-rates/
func Ratelimit_check_rates[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	entry string,
	rc1 string, // IDENT
	delta1, window1, limit1 int64,
	rc2 string, // IDENT
	delta2, window2, limit2 int64,
	pb string, // IDENT
	ttl time.Time,
) (bool, error) {
	// Need to be implemented
	return false, errors.FunctionError(
		Ratelimit_check_rates_Name, "not Implemented",
	)
}
