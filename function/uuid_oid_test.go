package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.oid
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-oid/
func Test_Uuid_oid(t *testing.T) {
	ret, err := Uuid_oid(newTestRuntime())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if diff := cmp.Diff(ret, "6ba7b812-9dad-11d1-80b4-00c04fd430c8"); diff != "" {
		t.Errorf("Return value unmatch, diff=%s", diff)
	}
}
