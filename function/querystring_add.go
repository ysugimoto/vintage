package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_add_Name = "querystring.add"

// Fastly built-in function implementation of querystring.add
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-add/
func Querystring_add[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url, name, value string,
) (string, error) {

	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_add_Name,
			"Failed to parse querystring: %s, error: %w", url, err,
		)
	}

	query.Add(name, value)
	return query.String(), nil
}
