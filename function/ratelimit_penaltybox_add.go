package function

import (
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Ratelimit_penaltybox_add_Name = "ratelimit.penaltybox_add"

// Fastly built-in function implementation of ratelimit.penaltybox_add
// Arguments may be:
// - ID, STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-penaltybox-add/
func Ratelimit_penaltybox_add[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	pb string, // IDENT
	entry string,
	ttl time.Time,
) error {
	// Need to be implemented
	return errors.FunctionError(
		Ratelimit_penaltybox_add_Name, "Not Implemented",
	)
}
