package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Xml_escape_Name = "xml_escape"

// Fastly built-in function implementation of xml_escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/xml-escape/
func Xml_escape[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	var escaped []byte
	for _, b := range []byte(input) {
		switch b {
		case 0x26: // "&"
			escaped = append(escaped, []byte("&amp;")...)
		case 0x3C: // "<"
			escaped = append(escaped, []byte("&lt;")...)
		case 0x3E: // ">"
			escaped = append(escaped, []byte("&gt;")...)
		case 0x27: // "'"
			escaped = append(escaped, []byte("&apos;")...)
		case 0x22: // '"'
			escaped = append(escaped, []byte("&quot;")...)
		default:
			escaped = append(escaped, b)
		}
	}

	return string(escaped), nil
}
