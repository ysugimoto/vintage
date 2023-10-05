package function

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Boltsort_sort_Name = "boltsort.sort"

// Fastly built-in function implementation of boltsort.sort
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/boltsort-sort/
func Boltsort_sort[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {

	parsed, err := url.ParseRequestURI(input)
	if err != nil {
		// This function does not raise an error
		return input, nil
	}
	query := parsed.Query()
	var sorted []string
	for k := range query {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)

	// Build RawQuery by sorted order
	var rawQuery []string
	for i := range sorted {
		qs := query[sorted[i]]
		sort.Strings(qs)
		for j := range qs {
			rawQuery = append(rawQuery, fmt.Sprintf(
				"%s=%s",
				url.QueryEscape(sorted[i]), url.QueryEscape(qs[j]),
			))
		}
	}
	parsed.RawQuery = strings.Join(rawQuery, "&")

	return parsed.String(), nil
}
