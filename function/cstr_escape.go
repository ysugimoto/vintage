package function

import (
	"fmt"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Cstr_escape_Name = "cstr_escape"

var Cstr_escape_CharacterMap = map[byte][]byte{
	0x22: []byte("\""),
	0x5C: []byte("\\"),
	0x08: []byte("\\b"),
	0x09: []byte("\\t"),
	0x0A: []byte("\\n"),
	0x0B: []byte("\\v"),
	0x0D: []byte("\\r"),
}

// Fastly built-in function implementation of cstr_escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/cstr-escape/
func Cstr_escape[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	str string,
) (string, error) {
	var escaped []byte
	for _, b := range []byte(str) {
		if v, ok := Cstr_escape_CharacterMap[b]; ok {
			escaped = append(escaped, v...)
			continue
		}
		if b < 0x1F || 0x7F < b {
			escaped = append(escaped, []byte(fmt.Sprintf("\\x%x", b))...)
			continue
		}
		escaped = append(escaped, b)
	}

	return string(escaped), nil
}
