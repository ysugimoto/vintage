package function

import (
	"net"
	"net/netip"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Addr_is_ipv6_Name = "addr.is_ipv6"

// Fastly built-in function implementation of addr.is_ipv6
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-is-ipv6/
func Addr_is_ipv6[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	ip net.IP,
) (bool, error) {
	addr, err := netip.ParseAddr(ip.String())
	if err != nil {
		return false, errors.FunctionError(
			Addr_is_ipv6_Name,
			"Failed to parse IP address %s, %s", ip.String(), err,
		)
	}
	return addr.Is6(), nil
}
