package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.version3
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version3/
func Test_Uuid_version3(t *testing.T) {
	tests := []struct {
		namespace string
		input     string
		expect    string
	}{
		{
			namespace: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			input:     "www.fastly.com",
			expect:    "3f22bcdf-f888-31a6-9575-d1588cb14ff4",
		},
	}

	for i, tt := range tests {
		ret, err := Uuid_version3(
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
