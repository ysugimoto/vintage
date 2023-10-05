package function

import (
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_is_after_Name = "time.is_after"

// Fastly built-in function implementation of time.is_after
// Arguments may be:
// - TIME, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-is-after/
func Time_is_after[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	t1, t2 time.Time,
) (bool, error) {

	return t1.After(t2), nil
}
