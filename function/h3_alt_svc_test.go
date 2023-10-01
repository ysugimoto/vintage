package function

import (
	"testing"
)

// Fastly built-in function testing implementation of h3.alt_svc
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h3-alt-svc/
func Test_H3_alt_svc(t *testing.T) {
	ctx := newTestRuntime()
	err := H3_alt_svc(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	// compare stacked result
	if ctx.H3AltSvc != true {
		t.Errorf("AltSvc should be true, got=%t", ctx.H3AltSvc)
	}
}
