package function

import (
	"strconv"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Parse_time_delta_Name = "parse_time_delta"

// Fastly built-in function implementation of parse_time_delta
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/parse-time-delta/
func Parse_time_delta[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	specifier string,
) (int64, error) {

	var delta int64
	var stack []byte

	// Golang's time.ParseDuration does not recognize "d" and "D", date duration
	// so we need to parse manually.
	// And, Fastly document says that this function supports only individual parse.
	for _, s := range []byte(specifier) {
		switch s {
		case 0x64, 0x44: // "d" or "D"
			v, err := strconv.ParseInt(string(stack), 10, 64)
			if err != nil {
				return 0, errors.FunctionError(
					Parse_time_delta_Name,
					"Failed to parse dates as int: %s", string(stack),
				)
			}
			delta += v * 24 * 60 * 60
			goto OUTER
		case 0x68, 0x48: // "h" or "H"
			v, err := strconv.ParseInt(string(stack), 10, 64)
			if err != nil {
				return 0, errors.FunctionError(
					Parse_time_delta_Name,
					"Failed to parse hours as int: %s", string(stack),
				)
			}
			delta += v * 60 * 60
			goto OUTER
		case 0x6D, 0x4D: // "m" or "M"
			v, err := strconv.ParseInt(string(stack), 10, 64)
			if err != nil {
				return 0, errors.FunctionError(
					Parse_time_delta_Name,
					"Failed to parse minutes as int: %s", string(stack),
				)
			}
			delta += v * 60
			goto OUTER
		case 0x73, 0x53: // "s" or "S"
			v, err := strconv.ParseInt(string(stack), 10, 64)
			if err != nil {
				return 0, errors.FunctionError(
					Parse_time_delta_Name,
					"Failed to parse seconds as int: %s", string(stack),
				)
			}
			delta += v
			goto OUTER
		default:
			if s < 0x30 || s > 0x39 {
				return 0, errors.FunctionError(
					Parse_time_delta_Name,
					"Invalid character found: %s", string([]byte{s}),
				)
			}
			stack = append(stack, s)
		}
	}

OUTER:
	return delta, nil
}
