package function

import (
	"net"
	"net/netip"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_ip_Name = "std.ip"

// Fastly built-in function implementation of std.ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-ip/
func Std_ip[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	addr, fallback string,
) (net.IP, error) {

	ip, err := netip.ParseAddr(addr)
	if err != nil {
		ip, err = netip.ParseAddr(fallback)
		if err != nil {
			return nil, errors.FunctionError(
				Std_ip_Name,
				"Failed to parse IP: %w", err,
			)
		}
	}

	switch {
	case ip.Is6():
		v := ip.As16()
		return net.IP(v[:]), nil
	case ip.Is4():
		v := ip.As4()
		return net.IP(v[:]), nil
	default:
		return nil, errors.FunctionError(
			Std_ip_Name, "Unexpected IP string: %s", ip.String(),
		)
	}
}
