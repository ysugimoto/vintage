package function

import (
	"regexp"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_regfilter_Name = "querystring.regfilter"

// Fastly built-in function implementation of querystring.regfilter
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-regfilter/
func Querystring_regfilter[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url, name string,
) (string, error) {

	query, err := lib.ParseQuery(url)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_regfilter_Name,
			"Failed to parse query: %s, error: %w", url, err,
		)
	}

	re, err := regexp.Compile(name)
	if err != nil {
		return "", errors.FunctionError(
			Querystring_regfilter_Name,
			"Invalid regexp pattern: %s, error: %s", name, err,
		)
	}

	query.Filter(func(key string) bool {
		return !re.MatchString(key)
	})

	return query.String(), nil
}
