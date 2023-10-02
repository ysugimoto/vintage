package function

import (
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Ratelimit_check_rate_Name = "ratelimit.check_rate"

// Fastly built-in function implementation of ratelimit.check_rate
// Arguments may be:
// - STRING, ID, INTEGER, INTEGER, INTEGER, ID, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-check-rate/
func Ratelimit_check_rate[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	entry string,
	rc string, // IDENT
	delta, window, limit int64,
	pb string, // IDENT
	ttl time.Time,
) (bool, error) {
	// Need to be implemented
	return false, errors.FunctionError(
		Ratelimit_check_rate_Name, "Not Implemented",
	)
}
