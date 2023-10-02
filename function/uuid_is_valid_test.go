package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of uuid.is_valid
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-valid/
func Test_Uuid_is_valid(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{input: "6ba7b810-9dad-11d1-80b4-00c04fd430c8", expect: true},
		{input: "ba7b810-9dad-11d1-80b4-00c04fd430c8", expect: false},
	}

	for i, tt := range tests {
		ret, err := Uuid_is_valid(newTestRuntime(), tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
