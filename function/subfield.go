package function

import (
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Subfield_Name = "subfield"

// Fastly built-in function implementation of subfield
// Arguments may be:
// - STRING, STRING, STRING
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/subfield/
func Subfield[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	subject, fieldName string,
	optional ...string,
) (string, error) {
	separator := ","
	if len(optional) > 0 {
		if len(optional[0]) > 1 {
			return "", errors.FunctionError(
				Subfield_Name, "Separator %s character must be a character", separator,
			)
		}
		separator = optional[0]
	}

	for _, v := range strings.Split(subject, separator) {
		kv := strings.SplitN(v, "=", 2)
		if kv[0] != fieldName {
			continue
		}
		if len(kv) > 1 {
			return kv[1], nil
		}
		return "", nil
	}
	return "", nil
}
