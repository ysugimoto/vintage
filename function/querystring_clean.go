package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_clean_Name = "querystring.clean"

// Fastly built-in function implementation of querystring.clean
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-clean/
func Querystring_clean[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url string,
) (string, error) {
	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_clean_Name,
			"Failed to parse url: %s, error: %w", url, err,
		)
	}

	query.Clean()
	return query.String(), nil
}
