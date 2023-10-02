package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.itoa_charset
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa-charset/
func Test_Std_itoa_charset(t *testing.T) {
	tests := []struct {
		input   int64
		charset string
		expect  string
	}{
		{input: 8, charset: "ab", expect: "baaa"},
		{input: 0xdeadbeef, charset: "0123456789ABCDEF", expect: "DEADBEEF"},
		{input: 0xdeadbeef, charset: "0123456789ABCDE", expect: "16CEB1BDE"},
		{input: 4825434263756878946, charset: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", expect: "5KsxueeBfd8"},
	}

	for i, tt := range tests {
		ret, err := Std_itoa_charset(
			newTestRuntime(),
			tt.input,
			tt.charset,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
