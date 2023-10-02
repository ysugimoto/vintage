package function

import (
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_filter_Name = "querystring.filter"

// Fastly built-in function implementation of querystring.filter
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-filter/
func Querystring_filter[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url, names string,
) (string, error) {
	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_filter_Name,
			"Failed to parse query: %s, error: %w", url, err,
		)
	}
	filterMap := make(map[string]struct{})
	for _, f := range strings.Split(names, Querystring_filtersep_Sign) {
		filterMap[f] = struct{}{}
	}

	query.Filter(func(name string) bool {
		_, ok := filterMap[name]
		return !ok
	})

	return query.String(), nil
}
