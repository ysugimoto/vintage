package function

import (
	"net"
	"net/netip"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_str2ip_Name = "std.str2ip"

// Fastly built-in function implementation of std.str2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-str2ip/
func Std_str2ip[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	addr, fallback string,
) (net.IP, error) {

	ip, err := netip.ParseAddr(addr)
	if err != nil {
		ip, err = netip.ParseAddr(fallback)
		if err != nil {
			return nil, errors.FunctionError(
				Std_str2ip_Name, "Failed to parse IP: %w", err,
			)
		}
	}

	if ip.Is6() {
		v := ip.As16()
		return net.IP(v[:]), nil
	} else if ip.Is4() {
		v := ip.As4()
		return net.IP(v[:]), nil
	}
	return nil, errors.FunctionError(
		Std_str2ip_Name,
		"Unexpected IP string: %s", ip.String(),
	)
}
