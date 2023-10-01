package function

import (
	"testing"
)

// Fastly built-in function testing implementation of h2.disable_header_compression
// Arguments may be:
// - STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h2-disable-header-compression/
func Test_H2_disable_header_compression(t *testing.T) {

	ctx := newTestRuntime()
	err := H2_disable_header_compression(
		ctx,
		"Authorization",
		"Secret",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	// compare stacked result
	headers := ctx.DisableCompressionHeaders
	if len(headers) != 2 {
		t.Errorf("Disabled heaers count must be 2, got %d", len(headers))
		return
	}
	for i, v := range []string{"Authorization", "Secret"} {
		if headers[i] != v {
			t.Errorf("%d header value must be %s, got=%s", i, v, headers[i])
		}
	}
}
