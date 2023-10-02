package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.version5
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version5/
func Test_Uuid_version5(t *testing.T) {
	tests := []struct {
		namespace string
		input     string
		expect    string
	}{
		{
			namespace: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			input:     "www.fastly.com",
			expect:    "86573da0-058f-5871-a5b7-f3cb33447360",
		},
	}

	for i, tt := range tests {
		ret, err := Uuid_version5(
			newTestRuntime(),
			tt.namespace,
			tt.input,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
