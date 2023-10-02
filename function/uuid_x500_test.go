package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.x500
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-x500/
func Test_Uuid_x500(t *testing.T) {
	ret, err := Uuid_x500(newTestRuntime())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if diff := cmp.Diff(ret, "6ba7b814-9dad-11d1-80b4-00c04fd430c8"); diff != "" {
		t.Errorf("Return value unmatch, diff=%s", diff)
	}
}
