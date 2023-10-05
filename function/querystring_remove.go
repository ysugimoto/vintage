package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_remove_Name = "querystring.remove"

// Fastly built-in function implementation of querystring.remove
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-remove/
func Querystring_remove[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	url string,
) (string, error) {

	if idx := strings.Index(url, "?"); idx != -1 {
		url = url[0:idx]
	}

	return url, nil
}
