package function

import (
	"math/big"
	"net"
	"net/netip"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Addr_extract_bits_Name = "addr.extract_bits"

// Fastly built-in function implementation of addr.extract_bits
// Arguments may be:
// - IP, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-extract-bits/
func Addr_extract_bits[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	ip net.IP,
	startBit, bitCount int64,
) (int64, error) {

	if bitCount > 32 {
		return 0, errors.FunctionError(Addr_extract_bits_Name, "start_bit must be less than 32")
	}
	if bitCount+startBit > 128 {
		return 0, errors.FunctionError(Addr_extract_bits_Name, "start_bit plus bit_count must be less than 128")
	}

	addr, err := netip.ParseAddr(ip.String())
	if err != nil {
		return 0, errors.FunctionError(
			Addr_extract_bits_Name,
			"Failed to parse IP address %s, %s", ip.String(), err,
		)
	}
	bits := addr.AsSlice()
	if len(bits) == 4 { // If ipv4, pad with zeros on the left
		bits = append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, bits...)
	}

	// IPv6 (128 bits) integer is too large so need to calculate with bigint
	bi := new(big.Int)
	bi.SetBytes(bits)

	// Right shift of start_bit
	bi.Rsh(bi, uint(startBit))

	// Create mask bit of bit_count and calculate AND
	mask := big.NewInt(0)
	for i := 0; i < int(bitCount); i++ {
		c := big.NewInt(1)
		mask.Or(mask, c.Lsh(c, uint(i)))
	}
	bi.And(bi, mask)

	return bi.Int64(), nil
}
