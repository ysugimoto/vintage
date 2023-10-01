package function

import (
	"encoding/binary"
	"fmt"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_hash_crc32b_Name = "digest.hash_crc32b"

// Fastly built-in function implementation of digest.hash_crc32b
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-crc32b/
func Digest_hash_crc32b[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input string,
) (string, error) {
	// https://github.com/whik/crc-lib-c/blob/master/crcLib.c#L527
	var crc uint32 = 0xffffffff
	for _, c := range []byte(input) {
		crc = crc ^ (uint32)(c)
		for i := 0; i < 8; i++ {
			if crc&0x1 != 0 {
				crc = (crc >> 1) ^ 0xEDB88320
			} else {
				crc = (crc >> 1)
			}
		}
	}
	crc = 0xffffffff ^ crc
	buf := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(buf, crc)

	return fmt.Sprintf("%08x", buf), nil
}
