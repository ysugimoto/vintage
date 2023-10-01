package function

import (
	"net"
	"testing"
)

// Fastly built-in function testing implementation of addr.extract_bits
// Arguments may be:
// - IP, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-extract-bits/
func Test_Addr_extract_bits(t *testing.T) {

	ret, err := Addr_extract_bits(
		newTestRuntime(),
		net.ParseIP("151.101.2.217"),
		0,
		8,
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret != 217 {
		t.Errorf("Unexpected value returned, expect=217, got=%d", ret)
	}
}
