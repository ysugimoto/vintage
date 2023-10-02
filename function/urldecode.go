package function

import (
	"net/url"
	"strings"

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
	// "%" string is also encoded as "%25" so we need to decode properly
	input = strings.ReplaceAll(input, "%25", "%")

	dec, err := url.PathUnescape(input)
	if err != nil {
		return "", errors.FunctionError(
			Urldecode_Name,
			"Failed to urldecode string: %s", input,
		)
	}
	// url.PathUnescape does not decode "+" sign to white space so we also need to call url.QueryUnescape
	// in order to decode "+" sign into white space
	dec, err = url.QueryUnescape(dec)
	if err != nil {
		return "", errors.FunctionError(
			Urldecode_Name,
			"Failed to urldecode string: %s", input,
		)
	}

	return dec, nil
}
