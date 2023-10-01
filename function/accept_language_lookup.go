package function

import (
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Accept_language_lookup_Name = "accept.language_lookup"

// Fastly built-in function implementation of accept.language_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-language-lookup/
func Accept_language_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	lookup, defaultValue, language string,
) (string, error) {

	var languages []string
	for _, v := range strings.Split(lookup, ":") {
		languages = append(languages, v)
	}

	index := len(languages)
	for _, v := range strings.Split(language, ",") {
		v = strings.TrimSpace(v)
		if idx := strings.Index(v, ";"); idx != -1 {
			v = v[0:idx]
		}
		for i := range languages {
			if languages[i] == v {
				if i < index {
					index = i
				}
			}
		}
	}

	if index < len(languages) {
		return languages[index], nil
	}
	return defaultValue, nil
}
