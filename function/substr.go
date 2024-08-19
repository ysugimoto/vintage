package function

import (
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
		if start < 0 {
			return "", nil
		}
	} else {
		start = int(offset)
	}

	switch {
	case length == nil:
		end = len(input)
	case *length < 0:
		end = len(input) + int(*length)
	default:
		end = start + int(*length)
		if end < 0 {
			return "", nil
		}
	}
	if end > len(input) {
		end = len(input)
	}

	if start > len(input) {
		return "", nil
	}
	if end <= start {
		return "", nil
	}
	return string(input[start:end]), nil
}
