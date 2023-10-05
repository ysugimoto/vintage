package function

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_units_Name = "time.units"

// Fastly built-in function implementation of time.units
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-units/
func Time_units[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	unit string,
	t time.Time,
) (string, error) {

	switch unit {
	case "s":
		return fmt.Sprintf("%d", t.Unix()), nil
	case "ms":
		v := float64(t.UnixMilli()) / 1000
		return strconv.FormatFloat(v, 'f', 3, 64), nil
	case "us":
		v := float64(t.UnixMicro()) / 1000000
		return strconv.FormatFloat(v, 'f', 6, 64), nil
	case "ns":
		v := float64(t.UnixNano()) / 1000000000
		return strconv.FormatFloat(v, 'f', 9, 64), nil
	default:
		ctx.FastlyError = "EINVAL"
		return "", errors.FunctionError(
			Time_units_Name,
			"Invalid unit string %s, allow either of s, ms, us and ns", unit,
		)
	}
}
