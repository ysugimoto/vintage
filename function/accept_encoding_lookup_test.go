package function

import (
	"testing"
)

// Fastly built-in function testing implementation of accept.encoding_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-encoding-lookup/
func Test_Accept_encoding_lookup(t *testing.T) {
	ret, err := Accept_encoding_lookup(
		newTestRuntime(),
		"br:compress:deflate:gzip",
		"identity",
		"deflate, br, unknown",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret != "br" {
		t.Errorf("Unexpected value returned, expect=br, got=%s", ret)
	}
}
