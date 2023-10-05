package function

import (
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_sub_Name = "time.sub"

// Fastly built-in function implementation of time.sub
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-sub/
func Time_sub[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	t1 time.Time,
	rtime time.Duration,
) (time.Time, error) {

	return t1.Add(-rtime), nil
}
