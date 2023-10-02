package function

import (
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_add_Name = "time.add"

// Fastly built-in function implementation of time.add
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-add/
func Time_add[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	t1 time.Time,
	t2 time.Duration,
) (time.Time, error) {
	return t1.Add(t2), nil
}
