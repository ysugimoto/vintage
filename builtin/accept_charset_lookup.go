package builtin

import (
	"strings"

	"github.com/ysugimoto/vintage"
)

// Fastly built-in function implementation of accept.charset_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-charset-lookup/
func Accept_charset_lookup(
	ctx *vintage.Context,
	requestedCharsets, defaultValue, acceptHeader string,
) string {
	var charsets []string
	for _, v := range strings.Split(requestedCharsets, ":") {
		charsets = append(charsets, v)
	}

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
		return charsets[index]
	}
	return defaultValue
}
