package function

import (
	"fmt"
	"unicode/utf8"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Json_escape_Name = "json.escape"

var Json_escape_CharacterMap = map[rune][]rune{
	0x22: []rune("\""),
	0x5C: []rune("\\"),
	0x08: []rune("\\b"),
	0x09: []rune("\\t"),
	0x0A: []rune("\\n"),
	0x0C: []rune("\\f"),
	0x0D: []rune("\\r"),
}

// ref: http://www5d.biglobe.ne.jp/~noocyte/Programming/CharCode.html (JP)
func Json_escape_toUTF16SurrogatePair(r rune) []rune {
	r -= 0x10000
	upper := ((r >> 10) & 0x03FF) + 0xD800
	lower := (r & 0x03FF) + 0xDC00
	return []rune(fmt.Sprintf("\\u%04X\\u%04X", upper, lower))
}

// Fastly built-in function implementation of json.escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/json-escape/
func Json_escape[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	str string,
) (string, error) {
	// Preparation: check provided is valid in UTF-8 sequence
	if !utf8.ValidString(str) {
		return "", nil
	}

	var escaped []rune
	for _, r := range str {
		if v, ok := Json_escape_CharacterMap[r]; ok {
			escaped = append(escaped, v...)
			continue
		}
		if r < 0x1F || r == 0x7F || r == 0x2028 || r == 0x2029 {
			escaped = append(escaped, []rune(fmt.Sprintf("\\u%04x", r))...)
			continue
		}
		if r > 0xFFFF {
			escaped = append(escaped, Json_escape_toUTF16SurrogatePair(r)...)
			continue
		}
		escaped = append(escaped, r)
	}

	return string(escaped), nil
}
