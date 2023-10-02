package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Substr_Name = "substr"

// Fastly built-in function implementation of substr
// Arguments may be:
// - STRING, INTEGER, INTEGER
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/substr/
func Substr[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
	offset int64,
	optional ...int64,
) (string, error) {
	var length *int64
	if len(optional) > 0 {
		length = &optional[0]
	}

	var start, end int
	if offset < 0 {
		start = len(input) + int(offset)
	} else {
		start = int(offset)
	}

	if length == nil {
		end = len(input)
	} else if *length < 0 {
		if offset < 0 {
			end = len(input) + int(*length) + 1
		} else {
			end = len(input) + int(*length)
		}
	} else {
		if offset < 0 {
			end = start + int(*length)
		} else {
			end = start + int(*length) + 1
		}
	}
	if end > len(input) {
		end = len(input)
	}

	if start > len(input) {
		return "", errors.FunctionError(
			Substr_Name,
			"Invalid start offset %d against input string %s", offset, input,
		)
	}
	if end <= start {
		return "", nil
	}
	return string(input[start:end]), nil
}
