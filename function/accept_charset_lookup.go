package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Accept_charset_lookup_Name = "accept.charset_lookup"

// Fastly built-in function implementation of accept.charset_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-charset-lookup/
func Accept_charset_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	requestedCharsets, defaultValue, acceptHeader string,
) (string, error) {

	charsets := strings.Split(requestedCharsets, ":")
	index := len(charsets)
	for _, v := range strings.Split(acceptHeader, ",") {
		v = strings.TrimSpace(v)
		if idx := strings.Index(v, ";"); idx != -1 {
			v = v[0:idx]
		}
		for i := range charsets {
			if charsets[i] == v {
				if i < index {
					index = i
				}
			}
		}
	}

	if index < len(charsets) {
		return charsets[index], nil
	}
	return defaultValue, nil
}
