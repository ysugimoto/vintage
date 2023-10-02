package function

import (
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_integer2time_Name = "std.integer2time"

// Fastly built-in function implementation of std.integer2time
// Arguments may be:
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/std-integer2time/
func Std_integer2time[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	val int64,
) (time.Time, error) {
	return time.Unix(val, 0).UTC(), nil
}
