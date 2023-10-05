package function

import (
	"sort"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Accept_language_filter_basic_Name = "accept.language_filter_basic"

// Fastly built-in function implementation of accept.language_filter_basic
// Arguments may be:
// - STRING, STRING, STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-language-filter-basic/
func Accept_language_filter_basic[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	lookup string,
	defaultValue string,
	language string,
	nmatches int64,
) (string, error) {

	languages := strings.Split(lookup, ":")
	var matches []int
	for _, v := range strings.Split(language, ",") {
		v = strings.TrimSpace(v)
		if idx := strings.Index(v, ";"); idx != -1 {
			v = v[0:idx]
		}
		for i := range languages {
			if languages[i] == v {
				matches = append(matches, i)
			}
		}
	}

	if len(matches) == 0 {
		return defaultValue, nil
	}

	if len(matches) > int(nmatches) {
		matches = matches[0:int(nmatches)]
	}
	sort.Ints(matches)
	ret := make([]string, len(matches))
	for i, m := range matches {
		ret[i] = languages[m]
	}

	return strings.Join(ret, ","), nil
}
