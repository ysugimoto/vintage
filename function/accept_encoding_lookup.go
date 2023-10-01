package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Accept_encoding_lookup_Name = "accept.encoding_lookup"

// Fastly built-in function implementation of accept.encoding_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-encoding-lookup/
func Accept_encoding_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	requestedContentEncodings, defaultValue, acceptHeader string,
) (string, error) {

	var encodings []string
	for _, v := range strings.Split(requestedContentEncodings, ":") {
		encodings = append(encodings, v)
	}

	index := len(encodings)
	for _, v := range strings.Split(acceptHeader, ",") {
		v = strings.TrimSpace(v)
		if idx := strings.Index(v, ";"); idx != -1 {
			v = v[0:idx]
		}
		for i := range encodings {
			if encodings[i] == v {
				if i < index {
					index = i
				}
			}
		}
	}

	if index < len(encodings) {
		return encodings[index], nil
	}
	return defaultValue, nil
}
