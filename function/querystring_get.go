package function

import (
	"net/url"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_get_Name = "querystring.get"

// Fastly built-in function implementation of querystring.get
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-get/
func Querystring_get[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	u, name string,
) (string, error) {

	var qs string
	if idx := strings.Index(u, "?"); idx != -1 {
		qs = u[idx+1:]
	}

	// url.Value could not treat correctly empty query value:
	// ?name  => should return empty string, but returns empty string
	// ?name= => should return not set, but returns empty string
	// so we try to parse from RawQuery string, without using url.Value
	for _, query := range strings.Split(qs, "&") {
		sp := strings.Split(query, "=")
		if len(sp) < 2 || sp[0] == "" {
			continue
		}
		n, err := url.QueryUnescape(sp[0])
		if err != nil {
			continue
		}
		if n == name {
			return sp[1], nil
		}
	}
	return "", nil
}
