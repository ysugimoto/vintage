package function

import (
	"testing"
)

// Fastly built-in function testing implementation of accept.language_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-language-lookup/
func Test_Accept_language_lookup(t *testing.T) {

	ret, err := Accept_charset_lookup(
		newTestRuntime(),
		"en:de:fr:nl",
		"nl",
		"ja, unknown",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret != "nl" {
		t.Errorf("Unexpected value returned, expect=nl, got=%s", ret)
	}
}
