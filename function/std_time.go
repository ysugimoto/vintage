package function

import (
	"strconv"
	"strings"
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_time_Name = "std.time"

var Std_time_SupportFormats = []string{
	time.RFC1123,
	time.RFC822,
	time.RFC850,
	time.ANSIC,
	"2006-01-02 15:04:05", // ISO 8601 subset
}

// Fastly built-in function implementation of std.time
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/std-time/
func Std_time[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	fallback time.Time,
) (time.Time, error) {

	var t time.Time
	var err error
	for _, format := range Std_time_SupportFormats {
		t, err = time.ParseInLocation(format, input, time.UTC)
		if err == nil {
			return t, nil
		}
	}

	// If all formats are invalid, try to parse from unix epoch seconds
	ss := input
	if idx := strings.Index(input, "."); idx != -1 {
		ss = input[0:idx]
	}
	ts, err := strconv.ParseInt(ss, 10, 64)
	if err != nil {
		t = fallback.Add(0)
	} else {
		t = time.Unix(ts, 0)
	}

	if t.Unix() < 0 {
		return time.Time{}, nil
	}
	return t, nil
}
