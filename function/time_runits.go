package function

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_runits_Name = "time.runits"

// Fastly built-in function implementation of time.runits
// Arguments may be:
// - STRING, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-runits/
func Time_runits[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	unit string,
	rtime time.Duration,
) (string, error) {
	switch unit {
	case "s":
		return fmt.Sprintf("%d", int64(rtime.Seconds())), nil
	case "ms":
		v := float64(rtime.Milliseconds()) / 1000
		return strconv.FormatFloat(v, 'f', 3, 64), nil
	case "us":
		v := float64(rtime.Microseconds()) / 1000000
		return strconv.FormatFloat(v, 'f', 6, 64), nil
	case "ns":
		v := float64(rtime.Nanoseconds()) / 1000000000
		return strconv.FormatFloat(v, 'f', 9, 64), nil
	default:
		ctx.FastlyError = "EINVAL"
		return "", errors.FunctionError(
			Time_runits_Name,
			"Invalid unit string %s, allow either of s, ms, us and ns", unit,
		)
	}
}
