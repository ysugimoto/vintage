package function

import (
	"github.com/ysugimoto/vintage/runtime/core"
)

const Querystring_filtersep_Name = "querystring.filtersep"

var Querystring_filtersep_Sign = string([]byte{0xFF})

// Fastly built-in function implementation of querystring.filtersep
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-filtersep/
func Querystring_filtersep[T core.EdgeRuntime](ctx *core.Runtime[T]) (string, error) {
	return Querystring_filtersep_Sign, nil
}
