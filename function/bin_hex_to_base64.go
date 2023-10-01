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

const Bin_hex_to_base64_Name = "bin.hex_to_base64"

// Fastly built-in function implementation of bin.hex_to_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/bin-hex-to-base64/
func Bin_hex_to_base64[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	var buf bytes.Buffer
	if _, err := io.Copy(
		base64.NewEncoder(base64.StdEncoding, &buf),
		hex.NewDecoder(strings.NewReader(input)),
	); err != nil {
		// If the hex string s is not valid, then fastly.error will be set to EINVAL.
		ctx.FastlyError = "EINVAL"
		return "", errors.FunctionError(
			Bin_hex_to_base64_Name,
			"Failed to decode hex string / encode to base64 string",
		)
	}

	return buf.String(), nil
}
