package function

import (
	"net/url"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Urlencode_Name = "urlencode"

// Fastly built-in function implementation of urlencode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/urlencode/
func Urlencode[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {
	// url.QueryEscape encodes white space to "+" so we should use url.PathEscape
	// in order to encode white space to "%20"
	return url.PathEscape(input), nil
}
