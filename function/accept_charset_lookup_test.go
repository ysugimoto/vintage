package function

import (
	"testing"
)

// Fastly built-in function testing implementation of accept.charset_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-charset-lookup/
func Test_Accept_charset_lookup(t *testing.T) {
	ret, err := Accept_charset_lookup(
		newTestRuntime(),
		"iso-8859-5:iso-8859-2",
		"utf-8",
		"utf-8, iso-8859-1;q=0.5, *;q=0.1",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if ret != "utf-8" {
		t.Errorf("Unexpected value returned, expect=utf-8, got=%s", ret)
	}
}
