package function

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strtof_Name = "std.strtof"

func Std_strtof_Decimal(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtof_Name, "Failed to parse string to decimal float: %s", s,
		)
	}
	return f, nil
}

var hexMantissaSuffix = regexp.MustCompile(`p[+-][0-9]+$`)

func Std_strtof_Hex(s string) (float64, error) {
	// To convert hex float string correctly in Golang, mantissa suffix is needed
	if !hexMantissaSuffix.MatchString(s) {
		s += "p0"
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.FunctionError(
			Std_strtof_Name, "Failed to parse string to hex float: %s", s,
		)
	}
	return f, nil
}

// Fastly built-in function implementation of std.strtof
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strtof/
func Std_strtof[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	base int64,
) (float64, error) {
	switch base {
	case 0:
		if strings.HasPrefix(input, "0x") {
			return Std_strtof_Hex(input)
		}
		return Std_strtof_Decimal(input)
	case 10:
		if strings.HasPrefix(input, "0x") {
			ctx.FastlyError = "EPARSENUM"
			return 0, errors.FunctionError(
				Std_strtof_Name,
				"string must not have 0x prefix of when base number is 10",
			)
		}
		return Std_strtof_Decimal(input)
	case 16:
		if !strings.HasPrefix(input, "0x") {
			ctx.FastlyError = "EPARSENUM"
			return 0, errors.FunctionError(
				Std_strtof_Name,
				"string must have 0x prefix of when base number is 16",
			)
		}
		return Std_strtof_Hex(input)
	default:
		ctx.FastlyError = "EPARSENUM"
		return 0, errors.FunctionError(
			Std_strtof_Name,
			"Base number accepts only 0, 10 and 16",
		)
	}
}
