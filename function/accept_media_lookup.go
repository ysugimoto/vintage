package function

import (
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Accept_media_lookup_Name = "accept.media_lookup"

// Fastly built-in function implementation of accept.media_lookup
// Arguments may be:
// - STRING, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-media-lookup/
func Accept_media_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	lookup, defaultValue, pattern, accept string,
) (string, error) {

	mediaTypes := make(map[string]string)
	for _, v := range strings.Split(lookup, ":") {
		mediaTypes[v] = v
	}

	patterns := make(map[string]string)
	for _, v := range strings.Split(pattern, ":") {
		// Duplicate media types are not allowed among the first three arguments.
		if _, ok := mediaTypes[v]; ok {
			return "", errors.FunctionError(
				Accept_media_lookup_Name,
				"Third argument media must not duplicate in first argument",
			)
		}
		patterns[v] = v

		// Also add to group pattern
		if idx := strings.Index(v, "/"); idx != -1 {
			patterns[v[0:idx]+"/*"] = v
		}
	}

	for _, v := range strings.Split(accept, ",") {
		v = strings.TrimSpace(v)
		if idx := strings.Index(v, ";"); idx != -1 {
			v = v[0:idx]
		}
		if m, ok := mediaTypes[v]; ok {
			return m, nil
		} else if m, ok := patterns[v]; ok {
			return m, nil
		} else if v == "*/*" {
			return defaultValue, nil
		}
	}

	return "", nil
}
