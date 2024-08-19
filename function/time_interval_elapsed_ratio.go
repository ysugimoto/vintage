package function

import (
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_interval_elapsed_ratio_Name = "time.interval_elapsed_ratio"

// Fastly built-in function implementation of substr
// Arguments may be:
// - STRING, INTEGER, INTEGER
// - STRING, INTEGER
// Reference: https://developer.fastly.com/documentation/reference/vcl/functions/date-and-time/time-interval-elapsed-ratio/
func Time_interval_elapsed_ratio[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	ref, start, end time.Time,
) (float64, error) {

	return float64(ref.Unix()-start.Unix()) / float64(end.Unix()-start.Unix()), nil
}
