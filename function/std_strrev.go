package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_strrev_Name = "std.strrev"

// Fastly built-in function implementation of std.strrev
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrev/
func Std_strrev[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {
	// Argument validations
	// Note: Fastly does not consider multibyte. When string includes multibyte, return empty string
	// To check it, compare byte slice and rune slice length are same. It means strings are all byte representation
	bs := []byte(input)
	if len(bs) != len([]rune(input)) {
		return "", nil
	}
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}

	return string(bs), nil
}
