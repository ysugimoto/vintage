package function

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.str2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-str2ip/
func Test_Std_str2ip(t *testing.T) {
	tests := []struct {
		input    string
		fallback string
		expect   string
	}{
		{input: "192.0.2.1", fallback: "192.0.2.2", expect: "192.0.2.1"},
		{input: "192.0.2.256", fallback: "192.0.2.2", expect: "192.0.2.2"},
		{input: "2001:db8::1d", fallback: "2001:db8::1e", expect: "2001:db8::1d"},
		{input: "2001:db8::-1", fallback: "2001:db8::1e", expect: "2001:db8::1e"},
	}

	for i, tt := range tests {
		ret, err := Std_str2ip(
			newTestRuntime(),
			tt.input,
			tt.fallback,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, net.ParseIP(tt.expect)); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
