package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_itoa_charset_Name = "std.itoa_charset"

// Fastly built-in function implementation of std.itoa_charset
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa-charset/
func Std_itoa_charset[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input int64,
	charsets string,
) (string, error) {
	// ref: strconv.FormatInt implementation of general case
	cs := []byte(charsets)
	base := int64(len(cs))

	var encoded []byte
	for input >= base {
		v := input / base
		encoded = append(encoded, cs[input-(v*base)])
		input = v
	}
	encoded = append(encoded, cs[input])
	reversed := make([]byte, len(encoded))
	for i := 0; i < len(encoded); i++ {
		reversed[len(encoded)-1-i] = encoded[i]
	}

	return string(reversed), nil
}
