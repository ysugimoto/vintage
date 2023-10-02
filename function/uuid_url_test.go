package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.url
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-url/
func Test_Uuid_url(t *testing.T) {
	ret, err := Uuid_url(newTestRuntime())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if diff := cmp.Diff(ret, "6ba7b811-9dad-11d1-80b4-00c04fd430c8"); diff != "" {
		t.Errorf("Return value unmatch, diff=%s", diff)
	}
}
