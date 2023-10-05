package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_set_Name = "querystring.set"

// Fastly built-in function implementation of querystring.set
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-set/
func Querystring_set[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url, name, value string,
) (string, error) {

	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_set_Name,
			"Failed to parse url query: %s, error: %w", url, err,
		)
	}

	query.Set(name, value)
	return query.String(), nil
}
