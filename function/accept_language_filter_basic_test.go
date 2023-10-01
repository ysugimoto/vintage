package function

import (
	"testing"
)

// Fastly built-in function testing implementation of accept.language_filter_basic
// Arguments may be:
// - STRING, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-language-filter-basic/
func Test_Accept_language_filter_basic(t *testing.T) {
	ret, err := Accept_language_filter_basic(
		newTestRuntime(),
		"en:de:fr:nl",
		"nl",
		"de,nl,ja",
		2,
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret != "de,nl" {
		t.Errorf("Unexpected value returned, expect=de,nl, got=%s", ret)
	}
}
