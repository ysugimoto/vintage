package function

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.anystr2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-anystr2ip/
func Test_Std_anystr2ip(t *testing.T) {
	tests := []struct {
		input    string
		fallback string
		expect   net.IP
	}{
		{
			input:    "0x8.010.2056",
			fallback: "10.0.0.0",
			expect:   net.ParseIP("8.8.8.8"),
		},
		{
			input:    "0x8.010.foo",
			fallback: "10.0.0.0",
			expect:   net.ParseIP("10.0.0.0"),
		},
		{
			input:    "0xc0.0.01001",
			fallback: "10.0.0.0",
			expect:   net.ParseIP("192.0.2.1"),
		},
	}

	for i, tt := range tests {
		ret, err := Std_anystr2ip(
			newTestRuntime(),
			tt.input,
			tt.fallback,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
