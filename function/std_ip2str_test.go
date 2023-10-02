package function

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Fastly built-in function testing implementation of std.ip2str
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-ip2str/
func Test_Std_ip2str(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "192.0.2.1", expect: "192.0.2.1"},
		{input: "2001:db8::1d", expect: "2001:db8::1d"},
	}

	for i, tt := range tests {
		ret, err := Std_ip2str(
			newTestRuntime(),
			net.ParseIP(tt.input),
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
