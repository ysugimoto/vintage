package function

import (
	"strconv"
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Time_hex_to_time_Name = "time.hex_to_time"

// Fastly built-in function implementation of time.hex_to_time
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-hex-to-time/
func Time_hex_to_time[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	divisor int64,
	dividend string,
) (time.Time, error) {
	if divisor == 0 {
		return time.Time{}, errors.FunctionError(
			Time_hex_to_time_Name, "Could not divide by zero",
		)
	}

	ts, err := strconv.ParseInt(dividend, 16, 64)
	if err != nil {
		return time.Time{}, errors.FunctionError(
			Time_hex_to_time_Name, "Failed to decode hex string to inteter: %w", err,
		)
	}

	return time.Unix(int64(ts/divisor), 0), nil
}
