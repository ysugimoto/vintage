package function

import (
	"net"
	"testing"
)

// Fastly built-in function testing implementation of addr.is_ipv4
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-is-ipv4/
func Test_Addr_is_ipv4(t *testing.T) {

	table := []struct {
		ip     net.IP
		expect bool
	}{
		{
			ip:     net.ParseIP("127.0.0.1"),
			expect: true,
		},
		{
			ip:     net.ParseIP("2001:DB8:0:0:8:800:200C:417A"),
			expect: false,
		},
	}

	for _, tt := range table {
		ret, err := Addr_is_ipv4(
			newTestRuntime(),
			tt.ip,
		)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.expect {
			t.Errorf("Unexpected value returned, expect=%t, got=%t", tt.expect, ret)
		}
	}
}
