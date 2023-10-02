package function

import (
	"github.com/gobwas/glob"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_globfilter_except_Name = "querystring.globfilter_except"

// Fastly built-in function implementation of querystring.globfilter_except
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-globfilter-except/
func Querystring_globfilter_except[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url, name string,
) (string, error) {
	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_globfilter_except_Name,
			"Failed to parse query: %s, error: %w", url, err,
		)
	}

	pattern, err := glob.Compile(name)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_globfilter_Name,
			"Invalid glob filter string: %s, error: %w", name, err,
		)
	}

	query.Filter(func(v string) bool {
		return pattern.Match(v)
	})
	return query.String(), nil
}
