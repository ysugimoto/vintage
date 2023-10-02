package function

import (
	"fmt"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Http_status_matches_Name = "http_status_matches"

// Fastly built-in function implementation of http_status_matches
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/http-status-matches/
func Http_status_matches[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	status int64,
	format string,
) (bool, error) {
	var inverse bool
	if format[0] == 0x21 { // prefixed with "!"
		inverse = true
		format = format[1:]
	}

	match := fmt.Sprint(status)
	var isMatch bool
	for _, code := range strings.Split(format, ",") {
		if match == strings.TrimSpace(code) {
			isMatch = true
			break
		}
	}

	if isMatch {
		return !inverse, nil
	}
	return inverse, nil
}
