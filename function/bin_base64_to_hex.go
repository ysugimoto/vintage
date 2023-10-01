package function

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Bin_base64_to_hex_Name = "bin.base64_to_hex"

// Fastly built-in function implementation of bin.base64_to_hex
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/bin-base64-to-hex/
func Bin_base64_to_hex[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {
	// If input value is empty, return empty string
	if input == "" {
		return "", nil
	}

	var buf bytes.Buffer
	if _, err := io.Copy(
		hex.NewEncoder(&buf),
		base64.NewDecoder(base64.StdEncoding, strings.NewReader(input)),
	); err != nil {
		// If the Base64-encoded string s is not valid Base64, then fastly.error will be set to EINVAL.
		ctx.FastlyError = "EINVAL"
		return "", errors.FunctionError(
			Bin_base64_to_hex_Name,
			"Failed to decode base64 string / encode hex string",
		)
	}
	return strings.ToUpper(buf.String()), nil
}
