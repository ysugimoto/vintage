package function

import (
	"net/url"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Urldecode_Name = "urldecode"

// Fastly built-in function implementation of urldecode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/urldecode/
func Urldecode[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	dec, err := url.QueryUnescape(input)
	if err != nil {
		return "", errors.FunctionError(
			Urldecode_Name,
			"Failed to urldecode string: %s", input,
		)
	}

	return dec, nil
}
