package function

import (
	"math"
	"strconv"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strtol_Name = "std.strtol"

func Std_strtol_Hex(s string) (int64, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(s, "0x"), 16, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtol_Name, "Failed to parse string with base 16: %w", err,
		)
	}
	return i, nil
}

func Std_strtol_Octet(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtol_Name, "Failed to parse string with base 8: %w", err,
		)
	}
	return i, nil
}

func Std_strtol_Decimal(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtol_Name, "Failed to parse string with base 10: %w", err,
		)
	}
	return i, nil
}

// Special case: base 36 number could present "x" as number.
// So hex string like "0xABC" should treat as "xABC"
func Std_strtol_36(s string) (int64, error) {
	if s == "0" {
		return 0, nil
	}
	i, err := strconv.ParseInt(strings.TrimPrefix(s, "0"), 36, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtol_Name, "Failed to parse string with base 36: %w", err,
		)
	}
	return i, nil
}

func Std_strtol_Other(s string, base int64) (int64, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(s, "0x"), int(base), 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtol_Name, "Failed to parse string with base %d: %w", base, err,
		)
	}
	return i, nil
}

// Fastly built-in function implementation of std.strtol
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strtol/
func Std_strtol[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	base int64,
) (int64, error) {

	var i int64
	var err error
	switch base {
	case 0: // auto detection
		switch {
		case strings.HasPrefix(input, "0x"):
			i, err = Std_strtol_Hex(input)
		case strings.HasPrefix(input, "0"):
			i, err = Std_strtol_Octet(input)
		default:
			i, err = Std_strtol_Decimal(input)
		}
	case 8: // octet conversion
		i, err = Std_strtol_Octet(input)
	case 16: // hex conversion
		i, err = Std_strtol_Hex(input)
	case 36: // special case conversion
		i, err = Std_strtol_36(input)
	default: // hex conversion
		if base > 36 {
			err = errors.FunctionError(
				Std_strtol_Name, "Invalid base int. base must be from 0 to 36",
			)
		} else {
			i, err = Std_strtol_Other(input, base)
		}
	}

	if err != nil {
		ctx.FastlyError = "EPARSENUM"
		return 0, err
	}

	if i > int64(math.MaxInt64) || i < int64(math.MinInt64) {
		ctx.FastlyError = "ERANGE"
		return 0, errors.FunctionError(Std_strtol_Name, "Value overflow")
	}

	return i, nil
}
