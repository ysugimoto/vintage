package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_sort_Name = "querystring.sort"

// Fastly built-in function implementation of querystring.sort
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-sort/
func Querystring_sort[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url string,
) (string, error) {

	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_sort_Name,
			"Failed to parse urquery: %s, error: %w", url, err,
		)
	}

	query.Sort(lib.SortAsc)
	return query.String(), nil
}
