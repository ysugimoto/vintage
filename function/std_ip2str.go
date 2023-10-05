package function

import (
	"net"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_ip2str_Name = "std.ip2str"

// Fastly built-in function implementation of std.ip2str
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-ip2str/
func Std_ip2str[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input net.IP,
) (string, error) {

	return input.String(), nil
}
