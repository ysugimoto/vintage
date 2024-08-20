package function

import (
	"regexp"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Regsuball_Name = "regsuball"

// Fastly built-in function implementation of regsuball
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/regsuball/
func Regsuball[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, pattern, replacement string,
) (string, error) {

	re, err := regexp.Compile(pattern)
	if err != nil {
		ctx.FastlyError = "EREGRECUR"
		return input, errors.FunctionError(
			Regsuball_Name, "Invalid regular expression pattern: %s, error: %w", pattern, err,
		)
	}

	expand := regsubExpandRE.ReplaceAllString(replacement, regsubExpandReplace)
	return re.ReplaceAllString(input, expand), nil
}
