package function

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Std_anystr2ip_Name = "std.anystr2ip"

func Std_anystr2ip_ParseString(v string) (int64, error) {
	// "0" always indicates zero
	if v == "0" {
		return 0, nil
	}

	switch {
	case strings.HasPrefix(v, "0x"): // hex
		return strconv.ParseInt(strings.TrimPrefix(v, "0x"), 16, 64)
	case strings.HasPrefix(v, "0"): // octet
		return strconv.ParseInt(strings.TrimPrefix(v, "0"), 8, 64)
	default: // decimal
		return strconv.ParseInt(v, 10, 64)
	}
}

func Std_anystr2ip_ParseIpv4(addr string) (net.IP, error) {
	var ip int64

	segments := strings.SplitN(addr, ".", 4)
	switch len(segments) {
	case 1:
		// first segment represetns all bits of IP (xxx.xxx.xxx.xxx)
		v, err := Std_anystr2ip_ParseString(segments[0])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		ip = v
	case 2:
		// first segment represetns first bits of IP (xxx.---.---.---)
		v1, err := Std_anystr2ip_ParseString(segments[0])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// second segment represetns remainings of IPs (---.xxx.xxx.xxx)
		v2, err := Std_anystr2ip_ParseString(segments[1])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		ip = (v1 << 24) | v2
	case 3:
		// first segment represetns first bits of IP (xxx.---.---.---)
		v1, err := Std_anystr2ip_ParseString(segments[0])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// second segment represetns second bits of IP (---.xxx.---.---)
		v2, err := Std_anystr2ip_ParseString(segments[1])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// third segment represetns remainings of IPs (---.---.xxx.xxx)
		v3, err := Std_anystr2ip_ParseString(segments[2])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		ip = (v1 << 24) | (v2 << 16) | v3
	case 4:
		// first segment represetns first bits of IP (xxx.---.---.---)
		v1, err := Std_anystr2ip_ParseString(segments[0])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// second segment represetns second bits of IP (---.xxx.---.---)
		v2, err := Std_anystr2ip_ParseString(segments[1])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// third segment represetns third bits of IP (---.---.xxx.---)
		v3, err := Std_anystr2ip_ParseString(segments[2])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		// last segment represetns fourth bits of IP (---.---.---.xxx)
		v4, err := Std_anystr2ip_ParseString(segments[3])
		if err != nil {
			return nil, errors.FunctionError(
				Std_anystr2ip_Name,
				"Failed to parse IPv4 string: %w", err,
			)
		}
		ip = (v1 << 24) | (v2 << 16) | (v3 << 8) | v4
	default:
		return nil, errors.FunctionError(
			Std_anystr2ip_Name,
			"Invalid IPv4 string: %s", addr,
		)
	}

	return net.ParseIP(
		fmt.Sprintf("%d.%d.%d.%d", ((ip >> 24) & 0xFF), ((ip >> 16) & 0xFF), ((ip >> 8) & 0xFF), (ip & 0xFF)),
	), nil
}

// Fastly built-in function implementation of std.anystr2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-anystr2ip/
func Std_anystr2ip[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	addr, fallback string,
) (net.IP, error) {
	// TODO: support IPv6 string to parse
	if strings.Contains(addr, ":") {
		return nil, errors.FunctionError(
			Std_anystr2ip_Name, "Does not support IPv6 format string",
		)
	}
	// IPv4 parsing
	if v, err := Std_anystr2ip_ParseIpv4(addr); err != nil {
		return net.ParseIP(fallback), nil
	} else {
		return v, nil
	}
}
