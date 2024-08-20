package function

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_base64url_decode_Name = "digest.base64url_decode"

// Fastly built-in function implementation of digest.base64url_decode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url-decode/
func Digest_base64url_decode[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	removed := Digest_base64url_decode_removeInvalidCharacters(input)
	dec, err := base64.URLEncoding.DecodeString(removed)
	if err != nil {
		return "", errors.FunctionError(
			Digest_base64url_decode_Name,
			"Failed to decode base64-url string, %s", err,
		)
	}

	return string(terminateNullByte(dec)), nil
}

func Digest_base64url_decode_removeInvalidCharacters(input string) string {
	removed := new(bytes.Buffer)
	r := bufio.NewReader(strings.NewReader(input))

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		switch {
		case b >= 0x41 && b <= 0x5A: // A-Z
			removed.WriteByte(b)
		case b >= 0x61 && b <= 0x7A: // a-z
			removed.WriteByte(b)
		case b >= 0x31 && b <= 0x39: // 0-9
			removed.WriteByte(b)
		case b == 0x2B: // + should replace to -
			removed.WriteByte(0x2D)
		case b == 0x2F: // / should replace to _
			removed.WriteByte(0x5F)
		case b == 0x2D || b == 0x5F: // + or /
			removed.WriteByte(b)
		case b == 0x3D: // =
			// If "=" sign found, next byte must also be "="
			if peek, err := r.Peek(1); err != nil && peek[0] == 0x3D {
				removed.WriteByte(b)
				removed.WriteByte(b)
				// nolint:errcheck
				r.ReadByte() // skip next "=" character
				continue
			}
			// Otherwise, treat as invalid character, stop decoding
			return string(base64_padding(removed.Bytes()))
		default:
			// Invalid characters, skip it
		}
	}

	return string(base64_padding(removed.Bytes()))
}
