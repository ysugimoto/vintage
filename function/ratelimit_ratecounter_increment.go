package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Ratelimit_ratecounter_increment_Name = "ratelimit.ratecounter_increment"

// Fastly built-in function implementation of ratelimit.ratecounter_increment
// Arguments may be:
// - ID, STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-ratecounter-increment/
func Ratelimit_ratecounter_increment[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	rc string, // IDENT
	entry string,
	delta int64,
) (int64, error) {
	// Need to be implemented
	return 0, errors.FunctionError(
		Ratelimit_ratecounter_increment_Name, "Not Implemented",
	)
}
